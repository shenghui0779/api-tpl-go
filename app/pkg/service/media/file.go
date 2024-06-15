package media

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"api/app/ent"
	"api/app/ent/media"
	"api/lib/log"
	"api/lib/result"

	"github.com/disintegration/imaging"
	"github.com/go-chi/chi/v5"
	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"
)

type MediaInfo struct {
	MediaID     string `json:"media_id"`
	Name        string `json:"name"`
	Size        int64  `json:"size"`
	SizeStr     string `json:"size_str"`
	Format      string `json:"format,omitempty"`      // 图片独有
	Width       int    `json:"width,omitempty"`       // 图片独有
	Height      int    `json:"height,omitempty"`      // 图片独有
	Orientation string `json:"orientation,omitempty"` // 图片独有
	Duration    int    `json:"duration,omitempty"`    // 视频独有
	Fingerprint string `json:"fingerprint"`
	ModTime     string `json:"mod_time"`
}

func File(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	mediaID := chi.URLParam(r, "mediaID")

	record, err := ent.DB.Media.Query().WithFile(func(fq *ent.FileQuery) {
		fq.Unique(false)
	}).Unique(false).Where(media.MediaID(mediaID)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			result.ErrNotFound(result.M("document not found")).JSON(w, r)
		} else {
			log.Error(ctx, "Error ent.Media.Query.First", zap.Error(err), zap.String("media_id", mediaID))
			result.ErrSystem().JSON(w, r)
		}

		return
	}

	// 真实文件
	file := record.Edges.File
	if file == nil {
		result.ErrNotFound(result.M("document not found")).JSON(w, r)
		return
	}

	query := r.URL.Query()

	// 文件信息
	if query.Has("info") {
		result.OK(result.V(&MediaInfo{
			MediaID:     mediaID,
			Name:        record.FileName,
			Size:        file.Size,
			SizeStr:     yiigo.Quantity(file.Size).String(),
			Format:      file.Format,
			Width:       file.Width,
			Height:      file.Height,
			Orientation: Orientation(file.Orientation).String(),
			Duration:    file.Duration,
			Fingerprint: file.Fingerprint,
			ModTime:     record.CreatedAt.Format(time.DateTime),
		})).JSON(w, r)

		return
	}

	format, _ := imaging.FormatFromFilename(record.FileName)
	filePath := MediaFile(file.Fingerprint, record.FileName)

	// 非图片，直接下载
	if format < 0 {
		w.Header().Set("content-disposition", "attachment; filename="+record.FileName)
		http.ServeFile(w, r, filePath)

		return
	}

	// 缩略图
	if v := query.Get("thumbnail"); len(v) != 0 {
		thumbnail(w, r, filePath, v)
		return
	}

	// 图片标注
	if v := query["label"]; len(v) != 0 {
		label(w, r, filePath, v)
		return
	}

	// 图片裁切
	if v := query.Get("crop"); len(v) != 0 {
		crop(w, r, filePath, v)
		return
	}

	http.ServeFile(w, r, filePath)
}

func thumbnail(w http.ResponseWriter, r *http.Request, mediaFile, thumbnail string) {
	ctx := r.Context()

	rect := new(Rect)
	quality := 80

	items := strings.Split(thumbnail, "/")
	count := len(items)

	for i := 0; i < count; i++ {
		next := i + 1
		if next >= count {
			break
		}

		switch items[i] {
		case "w":
			rect.W, _ = strconv.Atoi(items[next])
		case "h":
			rect.H, _ = strconv.Atoi(items[next])
		case "q":
			if q, _ := strconv.Atoi(items[next]); q > 0 && q <= 100 {
				quality = q
			}
		}
	}

	if rect.W < 0 || rect.H < 0 {
		result.ErrParams(result.M("Error thumbnail params")).JSON(w, r)
		return
	}

	if err := ImageThumbnail(w, mediaFile, rect, imaging.JPEGQuality(quality)); err != nil {
		log.Error(ctx, "Error image thumbnail", zap.Error(err), zap.String("file", mediaFile))
		result.ErrSystem().JSON(w, r)
	}
}

func label(w http.ResponseWriter, r *http.Request, mediaFile string, labels []string) {
	ctx := r.Context()

	rects := make([]*Rect, 0, len(labels))

	for _, v := range labels {
		items := strings.Split(v, "/")
		count := len(items)

		rect := new(Rect)

		for i := 0; i < count; i++ {
			next := i + 1
			if next >= count {
				break
			}

			switch items[i] {
			case "x":
				rect.X, _ = strconv.Atoi(items[next])
			case "y":
				rect.Y, _ = strconv.Atoi(items[next])
			case "w":
				rect.W, _ = strconv.Atoi(items[next])
			case "h":
				rect.H, _ = strconv.Atoi(items[next])
			}
		}

		if rect.X < 0 || rect.Y < 0 || rect.W <= 0 || rect.H <= 0 {
			result.ErrParams(result.M("Error label params")).JSON(w, r)
			return
		}

		rects = append(rects, rect)
	}

	if err := ImageLabel(w, mediaFile, rects, imaging.JPEGQuality(100)); err != nil {
		log.Error(ctx, "Error label image", zap.Error(err), zap.String("file", mediaFile))
		result.ErrSystem().JSON(w, r)
	}
}

func crop(w http.ResponseWriter, r *http.Request, mediaFile, crop string) {
	ctx := r.Context()

	rect := new(Rect)

	items := strings.Split(crop, "/")
	count := len(items)

	for i := 0; i < count; i++ {
		next := i + 1
		if next >= count {
			break
		}

		switch items[i] {
		case "x":
			rect.X, _ = strconv.Atoi(items[next])
		case "y":
			rect.Y, _ = strconv.Atoi(items[next])
		case "w":
			rect.W, _ = strconv.Atoi(items[next])
		case "h":
			rect.H, _ = strconv.Atoi(items[next])
		}
	}

	if rect.X < 0 || rect.Y < 0 || rect.W <= 0 || rect.H <= 0 {
		result.ErrParams(result.M("Error crop params")).JSON(w, r)
		return
	}

	if err := ImageCrop(w, mediaFile, rect, imaging.JPEGQuality(100)); err != nil {
		log.Error(ctx, "Error crop image", zap.Error(err), zap.String("file", mediaFile))
		result.ErrSystem().JSON(w, r)
	}
}
