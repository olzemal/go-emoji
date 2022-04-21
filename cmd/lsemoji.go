package main

import (
	"fmt"
	"os"

	emoji "github.com/olzemal/lsemoji"
)

const (
	cacheFileName = "cache"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Errorf("os.UserHomeDir: %v", err))
	}
	cacheDirPath := home + "/.cache/lsemoji"
	initCache(cacheDirPath)

	cache, err := emoji.ImportCache(cacheDirPath + "/" + cacheFileName)
	if err != nil {
		panic(fmt.Errorf("Emoji.ImportCache: %v", err))
	}

	if len(os.Args) < 2 {
		for key, value := range cache {
			fmt.Printf("%s=%s\n", key, value)
		}
	} else {

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
		cache, err := emoji.GenerateCache()
		if err != nil {
			panic(fmt.Errorf("emoji.GenerateCache: %v", err))
		}
		cacheFile, err := os.Create(cacheFilePath)
		if err != nil {
			panic(fmt.Errorf("os.Create: %v", err))
		}
		defer cacheFile.Close()
		emoji.ExportCache(cacheFile, &cache)
	}
}
