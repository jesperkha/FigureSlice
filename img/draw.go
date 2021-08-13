package img

import (
	"image"
	"image/color"
	"image/draw"
	"math"
)

func (m *Mask) drawRect(shape Shape) {
	// Draw uniform rect to mask at shapes point
	col := color.RGBA{0, 0, 0, uint8(shape.Opacity)}
	draw.Draw(m, shape.Rect(), &image.Uniform{col}, image.Pt(0, 0), draw.Src)
}

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

func (mask *Mask) drawPolygon(shape Shape) {
	pts := shape.Points

	var (
		rx  = mask.Bounds().Dx()
		ry  = mask.Bounds().Dy()
		rx2 = 0
		ry2 = 0
	)

	// Get border rect to not check every pixel
	for _, point := range pts {
		if point.X < rx {
			rx = point.X
		}
		if point.X > rx2 {
			rx2 = point.X
		}
		if point.Y < ry {
			ry = point.Y
		}
		if point.Y > ry2 {
			ry2 = point.Y
		}
	}

	rect := image.Rect(rx, ry, rx2, ry2)
	buffer := image.NewRGBA(rect)
	col := color.RGBA{0, 0, 0, uint8(shape.Opacity)}

	for x := buffer.Bounds().Min.X; x < buffer.Bounds().Max.X; x++ {
		for y := buffer.Bounds().Min.Y; y < buffer.Bounds().Max.Y; y++ {
			collision := false
			next := 0
			for i := 0; i < len(pts); i++ {
				next = i + 1
				if next == len(pts) {
					next = 0
				}

				pt1 := pts[i]
				pt2 := pts[next]
				
				// wtf
				if ((pt1.Y > y) != (pt2.Y > y)) && (x < (pt2.X - pt1.X) * (y - pt1.Y) / (pt2.Y - pt1.Y) + pt1.X) {
					collision = !collision;
				}
			}
			
			if collision {
				buffer.Set(x, y, col)
			}
		}
	}

	draw.Draw(mask, rect, buffer, image.Pt(rx, ry), draw.Src)
}