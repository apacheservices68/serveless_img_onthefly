package main

import (
	"bytes"
	"errors"
	"image"
	"tto-resize/lilliput"

	"github.com/h2non/bimg"
	"github.com/muesli/smartcrop"
	"github.com/muesli/smartcrop/nfnt"
)

// OperationsMap defines the allowed image transformation operations listed by name.
var OperationsMap = map[string]Operation{
	"origin":  Origin,
	"cache":   Cache,
	"crop":    Crop,
	"zoom":    Zoom,
	"fit":     Fit,
	"thumb_w": thumbWidth,
	"thumb_h": thumbHeight,
}

// Image stores an image binary buffer and its MIME type
type Image struct {
	Body []byte
	Mime string
}

// Operation implements an image transformation runnable interface
type Operation func([]byte, ImageOptions) (Image, error)

// Run performs the image transformation
func (o Operation) Run(buf []byte, opts ImageOptions) (Image, error) {
	return o(buf, opts)
}

// ImageInfo represents an image details and additional metadata
type ImageInfo struct {
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	Type        string `json:"type"`
	Space       string `json:"space"`
	Alpha       bool   `json:"hasAlpha"`
	Profile     bool   `json:"hasProfile"`
	Channels    int    `json:"channels"`
	Orientation int    `json:"orientation"`
}

// Origin origin
func Origin(buf []byte, o ImageOptions) (Image, error) {
	opts := BimgOptions(o)

	return Process(buf, opts)
}

// Cache cache
func Cache(buf []byte, o ImageOptions) (Image, error) {
	opts := BimgOptions(o)

	return Process(buf, opts)
}

// Resize resize
func Resize(buf []byte, o ImageOptions) (Image, error) {
	if o.Width == 0 && o.Height == 0 {
		return Image{}, NewError("Missing required param: height or width", BadRequest)
	}

	opts := BimgOptions(o)
	// opts.Embed = true

	dims, err := bimg.Size(buf)
	if err != nil {
		return Image{}, err
	}

	if dims.Width*o.Height > o.Width*dims.Height {
		if dims.Width != 0 {
			o.Height = o.Width * dims.Height / dims.Width
		}
	} else {
		if dims.Height != 0 {
			o.Width = o.Height * dims.Width / dims.Height
		}
	}

	if o.NoCrop == true {
		opts.Crop = true
		opts.Embed = false
	}

	return Process(buf, opts)
}

// Crop crop
func Crop(buf []byte, o ImageOptions) (Image, error) {
	if o.Width == 0 && o.Height == 0 {
		return Image{}, NewError("Missing required param: height or width", BadRequest)
	}

	opts := BimgOptions(o)
	opts.Crop = true

	return Process(buf, opts)
}

// Fit fit
func Fit(buf []byte, o ImageOptions) (Image, error) {
	if o.Width == 0 || o.Height == 0 {
		return Image{}, NewError("Missing required params: height, width", BadRequest)
	}

	dims, err := bimg.Size(buf)
	if err != nil {
		return Image{}, err
	}

	if dims.Width*o.Height > o.Width*dims.Height {
		if dims.Width != 0 {
			o.Height = o.Width * dims.Height / dims.Width
		}
	} else {
		if dims.Height != 0 {
			o.Width = o.Height * dims.Width / dims.Height
		}
	}

	opts := BimgOptions(o)
	opts.Embed = true
	opts.Crop = o.NoCrop

	return Process(buf, opts)
}

func thumbWidth(buf []byte, o ImageOptions) (Image, error) {
	if o.Width == 0 {
		return Image{}, NewError("Missing required param: height or width", BadRequest)
	}

	opts := BimgOptions(o)
	opts.Force = true

	return Process(buf, opts)
}

func thumbHeight(buf []byte, o ImageOptions) (Image, error) {
	if o.Height == 0 {
		return Image{}, NewError("Missing required param: height or width", BadRequest)
	}

	opts := BimgOptions(o)
	opts.Force = true

	return Process(buf, opts)
}

// Zoom zoom
func Zoom(buf []byte, o ImageOptions) (Image, error) {
	if o.Width == 0 && o.Height == 0 {
		return Image{}, NewError("Missing required param: height or width", BadRequest)
	}

	opts := BimgOptions(o)
	// opts.Crop = true
	// opts.Gravity = bimg.GravitySmart

	img, _, _ := image.Decode(bytes.NewReader(buf))
	analyzer := smartcrop.NewAnalyzer(nfnt.NewDefaultResizer())
	sizeCrop, _ := analyzer.FindBestCrop(img, o.Width, o.Height)

	opts.Top = sizeCrop.Min.Y
	opts.Left = sizeCrop.Min.X

	opts.AreaWidth = o.Width
	opts.AreaHeight = o.Height

	opts.Width = 0
	opts.Height = 0

	return Process(buf, opts)
}

// Process process
func Process(buf []byte, opts bimg.Options) (out Image, err error) {
	defer func() {
		if r := recover(); r != nil {
			switch value := r.(type) {
			case error:
				err = value
			case string:
				err = errors.New(value)
			default:
				err = errors.New("libvips internal error")
			}
			out = Image{}
		}
	}()

	mime := GetImageMimeType(bimg.DetermineImageType(buf))

	if mime == "image/gif" {
		decoder, err := lilliput.NewDecoder(buf)

		if err != nil {
			return Image{}, err
		}
		defer decoder.Close()

		ops := lilliput.NewImageOps(8192)
		defer ops.Close()

		// create a buffer to store the output image, 50MB in this case
		outputImg := make([]byte, 50*1024*1024)

		opts := &lilliput.ImageOptions{
			FileType:             ".gif",
			Width:                opts.Width,
			Height:               opts.Height,
			ResizeMethod:         lilliput.ImageOpsFit,
			NormalizeOrientation: true,
			// EncodeOptions:        EncodeOptions[outputType],
		}

		if opts.Height == 0 {
			opts.ResizeMethod = lilliput.ImageOpsNoResize
		}

		outputImg, err = ops.Transform(decoder, opts, outputImg)
		if err != nil {
			return Image{}, err
		}

		return Image{Body: outputImg, Mime: mime}, nil
	}

	buf, err = bimg.Resize(buf, opts)
	if err != nil {
		return Image{}, err
	}

	return Image{Body: buf, Mime: mime}, nil
}
