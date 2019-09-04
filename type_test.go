package main

import (
	"testing"

	"github.com/h2non/bimg"
)

func TestExtractImageTypeFromMime(t *testing.T) {
	files := []struct {
		mime     string
		expected string
	}{
		{"image/jpeg", "jpeg"},
		{"/png", "png"},
		{"png", ""},
		{"multipart/form-data; encoding=utf-8", "form-data"},
		{"", ""},
	}

	for _, file := range files {
		if ExtractImageTypeFromMime(file.mime) != file.expected {
			t.Fatalf("Invalid mime type: %s != %s", file.mime, file.expected)
		}
	}
}

func TestIsImageTypeSupported(t *testing.T) {
	files := []struct {
		name     string
		expected bool
	}{
		{"image/jpeg", true},
		{"image/png", true},
		{"IMAGE/JPEG", true},
		{"png", false},
		{"multipart/form-data; encoding=utf-8", false},
		{"application/json", false},
		{"image/gif", bimg.IsImageTypeSupportedByVips(bimg.GIF).Load},
		{"image/svg+xml", bimg.IsImageTypeSupportedByVips(bimg.SVG).Load},
		{"image/svg", bimg.IsImageTypeSupportedByVips(bimg.SVG).Load},
		{"image/tiff", bimg.IsImageTypeSupportedByVips(bimg.TIFF).Load},
		{"application/pdf", bimg.IsImageTypeSupportedByVips(bimg.PDF).Load},
		{"text/plain", false},
		{"blablabla", false},
		{"", false},
	}

	for _, file := range files {
		if IsImageMimeTypeSupported(file.name) != file.expected {
			t.Fatalf("Invalid type: %s != %t", file.name, file.expected)
		}
	}
}

func TestImageType(t *testing.T) {
	files := []struct {
		name     string
		expected bimg.ImageType
	}{
		{"jpeg", bimg.JPEG},
		{"png", bimg.PNG},
		{"gif", bimg.GIF},
		{"", bimg.UNKNOWN},
	}

	for _, file := range files {
		if ImageType(file.name) != file.expected {
			t.Fatalf("Invalid type: %s != %s", file.name, bimg.ImageTypes[file.expected])
		}
	}
}

func TestGetImageMimeType(t *testing.T) {
	files := []struct {
		name     bimg.ImageType
		expected string
	}{
		{bimg.JPEG, "image/jpeg"},
		{bimg.PNG, "image/png"},
		{bimg.GIF, "image/gif"},
		{bimg.UNKNOWN, "image/jpeg"},
	}

	for _, file := range files {
		if GetImageMimeType(file.name) != file.expected {
			t.Fatalf("Invalid type: %s != %s", bimg.ImageTypes[file.name], file.expected)
		}
	}
}