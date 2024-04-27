package media

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"net/http"
	"os"

	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"

	"api/ent"
	"api/ent/file"
	"api/lib/db"
	"api/lib/log"
	"api/pkg/internal"
	"api/pkg/result"
)

type RespUpload struct {
	MediaID  string `json:"media_id"`
	MediaURL string `json:"media_url"`
	FileName string `json:"file_name"`
	FileSize int64  `json:"file_size"`
}

func Upload(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	src, fh, err := r.FormFile("media")
	if err != nil {
		log.Error(ctx, "Error form file", zap.Error(err))
		result.ErrSystem().JSON(w, r)

		return
	}
	defer src.Close()

	h := md5.New()
	if _, err = io.Copy(h, src); err != nil {
		log.Error(ctx, "Error copy src to hash", zap.Error(err))
		result.ErrSystem().JSON(w, r)

		return
	}
	fingerprint := hex.EncodeToString(h.Sum(nil))
	mediaID := internal.MediaID(fingerprint)

	record, err := db.Client().File.Query().Unique(false).Select(file.FieldID, file.FieldSize).Where(file.Fingerprint(fingerprint)).First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		log.Error(ctx, "Error query file", zap.Error(err))
		result.ErrSystem().JSON(w, r)
		return
	}
	// 文件存在，则秒传
	if record != nil {
		_, err = db.Client().Media.Create().SetMediaID(mediaID).SetFileName(fh.Filename).SetFileID(record.ID).Save(ctx)
		if err != nil {
			log.Error(ctx, "Error create media", zap.Error(err))
			result.ErrSystem().JSON(w, r)
			return
		}

		result.OK(result.V(&RespUpload{
			MediaID:  mediaID,
			MediaURL: internal.MediaURL(mediaID),
			FileName: fh.Filename,
			FileSize: record.Size,
		})).JSON(w, r)

		return
	}

	// 将文件指针移到文件开头
	if _, err = src.Seek(0, 0); err != nil {
		log.Error(ctx, "Error seek src", zap.Error(err))
		result.ErrSystem().JSON(w, r)

		return
	}

	mediaPath := internal.MediaFile(mediaID, fh.Filename)

	dst, err := yiigo.CreateFile(mediaPath)
	if err != nil {
		log.Error(ctx, "Error create file", zap.Error(err))
		result.ErrSystem().JSON(w, r)

		return
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		log.Error(ctx, "Error copy src to file", zap.Error(err))
		result.ErrSystem().JSON(w, r)

		return
	}

	// 如果是图片，则获取图片宽高
	imgExif, err := ParseMediaEXIF(mediaPath)
	if err != nil {
		log.Error(ctx, "Error parse EXIF", zap.Error(err))
		result.ErrSystem().JSON(w, r)

		return
	}

	// 创建文件
	record, err = db.Client().File.Create().
		SetFingerprint(fingerprint).
		SetSize(fh.Size).
		SetFormat(imgExif.Format).
		SetWidth(imgExif.Width).
		SetHeight(imgExif.Height).
		SetOrientation(imgExif.Orientation).
		Save(ctx)
	if err != nil {
		// DB失败，删除文件
		os.RemoveAll(mediaPath)
		log.Error(ctx, "Error create file", zap.Error(err))
		result.ErrSystem().JSON(w, r)
		return
	}

	// 创建Media
	_, err = db.Client().Media.Create().SetMediaID(mediaID).SetFileName(fh.Filename).SetFileID(record.ID).Save(ctx)
	if err != nil {
		log.Error(ctx, "Error create media", zap.Error(err))
		result.ErrSystem().JSON(w, r)
		return
	}

	result.OK(result.V(&RespUpload{
		MediaURL: internal.MediaURL(mediaID),
		MediaID:  mediaID,
		FileName: fh.Filename,
		FileSize: fh.Size,
	})).JSON(w, r)
}
