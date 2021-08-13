package img

import (
	"image"
	"image/draw"
)

// Removes whitespace around image. Returns new image.
func TrimWhitespace(img image.Image) image.Image {
	var (
		dx = img.Bounds().Dx()
		dy = img.Bounds().Dy()

		closestLeft   = dx
		closestTop    = dy
		closestRight  = 0
		closestBottom = 0
	)

	for y := 0; y < dy; y++ {
		hasColor := false // for current row
		for x := 0; x < dx; x++ {
			_, _, _, a := img.At(x, y).RGBA()
			if a != 0 {
				hasColor = true
				if x < closestLeft {
					closestLeft = x
				}

				if x > closestRight {
					closestRight = x
				}
			} 
		}

		if hasColor {
			if y < closestTop {
				closestTop = y
			}
	
			if y > closestBottom {
				closestBottom = y
			}
		}
	}
	
	// Rect from edges of original image
	rect := image.Rect(closestLeft, closestTop, closestRight, closestBottom)

	// Draw original image to a RGBA buffer and get the sub image where the rect overlaps
	buffer := image.NewRGBA(img.Bounds())
	draw.Draw(buffer, img.Bounds(), img, image.Pt(0, 0), draw.Src)
	trimmed := buffer.SubImage(rect)

	return trimmed
}