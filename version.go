package main

import (
	"math"
	"runtime"
	"time"

	"github.com/h2non/bimg"
)

var start = time.Now()

// Version stores the current package semantic version
const Version = "0.5.1"

const MB float64 = 1.0 * 1024 * 1024

// Version represents the supported version
type Versions struct {
	ResizeVersion string `json:"tto"`
}

// CurrentVersions stores the current runtime system version metadata
var CurrentVersions = Versions{Version}

// Versions represents the supported version
type HealthStats struct {
	TTOResizeVersion     string  `json:"tto-resize"`
	BimgVersion          string  `json:"bimg"`
	VipsVersion          string  `json:"libvips"`
	Uptime               int64   `json:"uptime"`
	AllocatedMemory      float64 `json:"allocatedMemory"`
	TotalAllocatedMemory float64 `json:"totalAllocatedMemory"`
	Goroutines           int     `json:"goroutines"`
	NumberOfCPUs         int     `json:"cpus"`
}

// CurrentVersions stores the current runtime system version metadata

func GetHealthStats() *HealthStats {
	mem := &runtime.MemStats{}
	runtime.ReadMemStats(mem)

	return &HealthStats{
		TTOResizeVersion:     Version,
		BimgVersion:          bimg.Version,
		VipsVersion:          bimg.VipsVersion,
		Uptime:               GetUptime(),
		AllocatedMemory:      toMegaBytes(mem.Alloc),
		TotalAllocatedMemory: toMegaBytes(mem.TotalAlloc),
		Goroutines:           runtime.NumGoroutine(),
		NumberOfCPUs:         runtime.NumCPU(),
	}
}

func GetUptime() int64 {
	return time.Now().Unix() - start.Unix()
}

func toMegaBytes(bytes uint64) float64 {
	return toFixed(float64(bytes)/MB, 2)
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
