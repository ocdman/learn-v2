package conf

type ApiConfig struct {
	Tag      string   `json:"tag"`
	Services []string `json:"services"`
}