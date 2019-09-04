package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	d "runtime/debug"
	"time"

	"github.com/BurntSushi/toml"
)

type profile struct {
	Name   string
	Width  int
	Height int
	NoCrop bool
}

type profiles struct {
	Profile []profile
}

var (
	aVers         = flag.Bool("v", false, "Show version")
	aVersl        = flag.Bool("version", false, "Show version")
	aHelp         = flag.Bool("h", false, "Show help")
	aHelpl        = flag.Bool("help", false, "Show help")
	aGzip         = flag.Bool("gzip", false, "Enable gzip compression (deprecated)")
	aOrigin       = flag.String("origin", "", "Mount server origin local directory")
	aCache        = flag.String("cache", "", "Mount server cache local directory")
	aDefault      = flag.String("default", "", "Path Image default")
	aProfile      = flag.String("profile", "", "Path profile config (config.toml)")
	aHTTPCacheTTL = flag.Int("http-cache-ttl", -1, "The TTL in seconds")
	aConcurrency  = flag.Int("concurrency", 0, "Throttle concurrency limit per second")
	aBurst        = flag.Int("burst", 100, "Throttle burst max cache size")
	aMRelease     = flag.Int("mrelease", 30, "OS memory release interval in seconds")
	aCpus         = flag.Int("cpus", runtime.GOMAXPROCS(-1), "Number of cpu cores to use")
)

const usage = `tto-resize %s

Usage:
  tto-resize -concurrency 10
  tto-resize -origin ./tests/origin -cache ./tests/cache
  tto-resize -h | -help
  tto-resize -v | -version

Options:
  -h, -help                 Show help
  -v, -version              Show version
  -path-prefix <value>      Url path prefix to listen to [default: "/"]
  -gzip                     Enable gzip compression (deprecated) [default: false]
  -origin <path>            Mount server origin local directory
  -cache <path>             Mount server cache local directory
  -profile <path>           Path profile config (config.toml)
  -default <path>           Path Image default
  -http-cache-ttl <num>     The TTL in seconds. Adds caching headers to locally served files.
  -concurrency <num>        Throttle concurrency limit per second [default: disabled]
  -burst <num>              Throttle burst max cache size [default: 100]
  -mrelease <num>           OS memory release interval in seconds [default: 30]
  -cpus <num>               Number of used cpu cores.
                            (default for current machine is %d cores)
`

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage, Version, runtime.NumCPU()))
	}
	flag.Parse()

	if *aHelp || *aHelpl {
		showUsage()
	}
	if *aVers || *aVersl {
		showVersion()
	}

	// Only required in Go < 1.5
	runtime.GOMAXPROCS(*aCpus)

	var configs profiles

	if *aProfile != "" {
		checkProfileConfig(*aProfile)

		if _, err := toml.DecodeFile(*aProfile, &configs); err != nil {
			exitWithError("error config profile: %s", err)
		}
	}

	opts := ServerOptions{
		Concurrency:  *aConcurrency,
		Burst:        *aBurst,
		Origin:       *aOrigin,
		Cache:        *aCache,
		Default:      *aDefault,
		Profiles:     configs,
		HTTPCacheTTL: *aHTTPCacheTTL,
	}

	// Show warning if gzip flag is passed
	if *aGzip {
		fmt.Println("warning: -gzip flag is deprecated and will not have effect")
	}

	// Create a memory release goroutine
	if *aMRelease > 0 {
		memoryRelease(*aMRelease)
	}

	// Check if the mount directory exists, if present
	if *aOrigin != "" {
		checkMountDirectory(*aOrigin)
	}

	if *aCache != "" {
		checkMountDirectory(*aCache)
	}

	if *aDefault != "" {
		checkFileDefault(*aDefault)
	}

	// Validate HTTP cache param, if present
	if *aHTTPCacheTTL != -1 {
		checkHTTPCacheTTL(*aHTTPCacheTTL)
	}

	debug("tto-resize server listening on port 3300")

	// Load image source providers
	LoadSources(opts)

	// Start the server
	err := Server(opts)
	if err != nil {
		exitWithError("cannot start the server: %s", err)
	}
}

func showUsage() {
	flag.Usage()
	os.Exit(1)
}

func showVersion() {
	fmt.Println(Version)
	os.Exit(1)
}

func checkMountDirectory(path string) {
	src, err := os.Stat(path)
	if err != nil {
		exitWithError("error while mounting directory: %s", err)
	}
	if src.IsDir() == false {
		exitWithError("mount path is not a directory: %s", path)
	}
	if path == "/" {
		exitWithError("cannot mount root directory for security reasons")
	}
}

func checkFileDefault(path string) {
	_, err := os.Stat(path)
	if err != nil {
		exitWithError("error image default: %s", err)
	}
	if path == "/" {
		exitWithError("default image not exists")
	}
}

func checkProfileConfig(path string) {
	_, err := os.Stat(path)
	if err != nil {
		exitWithError("error config profile: %s", err)
	}
	if path == "/" {
		exitWithError("config profile not exists")
	}
}

func checkHTTPCacheTTL(ttl int) {
	if ttl < -1 || ttl > 31556926 {
		exitWithError("The -http-cache-ttl flag only accepts a value from 0 to 31556926")
	}

	if ttl == 0 {
		debug("Adding HTTP cache control headers set to prevent caching.")
	}
}

func memoryRelease(interval int) {
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	go func() {
		for range ticker.C {
			debug("FreeOSMemory()")
			d.FreeOSMemory()
		}
	}()
}

func exitWithError(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args)
	os.Exit(1)
}

func debug(msg string, values ...interface{}) {
	debug := os.Getenv("DEBUG")
	if debug == "tto-resize" || debug == "*" {
		log.Printf(msg, values...)
	}
}
