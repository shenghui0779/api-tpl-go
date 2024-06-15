package media

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"net/http"
	"os"

	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"

	"api/app/ent"
	"api/app/ent/file"
	"api/lib/log"
	"api/lib/result"
)

type RespUpload struct {
	MediaID    string `json:"media_id"`
	MediaURL   string `json:"media_url"`
	FileName   string `json:"file_name"`
	FileSize   int64  `json:"file_size"`
	FileFormat string `json:"file_format"`
	Duration   int64  `json:"duration"`
}

func Upload(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	src, fh, err := r.FormFile("media")
	if err != nil {
		log.Error(ctx, "Error FormFile(media)", zap.Error(err))
		result.ErrSystem().JSON(w, r)
		return
	}
	defer src.Close()

	h := md5.New()
	if _, err = io.Copy(h, src); err != nil {
		log.Error(ctx, "Error io.Copy", zap.Error(err))
		result.ErrSystem().JSON(w, r)
		return
	}
	fingerprint := hex.EncodeToString(h.Sum(nil))
	mediaID := MediaID(fingerprint)

	record, err := ent.DB.File.Query().Unique(false).Select(file.FieldID, file.FieldSize).Where(file.Fingerprint(fingerprint)).First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		log.Error(ctx, "Error ent.File.Query.First", zap.Error(err), zap.String("fingerprint", fingerprint))
		result.ErrSystem().JSON(w, r)
		return
	}
	// 文件存在，则秒传
	if record != nil {
		_, err = ent.DB.Media.Create().SetMediaID(mediaID).SetFileName(fh.Filename).SetFileID(record.ID).Save(ctx)
		if err != nil {
			log.Error(ctx, "Error ent.Media.Create", zap.Error(err))
			result.ErrSystem().JSON(w, r)
			return
		}

		result.OK(result.V(&RespUpload{
			MediaID:  mediaID,
			MediaURL: MediaURL(mediaID),
			FileName: fh.Filename,
			FileSize: record.Size,
		})).JSON(w, r)

		return
	}

	// 将文件指针移到文件开头
	if _, err = src.Seek(0, 0); err != nil {
		log.Error(ctx, "Error src.Seek(0,0)", zap.Error(err))
		result.ErrSystem().JSON(w, r)
		return
	}

	mediaPath := MediaFile(fingerprint, fh.Filename)

	dst, err := yiigo.CreateFile(mediaPath)
	if err != nil {
		log.Error(ctx, "Error yiigo.CreateFile", zap.Error(err), zap.String("path", mediaPath))
		result.ErrSystem().JSON(w, r)
		return
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		log.Error(ctx, "Error io.Copy", zap.Error(err))
		result.ErrSystem().JSON(w, r)
		return
	}

	// 如果是图片，则获取图片宽高
	exif, err := ParseMediaEXIF(mediaPath)
	if err != nil {
		log.Error(ctx, "Error ParseMediaEXIF", zap.Error(err), zap.String("path", mediaPath))
		result.ErrSystem().JSON(w, r)
		return
	}

	// 创建文件
	record, err = ent.DB.File.Create().
		SetFingerprint(fingerprint).
		SetSize(fh.Size).
		SetFormat(exif.Format).
		SetWidth(exif.Width).
		SetHeight(exif.Height).
		SetOrientation(exif.Orientation).
		Save(ctx)
	if err != nil {
		// DB失败，删除文件
		os.RemoveAll(mediaPath)
		log.Error(ctx, "Error ent.File.Create", zap.Error(err))
		result.ErrSystem().JSON(w, r)
		return
	}

	// 创建Media
	_, err = ent.DB.Media.Create().SetMediaID(mediaID).SetFileName(fh.Filename).SetFileID(record.ID).Save(ctx)
	if err != nil {
		log.Error(ctx, "Error ent.Media.Create", zap.Error(err))
		result.ErrSystem().JSON(w, r)
		return
	}

	result.OK(result.V(&RespUpload{
		MediaURL:   MediaURL(mediaID),
		MediaID:    mediaID,
		FileName:   fh.Filename,
		FileSize:   fh.Size,
		FileFormat: exif.Format,
	})).JSON(w, r)
}
