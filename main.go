package main

import (
	"log"

	"github.com/jesperkha/ImageMasker/img"
)

func main() {
	src, err := img.LoadImage("./image.png")
	if err != nil {
		log.Fatal(err)
	}

	mask := img.GetMask(src.Bounds(), []img.Shape{})
	err = img.WriteImage("new.png", img.GetMaskedImage(src, mask))
	if err != nil {
		log.Fatal(err)
	}
}