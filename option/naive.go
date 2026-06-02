package option

import (
	"github.com/sagernet/sing/common/auth"
	"github.com/sagernet/sing/common/byteformats"
	"github.com/sagernet/sing/common/json/badoption"
)

type QuicheCongestionControl string

const (
	QuicheCongestionControlDefault QuicheCongestionControl = ""
	QuicheCongestionControlBBR     QuicheCongestionControl = "TBBR"
	QuicheCongestionControlBBRv2   QuicheCongestionControl = "B2ON"
	QuicheCongestionControlCubic   QuicheCongestionControl = "QBIC"
	QuicheCongestionControlReno    QuicheCongestionControl = "RENO"
)

// FallbackSite FORKED
type FallbackSite struct {
	Address    string `json:"address,omitempty"`
	Port       int    `json:"port,omitempty"`
	ForceHTTPS bool   `json:"force_https,omitempty"` // If true, proxied to [https://], else proxied to [http://]
}

// NaiveInboundOptions FORKED
type NaiveInboundOptions struct {
	ListenOptions
	Users                 []auth.User  `json:"users,omitempty"`
	Network               NetworkList  `json:"network,omitempty"`
	QUICCongestionControl string       `json:"quic_congestion_control,omitempty"`
	FallbackSite          FallbackSite `json:"fallback_site,omitempty"`
	InboundTLSOptionsContainer
}

type NaiveOutboundOptions struct {
	DialerOptions
	ServerOptions
	Username                 string                   `json:"username,omitempty"`
	Password                 string                   `json:"password,omitempty"`
	InsecureConcurrency      int                      `json:"insecure_concurrency,omitempty"`
	ExtraHeaders             badoption.HTTPHeader     `json:"extra_headers,omitempty"`
	ReceiveWindow            *byteformats.MemoryBytes `json:"stream_receive_window,omitempty"`
	UDPOverTCP               *UDPOverTCPOptions       `json:"udp_over_tcp,omitempty"`
	QUIC                     bool                     `json:"quic,omitempty"`
	QUICCongestionControl    string                   `json:"quic_congestion_control,omitempty"`
	QUICSessionReceiveWindow *byteformats.MemoryBytes `json:"quic_session_receive_window,omitempty"`
	OutboundTLSOptionsContainer
}
