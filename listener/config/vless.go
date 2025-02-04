package config

import (
	"github.com/metacubex/mihomo/listener/sing"

	"encoding/json"
)

type VlessUser struct {
	Username string
	UUID     string
	Flow     string
}

type VlessServer struct {
	Enable        bool
	Listen        string
	Users         []VlessUser
	WsPath        string
	Certificate   string
	PrivateKey    string
	RealityConfig RealityConfig  `yaml:"reality-config" json:"reality-config"`
	MuxOption     sing.MuxOption `yaml:"mux-option" json:"mux-option,omitempty"`
}

func (t VlessServer) String() string {
	b, _ := json.Marshal(t)
	return string(b)
}

type RealityConfig struct {
	Dest              string
	PrivateKey        string `yaml:"private-key" json:"private-key"`
	ShortID           []string `yaml:"short-id" json:"short-id"`
	ServerNames       []string
	MaxTimeDifference int
	Proxy             string
}
