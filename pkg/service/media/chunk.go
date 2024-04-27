package media

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/shenghui0779/yiigo"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"api/ent"
	"api/ent/file"
	"api/lib/db"
	"api/lib/log"
	"api/pkg/internal"
	"api/pkg/result"
)

type FormChunk struct {
	FileID   string `form:"file_id" valid:"required"`
	FileMD5  string `form:"file_md5" valid:"required"`
	FileName string `form:"file_name" valid:"required"`
	Index    int    `form:"index" valid:"gt=0"`
	Blocks   int    `form:"blocks" valid:"gt=0"`
}

// Chunk 分片上传
func Chunk(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	form := new(FormChunk)
	if err := internal.BindForm(r, form); err != nil {
		result.ErrParams(result.E(err)).JSON(w, r)
		return
	}

	record, err := db.Client().File.Query().Unique(false).Select(file.FieldID, file.FieldSize).Where(file.Fingerprint(form.FileMD5)).First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		log.Error(ctx, "Error query file", zap.Error(err))
		result.ErrSystem().JSON(w, r)
		return
	}
	// 文件存在，则秒传
	if record != nil {
		// 非最后一块，直接返回成功
		if form.Index < form.Blocks {
			result.OK().JSON(w, r)
			return
		}

		mediaID := internal.MediaID(form.FileMD5)
		// 创建Media
		_, err = db.Client().Media.Create().SetMediaID(mediaID).SetFileName(form.FileName).SetFileID(record.ID).Save(ctx)
		if err != nil {
			log.Error(ctx, "Error create media", zap.Error(err))
			result.ErrSystem().JSON(w, r)
			return
		}

		result.OK(result.V(&RespUpload{
			MediaID:  mediaID,
			MediaURL: internal.MediaURL(mediaID),
			FileName: form.FileName,
			FileSize: record.Size,
		})).JSON(w, r)

		return
	}

	src, _, err := r.FormFile("media")
	if err != nil {
		log.Error(ctx, "Error form file", zap.Error(err))
		result.ErrSystem().JSON(w, r)
		return
	}
	defer src.Close()

	chunckPath := internal.MediaChunk(form.FileID, form.Index)
	dst, err := yiigo.CreateFile(chunckPath)
	if err != nil {
		log.Error(ctx, "Error create chunck", zap.Error(err))
		result.ErrSystem().JSON(w, r)
		return
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		log.Error(ctx, "Error copy src to file", zap.Error(err))
		result.ErrSystem().JSON(w, r)
		return
	}

	// 非最后一块，保存即可
	if form.Index < form.Blocks {
		result.OK().JSON(w, r)
		return
	}

	// 合并文件
	merge(ctx, form).JSON(w, r)
}

func merge(ctx context.Context, form *FormChunk) result.Result {
	defer func() {
		// 清除分片临时文件
		os.RemoveAll(path.Clean(fmt.Sprintf("%s/chuncks/%s", viper.GetString("media.path"), form.FileID)))
	}()

	mediaID := internal.MediaID(form.FileMD5)
	mediaPath := internal.MediaFile(mediaID, form.FileName)

	f, err := yiigo.CreateFile(mediaPath)
	if err != nil {
		log.Error(ctx, "Error create file", zap.Error(err))
		return result.ErrSystem()
	}
	defer f.Close()

	// 合并文件并校验MD5值
	h := md5.New()
	for i := 1; i <= form.Blocks; i++ {
		tmp, err := os.Open(internal.MediaChunk(form.FileID, i))
		if err != nil {
			log.Error(ctx, "Error open chunck", zap.Error(err))
			return result.ErrSystem()
		}

		io.Copy(h, tmp)
		tmp.Seek(0, 0)
		io.Copy(f, tmp)
		tmp.Close()
	}
	fingerprint := hex.EncodeToString(h.Sum(nil))
	// 校验MD5值
	if fingerprint != form.FileMD5 {
		// 不一致则作废，删除文件
		os.RemoveAll(mediaPath)
		return result.ErrData(result.E(errors.New("MD5值校验失败")))
	}

	// 如果是图片，则获取图片宽高
	imgExif, err := ParseMediaEXIF(mediaPath)
	if err != nil {
		log.Error(ctx, "Error parse EXIF", zap.Error(err))
		return result.ErrSystem()
	}

	stat, _ := f.Stat()

	// 创建文件
	record, err := db.Client().File.Create().
		SetFingerprint(fingerprint).
		SetSize(stat.Size()).
		SetFormat(imgExif.Format).
		SetWidth(imgExif.Width).
		SetHeight(imgExif.Height).
		SetOrientation(imgExif.Orientation).
		Save(ctx)
	if err != nil {
		// DB失败，删除文件
		os.RemoveAll(mediaPath)
		log.Error(ctx, "Error create file", zap.Error(err))

		return result.ErrSystem()
	}

	// 创建Media
	_, err = db.Client().Media.Create().SetMediaID(mediaID).SetFileName(form.FileName).SetFileID(record.ID).Save(ctx)
	if err != nil {
		log.Error(ctx, "Error create media", zap.Error(err))
		return result.ErrSystem()
	}

	return result.OK(result.V(&RespUpload{
		MediaID:  mediaID,
		MediaURL: internal.MediaURL(mediaID),
		FileName: form.FileName,
		FileSize: stat.Size(),
	}))
}
