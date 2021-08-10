package client

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jesperkha/ImageMasker/img"
)

// Todo add error page and handling

type handlerFunc func (res http.ResponseWriter, req *http.Request) (status int, err error)

var routes = map[string]handlerFunc {
	"/image": handleImageRequest,
	"/error": handleError,
}

func handleError(res http.ResponseWriter, req *http.Request) (status int, err error) {
	// Serve error.html
	return http.StatusOK, nil
}

func HandleRequest(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path

	if path == "/" {
		http.ServeFile(res, req, "./website/html/index.html")
	}

	for route, handle := range routes {
		if path == route {
			status, err := handle(res, req)
			// Debug
			// Implement actual error handling
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

	// Handle form data
	var imgBuffer bytes.Buffer
	var shapeData []img.Shape
	if file, _, err := req.FormFile("Image"); err == nil {
		_, err = bufio.NewReader(file).WriteTo(&imgBuffer)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		
		err = json.Unmarshal([]byte(req.PostForm.Get("Shapes")), &shapeData)
		if err != nil {
			return http.StatusInternalServerError, err
		}
	}
	
	// Get processed image
	rawImage, err := img.BLoadImage(imgBuffer.Bytes())
	if err != nil {
		return http.StatusInternalServerError, err
	}
	
	mask := img.GetMask(rawImage.Bounds(), shapeData)
	if finalImage, err := img.BWriteImage(img.GetMaskedImage(rawImage, mask)); err == nil {
		_, err = res.Write(finalImage)
		return http.StatusOK, err
	}

	return http.StatusInternalServerError, err
}
