package main

import (
	"fmt"

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

func main() {
	var config profiles

	if _, err := toml.DecodeFile("../tests/config.toml", &config); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Test \n")

	for _, p := range config.Profile {
		fmt.Printf("Profile: %s \n", p.Name)
	}
}
