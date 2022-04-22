package main

import (
	"fmt"
	"os"

	cache "github.com/olzemal/go-emoji/pkg/cache"
	"github.com/olzemal/go-emoji/pkg/query"
)

const (
	cacheFileName = "cache"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Errorf("os.UserHomeDir: %v", err))
	}
	cacheDirPath := home + "/.cache/go-emoji"
	initCache(cacheDirPath)

	c, err := cache.Import(cacheDirPath + "/" + cacheFileName)
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

// initCache ensures a filled Cache File is present in the cacheDirPath.
func initCache(cacheDirPath string) {
	if _, err := os.Stat(cacheDirPath); err != nil {
		err := os.Mkdir(cacheDirPath, 0755)
		if err != nil {
			panic(fmt.Errorf("Mkdir: %v", err))
		}
	}
	cacheFilePath := cacheDirPath + "/" + cacheFileName
	_, err := os.Stat(cacheFilePath)
	cacheFilePresent := err == nil
	if !cacheFilePresent {
		c, err := cache.Generate()
		if err != nil {
			panic(fmt.Errorf("emoji.GenerateCache: %v", err))
		}
		cacheFile, err := os.Create(cacheFilePath)
		if err != nil {
			panic(fmt.Errorf("os.Create: %v", err))
		}
		defer cacheFile.Close()
		cache.Export(cacheFile, &c)
	}
}
