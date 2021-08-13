package img

import (
	"bytes"
	"image"
	"image/png"
	"os"
)

type Vector struct {
	X int
	Y int
}

const (
	Rectangle = 0
	Circle    = 1
	Polygon   = 2
)

// Data comes in as array of MaskShape
type Shape struct {
	Type    int
	Opacity int
	Pos     Vector
	Size    Vector
	Points  []Vector
}

// Get image rect of shape
func (s *Shape) rect() image.Rectangle {
	// x1 y1 is origin
	// second args are second position but here a width/height is used
	// therefore the position x2 and y2 needs to be shifted by x1 and y1 to
	// represent a width/height
	return image.Rect(s.Pos.X, s.Pos.Y, s.Size.X + s.Pos.X, s.Size.Y + s.Pos.Y)
}

// Load from filename
func LoadImage(filename string) (img image.Image, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return img, err
	}
	defer file.Close()

	img, _, err = image.Decode(file)
	return img, err
}

// Load from byte array
func BLoadImage(byt []byte) (img image.Image, err error) {
	reader := bytes.NewReader(byt)
	img, _, err = image.Decode(reader)
	return img, err
}

// Write image as file
func WriteImage(path string, img image.Image) (err error) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, img)
}

// Get encoded image bytes
func BWriteImage(img image.Image) (byt []byte, err error) {
	var writer bytes.Buffer
	err = png.Encode(&writer, img)
	return writer.Bytes(), err
}