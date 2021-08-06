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
)

// Data comes in as array of MaskShape
type Shape struct {
	Type    int
	Pos     Vector
	Size    Vector
	Opacity int
}

// Get image rect of shape
func (s *Shape) rect() image.Rectangle {
	return image.Rect(s.Pos.X, s.Pos.Y, s.Size.X, s.Size.Y)
}

// Load from filename
func LoadImage(filename string) (img image.Image, err error) {
	reader, err := os.Open(filename)
	if err != nil {
		return img, err
	}
	defer reader.Close()

	img, _, err = image.Decode(reader)
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
	writer, err := os.OpenFile(path, os.O_CREATE, os.ModeAppend)
	if err != nil {
		return err
	}

	err = png.Encode(writer, img)
	return err
}

// Get encoded image bytes
func BWriteImage(img image.Image) (byt []byte, err error) {
	var writer bytes.Buffer
	err = png.Encode(&writer, img)
	return writer.Bytes(), err
}