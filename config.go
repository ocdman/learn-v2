// +build !confonly

package core

import (
	"io"
	"strings"
)

// ConfigFormat is a configurable format of V2Ray config file.
type ConfigFormat struct {
	Name      string
	Extension []string
	Loader    ConfigLoader
}

// ConfigLoader is a utility to load V2Ray config from external source.
type ConfigLoader func(input io.Reader) (*Config, error)

var (
	configLoaderByName = make(map[string]*ConfigFormat)
	configLoaderByExt  = make(map[string]*ConfigFormat)
)

func getExtension(filename string) string {
	idx := strings.LastIndexByte(filename, '.')
	if idx == -1 {
		return ""
	}
	return filename[idx+1:]
}

// LoadConfig loads config with given format from given source.
func LoadConfig(formatName string, filename string, input io.Reader) (*Config, error) {
	ext := getExtension(filename)
	if len(ext) > 0 {
		if f, found := configLoaderByExt[ext]; found {
			return f.Loader(input)
		}
	}

	if f, found := configLoaderByName[formatName]; found {
		return f.Loader(input)
	}

	return nil, newError("Unable to load config in ", formatName).AtWarning()
}
