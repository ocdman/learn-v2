package jsonem

import (
	"v2ray.com/core"
	"v2ray.com/core/infra/conf/serial"
)

func init() {
	core.RegisterConfigLoader(&core.ConfigFormat{
		Name:      "JSON",
		Extension: []string{"json"},
		Loader:    serial.LoadJSONConfig,
	})
}
