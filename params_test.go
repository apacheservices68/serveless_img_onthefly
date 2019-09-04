package main

import (
	"testing"
)

const fixture = "fixtures/large.jpg"

// func TestReadParams(t *testing.T) {
// 	q := url.Values{}
// 	q.Set("width", "100")
// 	q.Add("height", "80")

// 	path := "/zoom/100_80/ttnew/test.jpg"

// 	params, err := readParams(path, q)

// 	if err != nil {
// 		t.Error("Invalid parse params")
// 	}

// 	assert := params.Width == 100 &&
// 		params.Height == 80

// 	if assert == false {
// 		t.Error("Invalid params")
// 	}
// }

// func TestReadParamsTTNEW(t *testing.T) {
// 	q := url.Values{}
// 	q.Set("width", "480")
// 	q.Add("height", "270")

// 	path := "/ttnew/i/s480/2018/11/11/large.jpg"

// 	params, err := readParams(path, q)

// 	if err != nil {
// 		t.Error("Invalid parse params")
// 	}

// 	assert := params.Width == 480 &&
// 		params.Height == 270

// 	if assert == false {
// 		t.Error("Invalid params")
// 	}
// }

// func TestReadParamsWidth(t *testing.T) {
// 	q := url.Values{}
// 	q.Set("width", "100")

// 	path := "/thumb_w/100/ttnew/test.jpg"

// 	params, err := readParams(path, q)

// 	if err != nil {
// 		t.Error("Invalid parse params")
// 	}

// 	assert := params.Width == 100 &&
// 		params.Height == 0

// 	if assert == false {
// 		t.Error("Invalid params")
// 	}
// }

// func TestReadParamsHeight(t *testing.T) {
// 	q := url.Values{}
// 	q.Set("height", "100")

// 	path := "/thumb_h/100/ttnew/test.jpg"

// 	params, err := readParams(path, q)

// 	if err != nil {
// 		t.Error("Invalid parse params")
// 	}

// 	assert := params.Width == 0 &&
// 		params.Height == 100

// 	if assert == false {
// 		t.Error("Invalid params")
// 	}
// }

func TestIsNumeric(t *testing.T) {
	number := "100"
	if !IsNumeric(number) {
		t.Errorf("Invalid number: %s ", number)
	}
}

func TestCheckSize(t *testing.T) {
	number := "1000"
	if CheckSize(number) {
		t.Errorf("Invalid number: %s ", number)
	}
}

// func TestSizeProfile(t *testing.T) {
// 	configFile := "./tests/config.toml"
// 	var configs profiles

// 	if _, err := toml.DecodeFile(configFile, &configs); err != nil {
// 		t.Errorf("error config profile: %s", err)
// 	}

// 	profile := "s400"
// 	width, height, NoCrop := getSizeProfile(profile, configs)

// 	if width != 400 && height != 400 && NoCrop == true {
// 		t.Errorf("Invalid profile: %s ", profile)
// 	}
// }

func TestParseParam(t *testing.T) {
	intCases := []struct {
		value    string
		expected int
	}{
		{"1", 1},
		{"0100", 100},
		{"-100", 100},
		{"99.02", 99},
		{"99.9", 100},
	}

	for _, test := range intCases {
		val := parseParam(test.value, "int")
		if val != test.expected {
			t.Errorf("Invalid param: %s != %d", test.value, test.expected)
		}
	}
}

func TestReadMapParams(t *testing.T) {
	cases := []struct {
		params   map[string]interface{}
		expected ImageOptions
	}{
		{
			map[string]interface{}{
				"width": 100,
			},
			ImageOptions{
				Width: 100,
			},
		},
	}

	for _, test := range cases {
		opts := readMapParams(test.params)
		if opts.Width != test.expected.Width {
			t.Errorf("Invalid width: %d != %d", opts.Width, test.expected.Width)
		}
	}
}
