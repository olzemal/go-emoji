package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/olzemal/go-emoji/pkg/cache"
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

	stdin := bufio.NewScanner(os.Stdin)
	emoji, _ := regexp.Compile("\\:[a-z_]+\\:")
	for stdin.Scan() {
		line := stdin.Text()
		replaced := emoji.ReplaceAllStringFunc(line, func(match string) string {
			e, ok := c[strings.Trim(match, ":")]
			if !ok {
				return match
			}
			return e
		})
		fmt.Println(replaced)
	}
}
