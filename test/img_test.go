package tests

import (
	"testing"

	"github.com/jesperkha/FigureSlice/img"
)

// Generates 3 new images with each draw function in img/draw.go

func TestDraw(t *testing.T) {
	Shapes := []img.Shape {
		{
			Type: img.Rectangle,
			Pos: img.Vector{X: 50, Y: 50},
			Size: img.Vector{X: 200, Y: 100},
			Opacity: 255,
		},
		{
			Type: img.Circle,
			Pos: img.Vector{X: 100, Y: 300},
			Size: img.Vector{X: 50, Y: 0},
			Opacity: 255,
		},
		{
			Type: img.Polygon,
			Opacity: 255,
			Points: []img.Vector{
				{X: 50, Y: 500},
				{X: 200, Y: 500},
				{X: 50, Y: 700},
			},
		},
	}

	i, err := img.LoadImage("./src/base.png")
	if err != nil {
		t.Error(err.Error())
	}

	mask := img.NewMask(i.Bounds())
	mask.DrawShapes(Shapes)
	newImg := mask.DrawToImage(i)
	err = img.WriteImage("./src/result.png", newImg)
	if err != nil {
		t.Error(err.Error())
	}
}
