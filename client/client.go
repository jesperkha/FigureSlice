package client

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jesperkha/ImageMasker/img"
)

type handlerFunc func (res http.ResponseWriter, req *http.Request) (status int, err error)

var routes = map[string]handlerFunc {
	"/image": handleImageRequest,
	"/error": handleError,
}

func handleError(res http.ResponseWriter, req *http.Request) (status int, err error) {
	// serve error.html
	return http.StatusOK, nil
}

func HandleRequest(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path

	if path == "/" {
		http.ServeFile(res, req, "./client/html/index.html")
	}

	for route, handle := range routes {
		if path == route {
			status, err := handle(res, req)
			if err != nil {
				log.Fatal(err)
			}

			if status != http.StatusOK {
				http.Redirect(res, req, fmt.Sprintf("/error/%d", status), status)
			}
		}
	}
}

func handleImageRequest(res http.ResponseWriter, req *http.Request) (status int, err error) {
	if req.Method != "POST" {
		return http.StatusMethodNotAllowed, nil
	}

	stream, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	pic, err := img.BLoadImage(stream)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// Proto
	// Get actual shaope data sent with image
	shape := img.Shape{
		Type: img.Rectangle,
		Pos: img.Vector{X:0, Y:0},
		Size: img.Vector{X:100, Y:100},
		Opacity: 255,
	}

	mask := img.GetMask(pic.Bounds(), []img.Shape{shape})
	finalImage, err := img.BWriteImage(img.GetMaskedImage(pic, mask))
	if err != nil {
		return http.StatusInternalServerError, err
	}

	_, err = res.Write(finalImage)
	return http.StatusOK, err
}
