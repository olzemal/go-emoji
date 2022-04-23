package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/olzemal/go-emoji/pkg/cache"
	"github.com/olzemal/go-emoji/pkg/config"
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

	stdin := bufio.NewScanner(os.Stdin)
	emoji, _ := regexp.Compile("\\:[a-z_]+\\:")
	for stdin.Scan() {
		line := stdin.Text()
		replaced := emoji.ReplaceAllStringFunc(line, func(match string) string {
			name := strings.Trim(match, ":")
			s := cfg.EmojiFromAlias(name)
			if len(s) == 0 {
				s = name
			}
			e, ok := c[s]
			if !ok {
				return match
			}
			return e
		})
		fmt.Println(replaced)
	}
}
