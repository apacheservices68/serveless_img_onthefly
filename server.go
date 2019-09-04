package main

import (
	"net/http"
	"os"
	"time"
)

// ServerOptions config server options
type ServerOptions struct {
	Burst        int
	Concurrency  int
	HTTPCacheTTL int
	Gzip         bool
	Origin       string
	Cache        string
	Default      string
	Profiles     profiles
}

// Server init server
func Server(o ServerOptions) error {
	handler := NewLog(NewServerMux(o), os.Stdout)

	server := &http.Server{
		Addr:           ":3300",
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    time.Duration(60) * time.Second,
		WriteTimeout:   time.Duration(60) * time.Second,
	}

	return server.ListenAndServe()
}

// NewServerMux creates a new HTTP server route multiplexer.
func NewServerMux(o ServerOptions) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", Middleware(indexController, o))
	mux.Handle("/health", Middleware(healthController, o))

	image := ImageMiddleware(o)

	mux.Handle("/origin/", image(Origin))

	// ttnew
	mux.Handle("/ttnew/r/", image(Origin))
	mux.Handle("/ttnew/i/", image(Resize))

	// datg
	mux.Handle("/datg/r/", image(Origin))
	mux.Handle("/datg/i/", image(Resize))

	// ttc
	mux.Handle("/ttc/r/", image(Origin))
	mux.Handle("/ttc/i/", image(Resize))

	// ttc-sticker
	mux.Handle("/sticker/r/", image(Origin))
	mux.Handle("/sticker/i/", image(Resize))

	mux.Handle("/fit/", image(Fit))
	mux.Handle("/crop/", image(Crop))
	mux.Handle("/zoom/", image(Zoom))
	mux.Handle("/thumb_w/", image(thumbWidth))
	mux.Handle("/thumb_h/", image(thumbHeight))

	return mux
}
