package main

import (
	"io/ioutil"
	"net/http"
	"path"
	"strings"
)

const ImageSourceTypeFileSystem ImageSourceType = "fs"

type FileSystemImageSource struct {
	Config *SourceConfig
}

func NewFileSystemImageSource(config *SourceConfig) ImageSource {
	return &FileSystemImageSource{config}
}

func (s *FileSystemImageSource) Matches(r *http.Request) bool {
	action, size, project, file := s.getFileParam(r)
	return r.Method == "GET" && action != "" && s.checkSize(action, size) && project != "" && file != ""
}

func (s *FileSystemImageSource) GetImage(r *http.Request) ([]byte, error) {
	action, _, project, file := s.getFileParam(r)

	if action == "ttnew" || action == "datg" || action == "ttc" || action == "sticker" {
		file, err := s.buildPath("r", file)

		if err != nil {
			return nil, err
		}

		return s.read(file)
	}

	if file == "" {
		return nil, ErrMissingParamFile
	}

	file, err := s.buildPath(project, file)

	if err != nil {
		return nil, err
	}

	return s.read(file)
}

func (s *FileSystemImageSource) checkSize(action string, size string) bool {
	return true
}

func (s *FileSystemImageSource) buildPath(project string, file string) (string, error) {
	file = path.Clean(path.Join(s.Config.OriginPath, project, file))
	if strings.HasPrefix(file, path.Join(s.Config.OriginPath)) == false {
		file = path.Clean(s.Config.Default)
	}
	return file, nil
}

func (s *FileSystemImageSource) read(file string) ([]byte, error) {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		buf, _ = ioutil.ReadFile(s.Config.Default)
	}
	return buf, nil
}

func (s *FileSystemImageSource) getFileParam(r *http.Request) (string, string, string, string) {
	params := strings.Split(r.URL.Path, "/")

	if len(params) <= 3 {
		return "", "", "", ""
	}

	action := params[1]

	if action == "ttnew" || action == "datg" || action == "ttc" || action == "sticker" {
		if params[2] == "r" {
			size := ""
			project := params[2]
			file := strings.Join(params[3:], "/")

			return action, size, project, file
		}

		if params[2] == "i" {
			size := params[3]
			project := params[2]
			file := strings.Join(params[4:], "/")

			return action, size, project, file
		}
	}

	if action == "origin" {
		size := ""
		project := params[2]
		file := strings.Join(params[3:], "/")

		return action, size, project, file
	}

	size := params[2]
	project := params[3]
	file := strings.Join(params[4:], "/")

	return action, size, project, file
}

func init() {
	RegisterSource(ImageSourceTypeFileSystem, NewFileSystemImageSource)
}
