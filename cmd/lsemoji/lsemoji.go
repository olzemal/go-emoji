package main

import (
	"fmt"
	"os"

	cache "github.com/olzemal/go-emoji/pkg/cache"
	"github.com/olzemal/go-emoji/pkg/query"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Errorf("os.UserHomeDir: %v", err))
	}
	cacheDirPath := home + "/.cache/go-emoji"
	cacheFilePath := cache.Init(cacheDirPath)

	c, err := cache.Import(cacheFilePath)
	if err != nil {
		panic(fmt.Errorf("Emoji.ImportCache: %v", err))
	}

	if len(os.Args) > 1 {
		c = query.FilterContains(c, os.Args[1])
	}
	for key, value := range c {
		fmt.Printf("%s=%s\n", key, value)
	}
}
