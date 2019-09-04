package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

const (
	// Unavailable unavailable
	Unavailable uint8 = iota
	// BadRequest Bad Request
	BadRequest
	// NotAllowed not allow request
	NotAllowed
	// Unsupported not support
	Unsupported
	// Unauthorized not auth
	Unauthorized
	// InternalError internal error
	InternalError
	// NotFound not found
	NotFound
	// NotImplemented not implement
	NotImplemented
	// Forbidden Forbidden
	Forbidden
)

var (
	ErrNotFound             = NewError("Not found", NotFound)
	ErrInvalidApiKey        = NewError("Invalid or missing API key", Unauthorized)
	ErrMethodNotAllowed     = NewError("Method not allowed", NotAllowed)
	ErrUnsupportedMedia     = NewError("Unsupported media type", Unsupported)
	ErrOutputFormat         = NewError("Unsupported output image format", BadRequest)
	ErrEmptyBody            = NewError("Empty image", BadRequest)
	ErrMissingParamFile     = NewError("Missing required param: file", BadRequest)
	ErrInvalidFilePath      = NewError("Invalid file path", BadRequest)
	ErrInvalidImageURL      = NewError("Invalid image URL", BadRequest)
	ErrMissingImageSource   = NewError("Cannot process the image due to missing or invalid params", BadRequest)
	ErrNotImplemented       = NewError("Not implemented endpoint", NotImplemented)
	ErrInvalidURLSignature  = NewError("Invalid URL signature", BadRequest)
	ErrURLSignatureMismatch = NewError("URL signature mismatch", Forbidden)
)

type Error struct {
	Message string `json:"message,omitempty"`
	Code    uint8  `json:"code"`
}

func (e Error) JSON() []byte {
	buf, _ := json.Marshal(e)
	return buf
}

func (e Error) Error() string {
	return e.Message
}

func (e Error) HTTPCode() int {
	if e.Code == BadRequest {
		return http.StatusBadRequest
	}
	if e.Code == NotAllowed {
		return http.StatusMethodNotAllowed
	}
	if e.Code == Unsupported {
		return http.StatusUnsupportedMediaType
	}
	if e.Code == InternalError {
		return http.StatusInternalServerError
	}
	if e.Code == Unauthorized {
		return http.StatusUnauthorized
	}
	if e.Code == NotFound {
		return http.StatusNotFound
	}
	if e.Code == NotImplemented {
		return http.StatusNotImplemented
	}
	if e.Code == Forbidden {
		return http.StatusForbidden
	}
	return http.StatusServiceUnavailable
}

func NewError(err string, code uint8) Error {
	err = strings.Replace(err, "\n", "", -1)
	return Error{err, code}
}

func ErrorReply(req *http.Request, w http.ResponseWriter, err Error, o ServerOptions) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.HTTPCode())
	w.Write(err.JSON())
	return err
}
