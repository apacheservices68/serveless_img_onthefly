package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/h2non/bimg"
)

func TestIndex(t *testing.T) {
	ts := testServer(indexController)
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != 200 {
		t.Fatalf("Invalid response status: %s", res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	if strings.Contains(string(body), "tto") == false {
		t.Fatalf("Invalid body response: %s", body)
	}
}

func TestCrop(t *testing.T) {
	opts := ServerOptions{
		Origin: "./tests/origin",
	}
	fn := ImageMiddleware(opts)(Crop)
	LoadSources(opts)

	ts := httptest.NewServer(fn)
	url := ts.URL + "/crop/200_250/ttnew/large.jpg"
	defer ts.Close()

	res, err := http.Get(url)
	if err != nil {
		t.Fatal("Cannot perform the request")
	}

	if res.StatusCode != 200 {
		t.Fatalf("Invalid response status: %s", res.Status)
	}

	if res.Header.Get("Content-Length") == "" {
		t.Fatal("Empty content length response")
	}

	image, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	if len(image) == 0 {
		t.Fatalf("Empty response body")
	}

	err = assertSize(image, 200, 250)
	if err != nil {
		t.Error(err)
	}

	if bimg.DetermineImageTypeName(image) != "jpeg" {
		t.Fatalf("Invalid image type")
	}
}

// func TestResize(t *testing.T) {
// 	opts := ServerOptions{
// 		Origin: "./tests/ttnew",
// 	}
// 	fn := ImageMiddleware(opts)(Zoom)
// 	LoadSources(opts)

// 	ts := httptest.NewServer(fn)
// 	url := ts.URL + "/ttnew/i/s640/2018/11/11/large.jpg"
// 	defer ts.Close()

// 	res, err := http.Get(url)
// 	if err != nil {
// 		t.Fatal("Cannot perform the request")
// 	}

// 	if res.StatusCode != 200 {
// 		t.Fatalf("Invalid response status: %s", res.Status)
// 	}

// 	image, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if len(image) == 0 {
// 		t.Fatalf("Empty response body")
// 	}

// 	err = assertSize(image, 640, 360)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	if bimg.DetermineImageTypeName(image) != "jpeg" {
// 		t.Fatalf("Invalid image type")
// 	}
// }

func TestZoom(t *testing.T) {
	opts := ServerOptions{
		Origin: "./tests/origin",
	}
	fn := ImageMiddleware(opts)(Zoom)
	LoadSources(opts)

	ts := httptest.NewServer(fn)
	url := ts.URL + "/zoom/200_250/ttnew/large.jpg"
	defer ts.Close()

	res, err := http.Get(url)
	if err != nil {
		t.Fatal("Cannot perform the request")
	}

	if res.StatusCode != 200 {
		t.Fatalf("Invalid response status: %s", res.Status)
	}

	if res.Header.Get("Content-Length") == "" {
		t.Fatal("Empty content length response")
	}

	image, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	if len(image) == 0 {
		t.Fatalf("Empty response body")
	}

	err = assertSize(image, 200, 250)
	if err != nil {
		t.Error(err)
	}

	if bimg.DetermineImageTypeName(image) != "jpeg" {
		t.Fatalf("Invalid image type")
	}
}

func TestFit(t *testing.T) {
	opts := ServerOptions{
		Origin: "./tests/origin",
	}
	fn := ImageMiddleware(opts)(Fit)
	LoadSources(opts)

	ts := httptest.NewServer(fn)
	url := ts.URL + "/fit/200_250/ttnew/large.jpg"
	defer ts.Close()

	res, err := http.Get(url)
	if err != nil {
		t.Fatal("Cannot perform the request")
	}

	if res.StatusCode != 200 {
		t.Fatalf("Invalid response status: %s", res.Status)
	}

	image, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	if len(image) == 0 {
		t.Fatalf("Empty response body")
	}

	err = assertSize(image, 166, 250)
	if err != nil {
		t.Error(err)
	}

	if bimg.DetermineImageTypeName(image) != "jpeg" {
		t.Fatalf("Invalid image type")
	}
}

func TestThumbWidth(t *testing.T) {
	opts := ServerOptions{
		Origin: "./tests/origin",
	}
	fn := ImageMiddleware(opts)(thumbWidth)
	LoadSources(opts)

	ts := httptest.NewServer(fn)
	url := ts.URL + "/thumb_w/200/ttnew/large.jpg"
	defer ts.Close()

	res, err := http.Get(url)
	if err != nil {
		t.Fatal("Cannot perform the request")
	}

	if res.StatusCode != 200 {
		t.Fatalf("Invalid response status: %s", res.Status)
	}

	image, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	if len(image) == 0 {
		t.Fatalf("Empty response body")
	}

	err = assertSize(image, 200, 300)
	if err != nil {
		t.Error(err)
	}

	if bimg.DetermineImageTypeName(image) != "jpeg" {
		t.Fatalf("Invalid image type")
	}
}

func TestThumbHeight(t *testing.T) {
	opts := ServerOptions{
		Origin: "./tests/origin",
	}
	fn := ImageMiddleware(opts)(thumbHeight)
	LoadSources(opts)

	ts := httptest.NewServer(fn)
	url := ts.URL + "/thumb_h/300/ttnew/large.jpg"
	defer ts.Close()

	res, err := http.Get(url)
	if err != nil {
		t.Fatal("Cannot perform the request")
	}

	if res.StatusCode != 200 {
		t.Fatalf("Invalid response status: %s", res.Status)
	}

	image, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	if len(image) == 0 {
		t.Fatalf("Empty response body")
	}

	err = assertSize(image, 200, 300)
	if err != nil {
		t.Error(err)
	}

	if bimg.DetermineImageTypeName(image) != "jpeg" {
		t.Fatalf("Invalid image type")
	}
}

func TestOriginDirectory(t *testing.T) {
	opts := ServerOptions{
		Origin: "./tests/origin",
	}
	fn := ImageMiddleware(opts)(Origin)
	LoadSources(opts)

	ts := httptest.NewServer(fn)
	url := ts.URL + "/origin/ttnew/large.jpg"
	defer ts.Close()

	res, err := http.Get(url)
	if err != nil {
		t.Fatal("Cannot perform the request")
	}
	if res.StatusCode != 200 {
		t.Fatalf("Invalid response status: %d", res.StatusCode)
	}

	image, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	if len(image) == 0 {
		t.Fatalf("Empty response body")
	}

	err = assertSize(image, 700, 1050)
	if err != nil {
		t.Error(err)
	}

	if bimg.DetermineImageTypeName(image) != "jpeg" {
		t.Fatalf("Invalid image type")
	}
}

func TestTTNEWOriginDirectory(t *testing.T) {
	opts := ServerOptions{
		Origin: "./tests/ttnew",
	}
	fn := ImageMiddleware(opts)(Origin)
	LoadSources(opts)

	ts := httptest.NewServer(fn)
	url := ts.URL + "/ttnew/r/2018/11/11/large.jpg"
	defer ts.Close()

	res, err := http.Get(url)
	if err != nil {
		t.Fatal("Cannot perform the request")
	}
	if res.StatusCode != 200 {
		t.Fatalf("Invalid response status: %d", res.StatusCode)
	}

	image, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	if len(image) == 0 {
		t.Fatalf("Empty response body")
	}

	err = assertSize(image, 1000, 750)
	if err != nil {
		t.Error(err)
	}

	if bimg.DetermineImageTypeName(image) != "jpeg" {
		t.Fatalf("Invalid image type")
	}
}

func TestCacheDirectory(t *testing.T) {
	opts := ServerOptions{
		Origin: "./tests/cache",
	}
	fn := ImageMiddleware(opts)(Cache)
	LoadSources(opts)

	ts := httptest.NewServer(fn)
	url := ts.URL + "/zoom/700_800/ttnew/large.jpg"
	defer ts.Close()

	res, err := http.Get(url)
	if err != nil {
		t.Fatal("Cannot perform the request")
	}
	if res.StatusCode != 200 {
		t.Fatalf("Invalid response status: %d", res.StatusCode)
	}

	image, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	if len(image) == 0 {
		t.Fatalf("Empty response body")
	}

	err = assertSize(image, 700, 800)
	if err != nil {
		t.Error(err)
	}

	if bimg.DetermineImageTypeName(image) != "jpeg" {
		t.Fatalf("Invalid image type")
	}
}

func TestOriginInvalidDirectory(t *testing.T) {
	fn := ImageMiddleware(ServerOptions{Origin: "_invalid_"})(Crop)
	ts := httptest.NewServer(fn)
	url := ts.URL + "?top=100&left=100&areawidth=200&areaheight=120&file=large.jpg"
	defer ts.Close()

	res, err := http.Get(url)
	if err != nil {
		t.Fatal("Cannot perform the request")
	}

	if res.StatusCode != 400 {
		t.Fatalf("Invalid response status: %d", res.StatusCode)
	}
}

func TestOriginInvalidPath(t *testing.T) {
	fn := ImageMiddleware(ServerOptions{Origin: "_invalid_"})(Crop)
	ts := httptest.NewServer(fn)
	url := ts.URL + "?top=100&left=100&areawidth=200&areaheight=120&file=../../large.jpg"
	defer ts.Close()

	res, err := http.Get(url)
	if err != nil {
		t.Fatal("Cannot perform the request")
	}

	if res.StatusCode != 400 {
		t.Fatalf("Invalid response status: %s", res.Status)
	}
}

func controller(op Operation) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		buf, _ := ioutil.ReadAll(r.Body)
		imageHandler(w, r, buf, op, ServerOptions{})
	}
}

func testServer(fn func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(fn))
}

func readFile(file string) io.Reader {
	buf, _ := os.Open(path.Join("tests", file))
	return buf
}

func assertSize(buf []byte, width, height int) error {
	size, err := bimg.NewImage(buf).Size()
	if err != nil {
		return err
	}
	if size.Width != width || size.Height != height {
		return fmt.Errorf("Invalid image size: %dx%d", size.Width, size.Height)
	}
	return nil
}
