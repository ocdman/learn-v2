package conf

import (
	"encoding/json"
)

type KCPConfig struct {
	Mtu *uint32 `json:"mtu"`
	Tti *uint32 `json:"tti"`
	UpCap *uint32 `json:"uplinkCapacity"`
	DownCap *uint32 `json:"downlinkCapacity"`
	Congestion *bool `json:"congestion"`
	ReadBufferSize *uint32 `json:"readBufferSize"`
	WriteBufferSize *uint32 `json:"writeBufferSize"`
	HeaderConfig json.RawMessage `json:"header"`
}

type TCPConfig struct {
	HeaderConfig json.RawMessage `json:"header"`
}

type WebSocketConfig struct {
	Path string `json:"path"`
	Path2 string `json:"Path"` // The key was misspelled. For backward compatibility, we have to keep track the old key.
	Headers map[string]string `json:"headers"`
}

type HTTPConfig struct {
	Host *StringList `json:"host"`
	Path string      `json:"path"`
}

type QUICConfig struct {
	Header   json.RawMessage `json:"header"`
	Security string          `json:"security"`
	Key      string          `json:"key"`
}

type DomainSocketConfig struct {
	Path     string `json:"path"`
	Abstract bool   `json:"abstract"`
}

type TLSCertConfig struct {
	CertFile string `json:"certificateFile"`
	CertStr []string `json:"certificate"`
	KeyFile string `json:"keyFile"`
	KeyStr []string `json:"key"`
	Usage string `json:"usage"`
}

type TLSConfig struct {
	Insecure bool `json:"allowInsecure"`
	InsecureCiphers bool `json:"allowInsecureCiphers"`
	Certs []*TLSCertConfig `json:"certificates"`
	ServerName string `json:"serverName"`
	ALPN *StringList `json:"alpn"`
	DiableSystemRoot bool `json:"disableSystemRoot"`
}

type TransportProtocol string

type SocketConfig struct {
	Mark   int32  `json:"mark"`
	TFO    *bool  `json:"tcpFastOpen"`
	TProxy string `json:"tproxy"`
}

type StreamConfig struct {
	Network *TransportProtocol `json:"network"`
	Security string `json:"security"`
	TLSSettings *TLSConfig `json:"tlsSettings"`
	TCPSettings *TCPConfig `json:"tcpSettings"`
	KCPSettings *KCPConfig `json:"kcpSettings"`
	WSSettings *WebSocketConfig `json:"wsSettings"`
	HTTPSettings *HTTPConfig `json:"httpSettings"`
	DSSettings *DomainSocketConfig `json:"dsSettings"`
	QUICSettings *QUICConfig `json:"quicSettings"`
	SocketSettings *SocketConfig `json:"sockopt"`
}

type ProxyConfig struct {
	Tag string `json:"tag"`
}