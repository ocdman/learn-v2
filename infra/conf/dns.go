package conf

type NameServerConfig struct {
	Address *Address
	Port uint16
	Domains []string
	ExceptIPs StringList
}

type DnsConfig struct {
	Servers []*NameServerConfig `json:"servers"`
	Hosts map[string]*Address `json:"hosts"`
	ClientIP *Address `json:"clientIp"`
	Tag string `json:"tag"`
}