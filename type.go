package main

import (
	"strings"

	"github.com/h2non/bimg"
)

// ExtractImageTypeFromMime returns the MIME image type.
func ExtractImageTypeFromMime(mime string) string {
	mime = strings.Split(mime, ";")[0]
	parts := strings.Split(mime, "/")
	if len(parts) < 2 {
		return ""
	}
	name := strings.Split(parts[1], "+")[0]
	return strings.ToLower(name)
}

// IsImageMimeTypeSupported returns true if the image MIME
// type is supported by bimg.
func IsImageMimeTypeSupported(mime string) bool {
	format := ExtractImageTypeFromMime(mime)

	// Some payloads may expose the MIME type for SVG as text/xml
	if format == "xml" {
		format = "svg"
	}

	return bimg.IsTypeNameSupported(format)
}

// ImageType returns the image type based on the given image type alias.
func ImageType(name string) bimg.ImageType {
	ext := strings.ToLower(name)
	if ext == "jpeg" {
		return bimg.JPEG
	}
	if ext == "png" {
		return bimg.PNG
	}
	if ext == "gif" {
		return bimg.GIF
	}
	return bimg.UNKNOWN
}

// GetImageMimeType returns the MIME type based on the given image type code.
func GetImageMimeType(code bimg.ImageType) string {
	if code == bimg.PNG {
		return "image/png"
	}
	if code == bimg.GIF {
		return "image/gif"
	}
	return "image/jpeg"
}
