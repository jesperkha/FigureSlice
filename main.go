package main

import (
	"log"

	"github.com/jesperkha/ImageMasker/img"
)

func main() {
	// http.HandleFunc("/", client.HandleRequest)

	// var filePrefixes = map[string]string {
	// 	"/js/": "./website/js/",
	// 	"/css/": "./website/css/",
	// }

	// for key, value := range filePrefixes {
	// 	http.Handle(key, http.StripPrefix(key, http.FileServer(http.Dir(value))))
	// }

	// log.Print("Open on port 3000")
	// log.Fatal(http.ListenAndServe(":3000", nil))

	i, err := img.LoadImage("./test2.png")
	if err != nil {
		log.Fatal(err)
	}

	newImg := img.TrimWhitespace(i)
	err = img.WriteImage("./new.png", newImg)
	if err != nil {
		log.Fatal(err)
	}
}