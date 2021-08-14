package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jesperkha/FigureSlice/client"
)

func main() {
	http.HandleFunc("/", client.HandleRequest)

	var filePrefixes = map[string]string{
		"/js/":  "./website/",
		"/css/": "./website/css/",
		"/src/": "./website/src/",
	}

	for key, value := range filePrefixes {
		http.Handle(key, http.StripPrefix(key, http.FileServer(http.Dir(value))))
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}