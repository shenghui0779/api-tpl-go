package media

import (
	"image"
	"io"
	"os"
	"path/filepath"
	"strings"

	"git.sr.ht/~sbinet/gg"
	"github.com/disintegration/imaging"
	"github.com/pkg/errors"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/shopspring/decimal"

	"api/pkg/consts"
)

type MediaEXIF struct {
	Size        int64
	Format      string
	Width       int
	Height      int
	Orientation int
	Longitude   decimal.Decimal
	Latitude    decimal.Decimal
}

func ParseMediaEXIF(filename string) (*MediaEXIF, error) {
	// 非图片格式，不处理
	format, _ := imaging.FormatFromFilename(filename)
	if format < 0 {
		ext := filepath.Ext(filename)
		return &MediaEXIF{Format: strings.ToUpper(strings.TrimPrefix(ext, "."))}, nil
	}

	f, err := os.Open(filename)
	if err != nil {
		return nil, errors.Wrap(err, "open file")
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		return nil, errors.Wrap(err, "stat file")
	}

	data := &MediaEXIF{
		Size:   stat.Size(),
		Format: format.String(),
	}
	if x, err := exif.Decode(f); err == nil {
		if lat, lng, err := x.LatLong(); err == nil {
			data.Longitude = decimal.NewFromFloat(lng)
			data.Latitude = decimal.NewFromFloat(lat)
		}
		if tag, err := x.Get(exif.PixelXDimension); err == nil {
			if v, err := tag.Int(0); err == nil {
				data.Width = v
			}
		}
		if tag, err := x.Get(exif.PixelYDimension); err == nil {
			if v, err := tag.Int(0); err == nil {
				data.Height = v
			}
		}
		if tag, err := x.Get(exif.Orientation); err == nil {
			if v, err := tag.Int(0); err == nil {
				data.Orientation = v
			}
		}
	}
	if data.Width == 0 || data.Height == 0 {
		if img, err := imaging.Open(filename); err == nil {
			rect := img.Bounds()
			data.Width = rect.Dx()
			data.Height = rect.Dy()
		}
	}
	return data, nil
}

type Rect struct {
	X int
	Y int
	W int
	H int
}

func ImageThumbnail(w io.Writer, filename string, rect *Rect, options ...imaging.EncodeOption) error {
	if rect == nil || rect.W < 0 || rect.H < 0 {
		return errors.New("Error param rect")
	}

	img, err := imaging.Open(filename)
	if err != nil {
		return err
	}

	size := img.Bounds().Size()
	if rect.W == 0 && rect.H == 0 {
		rect.W = consts.MediaThumbnailWidth
		rect.H = rect.W * size.Y / size.X
	} else {
		if rect.W > size.X {
			rect.W = size.X
		}
		if rect.H > size.Y {
			rect.H = size.Y
		}
		if rect.W > 0 {
			if rect.H == 0 {
				rect.H = rect.W * size.Y / size.X
			}
		} else {
			rect.W = rect.H * size.X / size.Y
		}
	}
	thumbnail := imaging.Thumbnail(img, rect.W, rect.H, imaging.Lanczos)

	format, _ := imaging.FormatFromFilename(filename)
	return imaging.Encode(w, thumbnail, format, options...)
}

func ImageThumbnailFromReader(w io.Writer, r io.Reader, format imaging.Format, rect *Rect, options ...imaging.EncodeOption) error {
	if rect == nil || rect.W < 0 || rect.H < 0 {
		return errors.New("Error param rect")
	}

	img, err := imaging.Decode(r)
	if err != nil {
		return err
	}

	size := img.Bounds().Size()
	if rect.W == 0 && rect.H == 0 {
		rect.W = consts.MediaThumbnailWidth
		rect.H = rect.W * size.Y / size.X
	} else {
		if rect.W > size.X {
			rect.W = size.X
		}
		if rect.H > size.Y {
			rect.H = size.Y
		}
		if rect.W > 0 {
			if rect.H == 0 {
				rect.H = rect.W * size.Y / size.X
			}
		} else {
			rect.W = rect.H * size.X / size.Y
		}
	}
	thumbnail := imaging.Thumbnail(img, rect.W, rect.H, imaging.Lanczos)

	return imaging.Encode(w, thumbnail, format, options...)
}

func ImageCrop(w io.Writer, filename string, rect *Rect, options ...imaging.EncodeOption) error {
	if rect == nil || rect.X < 0 || rect.Y < 0 || rect.W <= 0 || rect.H <= 0 {
		return errors.New("Error param rect")
	}

	img, err := imaging.Open(filename)
	if err != nil {
		return err
	}
	crop := imaging.Crop(img, image.Rect(rect.X, rect.Y, rect.X+rect.W, rect.Y+rect.H))

	format, _ := imaging.FormatFromFilename(filename)

	return imaging.Encode(w, crop, format, options...)
}

func ImageCropFromReader(w io.Writer, r io.Reader, format imaging.Format, rect *Rect, options ...imaging.EncodeOption) error {
	if rect == nil || rect.X < 0 || rect.Y < 0 || rect.W < 0 || rect.H < 0 {
		return errors.New("Error param rect")
	}

	img, err := imaging.Decode(r)
	if err != nil {
		return err
	}
	crop := imaging.Crop(img, image.Rect(rect.X, rect.Y, rect.X+rect.W, rect.Y+rect.H))

	return imaging.Encode(w, crop, format, options...)
}

func ImageLabel(w io.Writer, filename string, rects []*Rect, options ...imaging.EncodeOption) error {
	img, err := imaging.Open(filename)
	if err != nil {
		return err
	}

	dc := gg.NewContextForImage(img)
	dc.SetRGB255(255, 0, 0)
	dc.SetLineWidth(8)
	for _, rect := range rects {
		if rect.X < 0 || rect.Y < 0 || rect.W <= 0 || rect.H <= 0 {
			return errors.New("Error param rects")
		}
		dc.DrawRectangle(float64(rect.X), float64(rect.Y), float64(rect.W), float64(rect.H))
	}
	dc.Stroke()

	format, _ := imaging.FormatFromFilename(filename)

	return imaging.Encode(w, dc.Image(), format, options...)
}

func ImageLabelFromReader(w io.Writer, r io.Reader, format imaging.Format, rects []*Rect, options ...imaging.EncodeOption) error {
	img, err := imaging.Decode(r)
	if err != nil {
		return err
	}

	dc := gg.NewContextForImage(img)
	dc.SetRGB255(255, 0, 0)
	dc.SetLineWidth(8)
	for _, rect := range rects {
		if rect.X < 0 || rect.Y < 0 || rect.W <= 0 || rect.H <= 0 {
			return errors.New("Error param rects")
		}
		dc.DrawRectangle(float64(rect.X), float64(rect.Y), float64(rect.W), float64(rect.H))
	}
	dc.Stroke()

	return imaging.Encode(w, dc.Image(), format, options...)
}
