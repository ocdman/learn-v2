// +build !confonly

package core

import (
	"fmt"
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

// RegisterConfigLoader add a new ConfigLoader.
func RegisterConfigLoader(format *ConfigFormat) error {
	name := strings.ToLower(format.Name)
	if _, found := configLoaderByName[name]; found {
		return newError(format.Name, " already registered.")
	}
	configLoaderByName[name] = format

	for _, ext := range format.Extension {
		lext := strings.ToLower(ext)
		if f, found := configLoaderByExt[lext]; found {
			return newError(ext, " already registered to ", f.Name)
		}
		configLoaderByExt[ext] = format
	}

	fmt.Println("configLoaderByName = ", configLoaderByName)
	fmt.Println("configLoaderByExt = ", configLoaderByExt)

	return nil
}

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
	fmt.Println("ext = ", ext)
	if len(ext) > 0 {
		if f, found := configLoaderByExt[ext]; found {
			fmt.Println("configLoaderByExt found")
			return f.Loader(input)
		}
		fmt.Println("configLoaderByExt not found")
	}

	if f, found := configLoaderByName[formatName]; found {
		fmt.Println("configLoaderByName found")
		return f.Loader(input)
	}
	fmt.Println("configLoaderByName not found")

	return nil, newError("Unable to load config in ", formatName).AtWarning()
}

func init() {
	fmt.Println("core config.go init")
	RegisterConfigLoader(&ConfigFormat{
		Name:      "Protobuf",
		Extension: []string{"pb"},
		// Loader: loadProtobufConfig,
	})
}
