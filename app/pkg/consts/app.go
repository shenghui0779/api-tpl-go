package consts

const MediaThumbnailWidth = 200

type Orientation int

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

func (o Orientation) String() string {
	s := ""

	switch o {
	case TopLeft:
		s = "Top-left"
	case TopRight:
		s = "Top-right"
	case BottomRight:
		s = "Bottom-right"
	case BottomLeft:
		s = "Bottom-left"
	case LeftTop:
		s = "Left-top"
	case RightTop:
		s = "Right-top"
	case RightBottom:
		s = "Right-bottom"
	case LeftBottom:
		s = "Left-bottom"
	}

	return s
}
