package main

import (
	"log"
	"net/http"

	"github.com/jesperkha/ImageMasker/client"
)

func main() {
	http.HandleFunc("/", client.HandleRequest)

	var filePrefixes = map[string]string {
		"/js/": "./client/js/",
		"/css/": "./client/css/",
	}

	for key, value := range filePrefixes {
		http.Handle(key, http.StripPrefix(key, http.FileServer(http.Dir(value))))
	}

	log.Print("Open on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}