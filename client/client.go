package client

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/jesperkha/ImageMasker/img"
)

type handlerFunc func (res http.ResponseWriter, req *http.Request) (status int, err error)

var routes = map[string]handlerFunc {
	"image": handleImageRequest,
	"error": handleError,
}

func HandleRequest(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path

	if path == "/" {
		temps, err := template.ParseGlob("./website/templates/index/*.html")
		if err != nil {
			log.Fatal(err)
		}

		err = temps.ExecuteTemplate(res, "main", "")
		if err != nil {
			log.Fatal(err)
		}
	}
	
	split := strings.Split(path, "/")
	for route, handle := range routes {
		if route == split[1] {
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

func handleError(res http.ResponseWriter, req *http.Request) (status int, err error) {
	errorCode := strings.Split(req.URL.Path, "/")[2]
	temp, err := template.ParseFiles("./website/templates/error.html")
	if err != nil {
		return http.StatusInternalServerError, err
	}
	
	return http.StatusOK, temp.Execute(res, errorCode)
}

func handleImageRequest(res http.ResponseWriter, req *http.Request) (status int, err error) {
	if req.Method != "POST" {
		return http.StatusMethodNotAllowed, nil
	}

	if strings.Split(req.Header.Get("Content-Type"), ";")[0] != "multipart/form-data" {
		return http.StatusBadRequest, nil
	}

	// Handle form data
	var imgBuffer bytes.Buffer
	var shapeData []img.Shape
	if file, _, err := req.FormFile("Image"); err == nil {
		_, err = bufio.NewReader(file).WriteTo(&imgBuffer)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		
		err = json.Unmarshal([]byte(req.PostFormValue("Shapes")), &shapeData)
		if err != nil {
			return http.StatusInternalServerError, err
		}
	}
	
	// Get processed image
	rawImage, err := img.BLoadImage(imgBuffer.Bytes())
	if err != nil {
		return http.StatusInternalServerError, err
	}
	
	mask := img.NewMask(rawImage.Bounds())
	mask.DrawShapes(shapeData)
	finalImage := mask.DrawToImage(rawImage)
	
	if req.PostFormValue("Trim") == "on" {
		finalImage = img.TrimWhitespace(finalImage)
	}
	
	if byt, err := img.BWriteImage(finalImage); err == nil {
		_, err = res.Write(byt)
		return http.StatusOK, err
	}

	return http.StatusInternalServerError, err
}
