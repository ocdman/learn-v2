package conf

import (
	"v2ray.com/core/common/net"
)

type StringList []string

type Address struct {
	net.Address
}

type PortRange struct {
	From uint32
	To uint32
}