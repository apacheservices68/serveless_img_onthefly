package main

import (
	"encoding/json"
	"mime"
	"net/http"
	"strconv"
	"strings"
)

func indexController(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorReply(r, w, ErrNotFound, ServerOptions{})
		return
	}

	body, _ := json.Marshal(CurrentVersions)
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func healthController(w http.ResponseWriter, r *http.Request) {
	health := GetHealthStats()
	body, _ := json.Marshal(health)
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func imageController(o ServerOptions, operation Operation) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		var imageSource = MatchSource(req)
		if imageSource == nil {
			ErrorReply(req, w, ErrMissingImageSource, o)
			return
		}

		buf, err := imageSource.GetImage(req)
		if err != nil {
			ErrorReply(req, w, NewError(err.Error(), BadRequest), o)
			return
		}

		if len(buf) == 0 {
			ErrorReply(req, w, ErrEmptyBody, o)
			return
		}

		imageHandler(w, req, buf, operation, o)
	}
}

func determineAcceptMimeType(accept string) string {
	for _, v := range strings.Split(accept, ",") {
		mediatype, _, _ := mime.ParseMediaType(v)
		if mediatype == "image/png" {
			return "png"
		} else if mediatype == "image/gif" {
			return "gif"
		} else if mediatype == "image/jpeg" {
			return "jpeg"
		}
	}
	// default
	return ""
}

func imageHandler(w http.ResponseWriter, r *http.Request, buf []byte, Operation Operation, o ServerOptions) {
	// Infer the body MIME type via mimesniff algorithm
	mimeType := http.DetectContentType(buf)

	// Finally check if image MIME type is supported
	if IsImageMimeTypeSupported(mimeType) == false {
		ErrorReply(r, w, ErrUnsupportedMedia, o)
		return
	}

	opts, err := readParams(r.URL.Path, r.URL.Query(), o.Profiles)
	if err != nil {
		ErrorReply(r, w, NewError("Error params image: "+err.Error(), BadRequest), o)
		return
	}

	image, err := Operation.Run(buf, opts)
	if err != nil {
		ErrorReply(r, w, NewError("Error while processing the image: "+err.Error(), BadRequest), o)
		return
	}

	// Expose Content-Length response header
	w.Header().Set("Content-Length", strconv.Itoa(len(image.Body)))
	w.Header().Set("Content-Type", image.Mime)
	w.Write(image.Body)
}
