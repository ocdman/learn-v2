package conf

type BridgeConfig struct {
	Tag    string `json:"tag"`
	Domain string `json:"domain"`
}

type PortalConfig struct {
	Tag    string `json:"tag"`
	Domain string `json:"domain"`
}

type ReverseConfig struct {
	Bridges []BridgeConfig `json:"bridges"`
	Portals []PortalConfig `json:"portals"`
}