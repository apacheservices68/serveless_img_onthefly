package main

import (
	"errors"
	"math"
	"net/url"
	"strconv"
	"strings"
)

var allowedParams = map[string]string{
	"width":  "int",
	"height": "int",
	"nocrop": "bool",
}

func readParams(path string, query url.Values, listProfile profiles) (ImageOptions, error) {
	params := make(map[string]interface{})

	for key, kind := range allowedParams {
		param := query.Get(key)
		params[key] = parseParam(param, kind)
	}

	configs := strings.Split(path, "/")

	action := configs[1]

	params["width"] = 0
	params["height"] = 0

	if action == "ttnew" || action == "datg" || action == "ttc" || action == "sticker" {
		if configs[2] == "r" {
			return mapImageParams(params), nil
		}

		if configs[2] == "i" {
			profile := configs[3]

			width, height, NoCrop := getSizeProfile(profile, listProfile)

			params["width"] = width
			params["height"] = height
			params["nocrop"] = NoCrop

			return mapImageParams(params), nil
		}
	}

	if action == "origin" {
		return mapImageParams(params), nil
	}

	if action == "thumb_w" {
		if !IsNumeric(configs[2]) || CheckSize(configs[2]) {
			return ImageOptions{}, errors.New("wrong type size")
		}

		params["width"] = parseParam(configs[2], "int")

		return mapImageParams(params), nil
	}

	if action == "thumb_h" {
		if !IsNumeric(configs[2]) || CheckSize(configs[2]) {
			return ImageOptions{}, errors.New("wrong type size")
		}
		params["height"] = parseParam(configs[2], "int")

		return mapImageParams(params), nil
	}

	size := strings.Split(configs[2], "_")

	if len(size) == 2 {
		if !IsNumeric(size[0]) || !IsNumeric(size[1]) || CheckSize(size[0]) || CheckSize(size[1]) {
			return ImageOptions{}, errors.New("wrong type size")
		}
		params["width"] = parseParam(size[0], "int")
		params["height"] = parseParam(size[1], "int")

		return mapImageParams(params), nil
	}

	return mapImageParams(params), nil
}

// getSizeProfile return width, height, NoCrop
func getSizeProfile(profile string, listProfile profiles) (int, int, bool) {
	for _, p := range listProfile.Profile {
		if p.Name == profile {
			return p.Width, p.Height, p.NoCrop
		}
	}

	return 200, 200, true
}

// IsNumeric is numeric
func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// CheckSize check width & height below 1024
func CheckSize(s string) bool {
	number := int(parseFloat(s))

	return number > 1024
}

func readMapParams(options map[string]interface{}) ImageOptions {
	params := make(map[string]interface{})

	for key, kind := range allowedParams {
		value, ok := options[key]
		if !ok {
			// Force type defaults
			params[key] = parseParam("", kind)
			continue
		}

		// Parse non JSON primitive types that would be represented as string types
		if kind == "int" {
			if v, ok := value.(float64); ok {
				params[key] = int(v)
			}
			if v, ok := value.(int); ok {
				params[key] = v
			}
		} else {
			params[key] = value
		}
	}

	return mapImageParams(params)
}

func parseParam(param, kind string) interface{} {
	if kind == "int" {
		return parseInt(param)
	}
	if kind == "bool" {
		return parseBool(param)
	}
	return param
}

func mapImageParams(params map[string]interface{}) ImageOptions {
	return ImageOptions{
		Width:  params["width"].(int),
		Height: params["height"].(int),
		NoCrop: params["nocrop"].(bool),
	}
}

func parseInt(param string) int {
	return int(math.Floor(parseFloat(param) + 0.5))
}

func parseFloat(param string) float64 {
	val, _ := strconv.ParseFloat(param, 64)
	return math.Abs(val)
}

func parseBool(val string) bool {
	value, _ := strconv.ParseBool(val)
	return value
}
