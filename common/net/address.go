package net

import (
	"net"
)

// AddressFamily is the type of address.
type AddressFamily byte

// Address represents a network address to be communicated with. It may be an  IP address or domain
// address, not both. This interface doesn't resolve IP address for a given domain.
type Address interface {
	IP() net.IP // IP of this Address
	Domain() string // Domain of this Address
	Family() AddressFamily

	String() string // String representation of this Address
}