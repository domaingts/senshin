package config

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestVlessParse(t *testing.T) {
	server := &VlessServer{
		Enable:        false,
		Listen:        "",
		Users:         []VlessUser{
			{
				Username: "",
				UUID:     "",
				Flow:     "",
			},
		},
		RealityConfig: RealityConfig{
			Dest:              "fijefjie",
			PrivateKey:        "erfergghh",
			ShortID:           []string{
				"fiefjie",
			},
			ServerNames:       []string{
				"fefijirr",
			},
		},
	}
	re, err := yaml.Marshal(server)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(re))
}