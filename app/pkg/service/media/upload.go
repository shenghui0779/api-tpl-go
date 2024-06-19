package media

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"net/http"

	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"

	"api/app/ent"
	"api/app/ent/file"
	"api/lib"
	"api/lib/log"
	"api/lib/result"
)

type RespUpload struct {
	MediaID    string `json:"media_id"`
	MediaURL   string `json:"media_url"`
	FileName   string `json:"file_name"`
	FileSize   int64  `json:"file_size"`
	FileFormat string `json:"file_format"`
	Duration   string `json:"duration"`
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

	record, err := ent.DB.File.Query().Unique(false).Select(file.FieldID, file.FieldSize, file.FieldFormat, file.FieldDuration).Where(file.Fingerprint(fingerprint)).First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		log.Error(ctx, "Error ent.File.Query", zap.Error(err), zap.String("fingerprint", fingerprint))
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
			MediaID:    mediaID,
			MediaURL:   MediaURL(mediaID),
			FileName:   fh.Filename,
			FileSize:   record.Size,
			FileFormat: record.Format,
			Duration:   record.Duration,
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
	exif, err := lib.ParseMediaEXIF(mediaPath)
	if err != nil {
		log.Error(ctx, "Error ParseMediaEXIF", zap.Error(err), zap.String("path", mediaPath))
		result.ErrSystem().JSON(w, r)
		return
	}

	// 创建文件
	err = createMedia(ctx, mediaID, mediaPath, fh.Filename, fingerprint, exif)
	if err != nil {
		log.Error(ctx, "Error createMedia", zap.Error(err))
		result.ErrSystem().JSON(w, r)
		return
	}

	result.OK(result.V(&RespUpload{
		MediaURL:   MediaURL(mediaID),
		MediaID:    mediaID,
		FileName:   fh.Filename,
		FileSize:   fh.Size,
		FileFormat: exif.Format,
		Duration:   exif.Duration,
	})).JSON(w, r)
}
