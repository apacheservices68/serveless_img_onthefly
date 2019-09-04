package main

import (
	"io/ioutil"
	"testing"
)

func TestImageGifResize(t *testing.T) {
	opts := ImageOptions{Width: 300, Height: 300}
	buf, _ := ioutil.ReadAll(readFile("car.gif"))

	img, err := Resize(buf, opts)
	if err != nil {
		t.Errorf("Cannot process image: %s", err)
	}
	if img.Mime != "image/gif" {
		t.Error("Invalid image MIME type")
	}
	// 550x740 -> 222x300
	if assertSize(img.Body, 300, 300) != nil {
		t.Errorf("Invalid image size, expected: %dx%d", opts.Width, opts.Height)
	}
}

func TestImageJpegResize(t *testing.T) {
	opts := ImageOptions{Width: 300, Height: 300}
	buf, _ := ioutil.ReadAll(readFile("imaginary.jpg"))

	img, err := Resize(buf, opts)
	if err != nil {
		t.Errorf("Cannot process image: %s", err)
	}
	if img.Mime != "image/jpeg" {
		t.Error("Invalid image MIME type")
	}
	// 550x740 -> 222x300
	if assertSize(img.Body, 300, 300) != nil {
		t.Errorf("Invalid image size, expected: %dx%d", opts.Width, opts.Height)
	}
}

func TestImagePngResize(t *testing.T) {
	opts := ImageOptions{Width: 300, Height: 300}
	buf, _ := ioutil.ReadAll(readFile("test.png"))

	img, err := Resize(buf, opts)
	if err != nil {
		t.Errorf("Cannot process image: %s", err)
	}
	if img.Mime != "image/png" {
		t.Error("Invalid image MIME type")
	}
	// 550x740 -> 222x300
	if assertSize(img.Body, 300, 300) != nil {
		t.Errorf("Invalid image size, expected: %dx%d", opts.Width, opts.Height)
	}
}
