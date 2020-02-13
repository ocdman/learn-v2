package platform_test

import (
	"os"
	"runtime"
	"testing"

	. "v2ray.com/core/common/platform"
)

func TestNormalizeEnvName(t *testing.T) {
	cases := []struct {
		input  string
		output string
	}{
		{
			input:  "a",
			output: "A",
		},
		{
			input:  "a.a",
			output: "A_A",
		},
		{
			input:  "A.A.B",
			output: "A_A_B",
		},
	}
	for _, test := range cases {
		if v := NormalizeEnvName(test.input); v != test.output {
			t.Error("unexpected output: ", v, " want ", test.output)
		}
	}
}

func TestEnvFlag(t *testing.T) {
	if v := (EnvFlag{
		Name: "xxxxx.y",
	}.GetValueAsInt(10)); v != 10 {
		t.Error("env value: ", v)
	}
}

func TestGetConfigurationPath(t *testing.T) {
	os.Setenv("v2ray.location.config", "/v2ray")
	if runtime.GOOS == "windows" {
		if v := GetConfigurationPath(); v != "\\v2ray\\config.json" {
			t.Error("config path: ", v)
		}
	} else {
		if v := GetConfigurationPath(); v != "/v2ray/config.json" {
			t.Error("config path: ", v)
		}
	}
}
