package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestFileSystemImageSource(t *testing.T) {
	var body []byte
	var err error
	const fixtureFile = "tests/origin/ttnew/large.jpg"

	source := NewFileSystemImageSource(&SourceConfig{
		OriginPath: "tests/origin",
		CachePath:  "tests/cache",
	})

	fakeHandler := func(w http.ResponseWriter, r *http.Request) {
		if !source.Matches(r) {
			t.Fatal("Cannot match the request")
		}

		body, err = source.GetImage(r)
		if err != nil {
			t.Fatalf("Error while reading the body: %s", err)
		}
		w.Write(body)
	}

	file, _ := os.Open(fixtureFile)
	r, _ := http.NewRequest("GET", "http://foo/zoom/120_120/ttnew/large.jpg", file)
	w := httptest.NewRecorder()
	fakeHandler(w, r)

	buf, _ := ioutil.ReadFile(fixtureFile)
	if len(body) != len(buf) {
		t.Error("Invalid response body")
	}
}

func TestFileSystemImageSourceOriginTTNEW(t *testing.T) {
	var body []byte
	var err error
	const fixtureFile = "tests/ttnew/r/2018/11/11/large.jpg"

	source := NewFileSystemImageSource(&SourceConfig{
		OriginPath: "tests/ttnew",
		CachePath:  "tests/ttnew",
	})

	fakeHandler := func(w http.ResponseWriter, r *http.Request) {
		if !source.Matches(r) {
			t.Fatal("Cannot match the request")
		}

		body, err = source.GetImage(r)
		if err != nil {
			t.Fatalf("Error while reading the body: %s", err)
		}
		w.Write(body)
	}

	file, _ := os.Open(fixtureFile)
	r, _ := http.NewRequest("GET", "http://foo/ttnew/r/2018/11/11/large.jpg", file)
	w := httptest.NewRecorder()
	fakeHandler(w, r)

	buf, _ := ioutil.ReadFile(fixtureFile)
	if len(body) != len(buf) {
		t.Error("Invalid response body")
	}
}

func TestFileSystemImageSourceCacheTTNEW(t *testing.T) {
	var body []byte
	var err error
	const fixtureFile = "tests/ttnew/r/2018/11/11/large.jpg"

	source := NewFileSystemImageSource(&SourceConfig{
		OriginPath: "tests/ttnew",
		CachePath:  "tests/ttnew",
	})

	fakeHandler := func(w http.ResponseWriter, r *http.Request) {
		if !source.Matches(r) {
			t.Fatal("Cannot match the request")
		}

		body, err = source.GetImage(r)
		if err != nil {
			t.Fatalf("Error while reading the body: %s", err)
		}
		w.Write(body)
	}

	file, _ := os.Open(fixtureFile)
	r, _ := http.NewRequest("GET", "http://foo/ttnew/i/s640/2018/11/11/large.jpg", file)
	w := httptest.NewRecorder()
	fakeHandler(w, r)

	buf, _ := ioutil.ReadFile(fixtureFile)
	if len(body) != len(buf) {
		t.Error("Invalid response body")
	}
}
