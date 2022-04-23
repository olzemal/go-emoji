package config

import (
	"fmt"
	"io"
	"os"

	"github.com/olzemal/go-emoji/pkg/util"
	"gopkg.in/yaml.v2"
)

/* Example config.yaml
---
aliases:
- emoji: slightly_smiling_face
  names:
  - smile
  - ":)"
*/

// Config holds the collection of all aliases
type Config struct {
	Aliases []Alias `yaml:"aliases"`
}

// Alias binds one or mor names to a emoji name
type Alias struct {
	Emoji string   `yaml:"emoji"`
	Names []string `yaml:"names"`
}

func (c Config) EmojiFromAlias(alias string) string {
	for _, a := range c.Aliases {
		for _, n := range a.Names {
			if n == alias {
				return a.Emoji
			}
		}
	}
	return ""
}

func LoadFile(path string) (Config, error) {
	if !util.CanReadFrom(path) {
		return Config{}, fmt.Errorf("Cannot read from path `%v`.", path)
	}
	f, _ := os.Open(path)
	defer f.Close()
	cfg, _ := io.ReadAll(f)
	var config Config
	err := yaml.Unmarshal(cfg, &config)
	return config, err
}
