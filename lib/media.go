package lib

import (
	"fmt"
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
	"github.com/tidwall/gjson"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

const MediaThumbnailWidth = 200

type Orientation int

func (o Orientation) String() string {
	s := ""

	switch o {
	case TopLeft:
		s = "Top-Left"
	case TopRight:
		s = "Top-Right"
	case BottomRight:
		s = "Bottom-Right"
	case BottomLeft:
		s = "Bottom-Left"
	case LeftTop:
		s = "Left-Top"
	case RightTop:
		s = "Right-Top"
	case RightBottom:
		s = "Right-Bottom"
	case LeftBottom:
		s = "Left-Bottom"
	}

	return s
}

const (
	TopLeft     Orientation = 1
	TopRight    Orientation = 2
	BottomRight Orientation = 3
	BottomLeft  Orientation = 4
	LeftTop     Orientation = 5
	RightTop    Orientation = 6
	RightBottom Orientation = 7
	LeftBottom  Orientation = 8
)

type MediaEXIF struct {
	Size        int64
	Format      string
	Width       int
	Height      int
	Orientation string
	Longitude   decimal.Decimal
	Latitude    decimal.Decimal
	Duration    string
}

func ParseMediaEXIF(filename string) (*MediaEXIF, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("os.Open: %w", err)
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		return nil, fmt.Errorf("os.File.Stat: %w", err)
	}
	data := &MediaEXIF{
		Size: stat.Size(),
	}

	format, _ := imaging.FormatFromFilename(filename)
	if format < 0 {
		data.Format = strings.ToUpper(strings.TrimPrefix(filepath.Ext(filename), "."))
		output, err := ffmpeg.Probe(filename)
		if err == nil {
			data.Duration = gjson.Get(output, "format.duration").String()
		}
		return data, nil
	}

	// 图片格式
	data.Format = format.String()

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
				data.Orientation = Orientation(v).String()
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
		rect.W = MediaThumbnailWidth
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
		rect.W = MediaThumbnailWidth
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
