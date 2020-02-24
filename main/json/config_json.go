package json

//go:generate errorgen

import (
	"v2ray.com/core"
)

func init() {
	core.RegisterConfigLoader(&core.ConfigFormat{
		Name:      "JSON",
		Extension: []string{"json"},
	})
}
