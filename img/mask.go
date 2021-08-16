package img

import (
	"image"
	"image/draw"
)

type Mask struct {
	draw.Image
}

func NewMask(rect image.Rectangle) *Mask {
	mask := image.NewRGBA(rect)
	return &Mask{mask}
}

// Draw single shape to mask
func (m *Mask) DrawShape(shape Shape) {
	m.DrawShapes([]Shape{shape})
}

// Draw multiple shapes to mask. Shapes will overlap eachother
func (m *Mask) DrawShapes(shapes []Shape) {
	for _, shape := range shapes {
		switch (shape.Type) {
			case Rectangle: m.drawRect(shape)
			case Circle:    m.drawCircle(shape)
			case Polygon:   m.drawPolygon(shape)
		}
	}
}

// Draws given image to a buffer using m as a mask. Returns said buffer
func (m *Mask) DrawToImage(img image.Image) image.Image {
	rect := img.Bounds()
	zero := image.Pt(0, 0)
	buffer := image.NewRGBA(rect)
	draw.DrawMask(buffer, rect, img, zero, m, zero, draw.Over)
	return buffer
}