package conf

import (
	"encoding/json"

	"v2ray.com/core"
	"v2ray.com/core/app/dispatcher"
	"v2ray.com/core/common/serial"
)

type SniffingConfig struct {
	Enabled      bool        `json:"enabled"`
	DestOverride *StringList `json:"destOverride"`
}

type MuxConfig struct {
	Enabled     bool  `json:"enabled"`
	Concurrency int16 `json:"concurrency"`
}

type InboundDetourAllocationConfig struct {
	Strategy string `json:"strategy"`
	Concurrency *uint32 `json:"concurrency"`
	RefreshMin *uint32 `json:"refresh"`
}

type InboundDetourConfig struct {
	Protocol string `json:"protocol"`
	PortRange *PortRange `json:"port"`
	ListenOn *Address `json:"listen"`
	Settings *json.RawMessage `json:"settings"`
	Tag string `json:"tag"`
	Allocation *InboundDetourAllocationConfig `json:"allocate"`
	StreamSetting *StreamConfig `json:"streamSettings"`
	DomainOverride *StringList `json:"domainOverride"`
	SniffingConfig *SniffingConfig `json:"sniffing"`
}

type OutboundDetourConfig struct {
	Protocol      string           `json:"protocol"`
	SendThrough   *Address         `json:"sendThrough"`
	Tag           string           `json:"tag"`
	Settings      *json.RawMessage `json:"settings"`
	StreamSetting *StreamConfig    `json:"streamSettings"`
	ProxySettings *ProxyConfig     `json:"proxySettings"`
	MuxSettings   *MuxConfig       `json:"mux"`
}

type StatsConfig struct{}

type Config struct {
	LogConfig *LogConfig `json:"log"`
	RouterConfig *RouterConfig `json:"routing"`
	DNSConfig *DnsConfig `json:"dns"`
	InboundConfigs []InboundDetourConfig `json:"inbounds"`
	OutboundConfigs []OutboundDetourConfig `json:"outbounds"`
	Transport *TransportConfig `json:"transport"`
	Policy *PolicyConfig `json:"policy"`
	Api *ApiConfig `json:"api"`
	Stats *StatsConfig `json:"stats"`
	Reverse *ReverseConfig `json:"reverse"`
}

// Build implements Buildable.
func (c *Config) Build() (*core.Config, error) {
	config := &core.Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
	}

	if c.Api != nil {
		apiConf, err := c.Api.Build()
		if err != nil {
			return nil, err
		}
		config.App = append(config.App, serial.ToTypedMessage(apiConf))
	}

	if c.Stats != nil {
		statsConf, err := c.Stats.Build()
		if err != nil {
			return nil, err
		}
		config.App = append(config.App, serial.ToTypedMessage(statsConf))
	}

	var logConfMsg *serial.TypedMessage
	if c.LogConfig != nil {
		logConfMsg = serial.ToTypedMessage(c.LogConfig.Build())
	} else {
		logConfMsg = serial.ToTypedMessage(DefaultLogConfig())
	}
	// let logger module be the first App to start,
	// so that other modules could print log during initiating
	config.App = append([]*serial.TypedMessage{logConfMsg}, config.App...)

	if c.RouterConfig != nil {
		routerConfig, err := c.RouterConfig.Build()
		if err != nil {
			return nil, err
		}
		config.App = append(config.App, serial.ToTypedMessage(routerConfig))
	}

	if c.DNSConfig != nil {
		dnsApp, err := c.DNSConfig.Build()
		if err != nil {
			return nil, newError("failed to parse DNS config").Base(err)
		}
		config.App = append(config.App, serial.ToTypedMessage(dnsApp))
	}

	if c.Policy != nil {
		pc, err := c.Policy.Build()
		if err != nil {
			return nil, err
		}
		config.App = append(config.App, serial.ToTypedMessage(pc))
	}

	if c.Reverse != nil {
		r, err := c.Reverse.Build()
		if err != nil {
			return nil, err
		}
		config.App = append(config.App, serial.ToTypedMessage(r))
	}

	var inbounds []InboundDetourConfig

	if c.InboundConfig != nil {
		inbounds = append(inbounds, *c.InboundConfig)
	}

	if len(c.InboundDetours) > 0 {
		inbounds = append(inbounds, c.InboundDetours...)
	}

	if len(c.InboundConfigs) > 0 {
		inbounds = append(inbounds, c.InboundConfigs...)
	}

	// Backward compatibility.
	if len(inbounds) > 0 && inbounds[0].PortRange == nil && c.Port > 0 {
		inbounds[0].PortRange = &PortRange{
			From: uint32(c.Port),
			To:   uint32(c.Port),
		}
	}

	for _, rawInboundConfig := range inbounds {
		if c.Transport != nil {
			if rawInboundConfig.StreamSetting == nil {
				rawInboundConfig.StreamSetting = &StreamConfig{}
			}
			applyTransportConfig(rawInboundConfig.StreamSetting, c.Transport)
		}
		ic, err := rawInboundConfig.Build()
		if err != nil {
			return nil, err
		}
		config.Inbound = append(config.Inbound, ic)
	}

	var outbounds []OutboundDetourConfig

	if c.OutboundConfig != nil {
		outbounds = append(outbounds, *c.OutboundConfig)
	}

	if len(c.OutboundDetours) > 0 {
		outbounds = append(outbounds, c.OutboundDetours...)
	}

	if len(c.OutboundConfigs) > 0 {
		outbounds = append(outbounds, c.OutboundConfigs...)
	}

	for _, rawOutboundConfig := range outbounds {
		if c.Transport != nil {
			if rawOutboundConfig.StreamSetting == nil {
				rawOutboundConfig.StreamSetting = &StreamConfig{}
			}
			applyTransportConfig(rawOutboundConfig.StreamSetting, c.Transport)
		}
		oc, err := rawOutboundConfig.Build()
		if err != nil {
			return nil, err
		}
		config.Outbound = append(config.Outbound, oc)
	}

	return config, nil
}