package core

import (
	serial "v2ray.com/core/common/serial"
)

type Config struct {
	// Inbound handler configurations. Must have at least one item.
	Inbound []*InboundHandlerConfig `protobuf:"bytes,1,rep,name=inbound,proto3" json:"inbound,omitempty"`
	// Outbound handler configurations. Must have at least one item. The first item is used as default for routing.
	Outbound []*OutboundHandlerConfig `protobuf:"bytes,2,rep,name=outbound,proto3" json:"outbound,omitempty"`
	// App is for configurations of all features in V2Ray. A feature must implement the Feature interface, and its config type must be registered through common.RegisterConfig.
	App []*serial.TypedMessage `protobuf:"bytes,4,rep,name=app,proto3" json:"app,omitempty"`
	// Configuration for extensions. The config may not work if corresponding extension is not loaded into V2Ray.
	// V2Ray will ignore such config during initialization.
	Extension            []*serial.TypedMessage `protobuf:"bytes,6,rep,name=extension,proto3" json:"extension,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

// InboundHandlerConfig is the configuration for inbound handler.
type InboundHandlerConfig struct {
	// Tag of the inbound handler. The tag must be unique among all inbound handlers.
	Tag string `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	// Settings for how this inbound proxy is handled.
	ReceiverSettings *serial.TypedMessage `protobuf:"bytes,2,opt,name=receiver_settings,json=receiverSettings,proto3" json:"receiver_settings,omitempty"`
	// Settings for inbound proxy. Must be one of inbound proxies.
	ProxySettings        *serial.TypedMessage `protobuf:"bytes,3,opt,name=proxy_settings,json=ProxySettings,proto3" json:proxy_settings,omitempty`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

// OutboundHandlerConfig is the configuration for outbound handler.
type OutboundHandlerConfig struct {
	// Tag of this outbound handler.
	Tag string `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	// Settings for how to dial connection for this outbound handler.
	SenderSettings *serial.TypedMessage `protobuf:"bytes,3,opt,name=sender_settings,json=senderSettings,proto3" json:"sender_settings,omitempty"`
	// Settings for this outbound proxy. Must be one of outbound proxies.
	ProxySettings        *serial.TypedMessage `protobuf:"bytes,3,opt,name=proxy_settings,json=ProxySettings,proto3" json:proxy_settings,omitempty`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}
