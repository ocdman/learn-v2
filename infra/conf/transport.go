package conf

type TransportConfig struct {
	TCPConfig  *TCPConfig          `json:"tcpSettings"`
	KCPConfig  *KCPConfig          `json:"kcpSettings"`
	WSConfig   *WebSocketConfig    `json:"wsSettings"`
	HTTPConfig *HTTPConfig         `json:"httpSettings"`
	DSConfig   *DomainSocketConfig `json:"dsSettings"`
	QUICConfig *QUICConfig         `json:"quicSettings"`
}