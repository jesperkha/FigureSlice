package img

import (
	"image"
	"image/color"
	"image/draw"
	"math"
)

// Returns new empty mask
func NewEmptyMask(rect image.Rectangle) *Mask {
	mask := image.NewRGBA(rect)
	return &Mask{mask}
}

// Get new mask with shapes draw to it
func GetMask(rect image.Rectangle, shapes []Shape) *Mask {
	mask := NewEmptyMask(rect)
	for _, shape := range shapes {
		switch (shape.Type) {
			case Rectangle: mask.drawRect(shape)
			case Circle:    mask.drawCircle(shape)
		}
	}
	
	return mask
}

// Draws image with mask to new empty image
func GetMaskedImage(src image.Image, mask *Mask) image.Image {
	rect := src.Bounds()
	zero := image.Pt(0, 0)
	canvas := image.NewRGBA(rect)
	draw.DrawMask(canvas, rect, src, zero, mask, zero, draw.Over)
	return canvas
}

type Mask struct {
	draw.Image
}

// Draws rect shape to mask
func (m *Mask) drawRect(shape Shape) {
	point := image.Point(shape.Pos)

	// Fill rect
	shapeRect := image.NewRGBA(shape.rect())
	col := color.RGBA{0, 0, 0, uint8(shape.Opacity)}
	draw.Draw(shapeRect, shapeRect.Bounds(), &image.Uniform{col}, image.Pt(0, 0), draw.Over)

	// Draw rect to incomplete mask
	draw.Draw(m, shapeRect.Bounds(), shapeRect, point, draw.Over)
}

// Draws circle shape to mask
func (m *Mask) drawCircle(shape Shape) {
	for x := 0; x < m.Bounds().Max.X; x++ {
		for y := 0; y < m.Bounds().Max.Y; y++ {
			// Get dist from center
			dx := shape.Pos.X - x
			dy := shape.Pos.Y - y
			c2 := float64(dx * dx + dy * dy)
			dist := int(math.Sqrt(c2))

			if dist < shape.Size.X {
				m.Set(x, y, color.RGBA{0, 0, 0, uint8(shape.Opacity)})
			}
		}
	}
}