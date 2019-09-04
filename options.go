package main

import "github.com/h2non/bimg"

// ImageOptions represent all the supported image transformation params as first level members
type ImageOptions struct {
	Height int
	Width  int
	NoCrop bool
}

// BimgOptions creates a new bimg compatible options struct mapping the fields properly
func BimgOptions(o ImageOptions) bimg.Options {
	opts := bimg.Options{
		Width:  o.Width,
		Height: o.Height,
		// AreaHeight:    0,
		// AreaWidth:     0,
		// Top:           0,
		// Left:          0,
		Quality:     80,
		Compression: 1,
		// Zoom:        0,
		Crop:      false,
		SmartCrop: false,
		Enlarge:   false,
		Embed:     false,
		Flip:      false,
		Flop:      false,
		Force:     false,
		// NoAutoRotate:  false,
		// NoProfile:     false,
		Interlace:     true,
		StripMetadata: true,
		Trim:          false,
		Lossless:      true,
		// Extend:        bimg.ExtendBlack,
		// Rotate:        bimg.D0,
		// Background:    bimg.Color{R: 0, G: 0, B: 0},
		Gravity: bimg.GravityCentre,
		// Watermark: bimg.Watermark{
		// 	Width:       0,
		// 	DPI:         0,
		// 	Margin:      0,
		// 	Opacity:     0,
		// 	NoReplicate: false,
		// 	Text:        "",
		// 	Font:        "",
		// 	Background:  bimg.Color{R: 0, G: 0, B: 0},
		// },
		// Type:           bimg.UNKNOWN,
		// Interpolator:   bimg.Bicubic,
		// Interpretation: bimg.InterpretationMultiband,
		// GaussianBlur: bimg.GaussianBlur{
		// 	Sigma:   0,
		// 	MinAmpl: 0,
		// },
		// Sharpen: bimg.Sharpen{
		// 	Radius: 0,
		// 	X1:     0,
		// 	Y2:     0,
		// 	Y3:     0,
		// 	M1:     0,
		// 	M2:     0,
		// },
		// Threshold: 0,
		// OutputICC: "",
	}

	return opts
}
