package main

import (
	"net/http"
)

type ImageSourceType string
type ImageSourceFactoryFunction func(*SourceConfig) ImageSource

type SourceConfig struct {
	OriginPath string
	CachePath  string
	Default    string
	Type       ImageSourceType
}

var imageSourceMap = make(map[ImageSourceType]ImageSource)
var imageSourceFactoryMap = make(map[ImageSourceType]ImageSourceFactoryFunction)

type ImageSource interface {
	Matches(*http.Request) bool
	GetImage(*http.Request) ([]byte, error)
}

func RegisterSource(sourceType ImageSourceType, factory ImageSourceFactoryFunction) {
	imageSourceFactoryMap[sourceType] = factory
}

func LoadSources(o ServerOptions) {
	for name, factory := range imageSourceFactoryMap {
		imageSourceMap[name] = factory(&SourceConfig{
			Type:       name,
			OriginPath: o.Origin,
			CachePath:  o.Cache,
			Default:    o.Default,
		})
	}
}

func MatchSource(req *http.Request) ImageSource {
	for _, source := range imageSourceMap {
		if source.Matches(req) {
			return source
		}
	}
	return nil
}
