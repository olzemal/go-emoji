package main

import (
	"fmt"
	"os"

	"github.com/olzemal/go-emoji/pkg/cache"
	"github.com/olzemal/go-emoji/pkg/config"
	"github.com/olzemal/go-emoji/pkg/query"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Errorf("os.UserHomeDir: %v", err))
	}
	cacheDirPath := home + "/.cache/go-emoji"
	cacheFilePath := cache.Init(cacheDirPath)
	configFilePath := home + "/.config/go-emoji/config.yaml"

	c, err := cache.Import(cacheFilePath)
	if err != nil {
		panic(fmt.Errorf("cache.Import: %v", err))
	}

	cfg, err := config.LoadFile(configFilePath)
	if err != nil {
		panic(fmt.Errorf("config.LoadFile: %v", err))
	}

	if len(os.Args) > 1 {
		s := cfg.EmojiFromAlias(os.Args[1])
		if len(s) == 0 {
			s = os.Args[1]
		}
		c = query.FilterContains(c, s)
	}
	for key, value := range c {
		fmt.Printf("%s=%s\n", key, value)
	}
}
