package main

import (
	"log"
	"net/http"

	"github.com/jesperkha/ImageMasker/client"
)

func main() {
	http.HandleFunc("/", client.HandleRequest)
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./client/js/"))))

	log.Print("Open on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}