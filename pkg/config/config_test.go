package config_test

import (
	"testing"

	"github.com/olzemal/go-emoji/pkg/config"
)

func TestLoadAlias(t *testing.T) {
	cfg, _ := config.LoadFile("../../opt/config.yaml")
	tests := []struct {
		I string
		O string
	}{
		{I: "", O: ""},
		{I: "smile", O: "slightly_smiling_face"},
		{I: ":)", O: "slightly_smiling_face"},
	}
	for _, test := range tests {
		got := cfg.EmojiFromAlias(test.I)
		if got != test.O {
			t.Errorf("Got: `%v`, Want: `%v`", got, test.O)
		}
	}
}
