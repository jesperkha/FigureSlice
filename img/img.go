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

type Shape struct {
	Type    int
	Opacity int
	Pos     Vector
	Size    Vector
	Points  []Vector
}

// Get image rect of the shape
func (s *Shape) Rect() image.Rectangle {
	// x1 y1 is origin
	// second args are second position but here a width/height is used
	// therefore the position x2 and y2 needs to be shifted by x1 and y1 to
	// represent a width/height
	return image.Rect(s.Pos.X, s.Pos.Y, s.Size.X + s.Pos.X, s.Size.Y + s.Pos.Y)
}

// Load image from file
func LoadImage(path string) (img image.Image, err error) {
	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return img, err
	}
	defer file.Close()

	img, _, err = image.Decode(file)
	return img, err
}

// Load image from bytes
func BLoadImage(byt []byte) (img image.Image, err error) {
	reader := bytes.NewReader(byt)
	img, _, err = image.Decode(reader)
	return img, err
}

// Write image as file
func WriteImage(path string, img image.Image) (err error) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, img)
}

// Encode image to byte array
func BWriteImage(img image.Image) (byt []byte, err error) {
	var writer bytes.Buffer
	err = png.Encode(&writer, img)
	return writer.Bytes(), err
}