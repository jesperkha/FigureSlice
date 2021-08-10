package main

import (
	"log"
	"net/http"

	"github.com/jesperkha/ImageMasker/client"
)

func main() {
	http.HandleFunc("/", client.HandleRequest)

	var filePrefixes = map[string]string {
		"/js/": "./website/js/",
		"/css/": "./website/css/",
	}

	for key, value := range filePrefixes {
		http.Handle(key, http.StripPrefix(key, http.FileServer(http.Dir(value))))
	}

	log.Print("Open on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))

	// i, err := img.LoadImage("./src/test.png")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// mask := img.GetMask(i.Bounds(), []img.Shape{
	// 	{
	// 		Type:    0,
	// 		Pos:     img.Vector{X: 50, Y: 50},
	// 		Size:    img.Vector{X: 200, Y: 100},
	// 		Opacity: 255,
	// 	},
	// })

	// newImg := img.GetMaskedImage(i, mask)
	// err = img.WriteImage("./src/result.png", newImg)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}