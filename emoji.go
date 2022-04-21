package emoji

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	util "github.com/olzemal/lsemoji/util"
	"golang.org/x/net/html"
)

type cache map[string]string

const (
	emojiUrl = "https://unicode.org/emoji/charts/full-emoji-list.html"
)

func GenerateCache() (cache, error) {
	resp, err := http.Get(emojiUrl)
	if err != nil {
		return nil, err
	}

	c := make(cache)
	z := html.NewTokenizer(resp.Body)
	end := false
	for !end {
		tt := z.Next()
		t := z.Token()
		if tt == html.EndTagToken && t.Data == "html" {
			end = true
			break
		}
		if tt != html.StartTagToken || t.Data != "tr" {
			continue
		}
		name := ""
		chars := ""
		for !(t.Data == "tr" && t.Type == html.EndTagToken) {
			z.Next()
			t = z.Token()
			if t.Data == "html" {
				end = true
				break
			}
			// Check for name / chars cell
			if t.Data == "td" &&
				strings.Contains(t.String(), "class=\"name\"") &&
				z.Next() == html.TextToken {
				name = string(z.Token().Data)
			} else if t.Data == "td" &&
				strings.Contains(t.String(), "class=\"chars\"") &&
				z.Next() == html.TextToken {
				chars = string(z.Token().Data)
			}
		}
		if len(name) > 0 && len(chars) > 0 {
			c[util.Format(name)] = chars
		}
	}
	return c, nil
}

func ImportCache(path string) (cache, error) {
	if !util.CanReadFrom(path) {
		return nil, fmt.Errorf("Invalid path: `%s`", path)
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(bytes), "\n")
	c := make(cache)
	for _, line := range lines {
		pair := strings.SplitN(line, "=", 2)
		if len(pair) == 2 {
			c[pair[0]] = pair[1]
		}
	}

	return c, nil
}

func ExportCache(path string, cache cache) error {
	for key, val := range cache {
		fmt.Printf("%s=%s\n", key, val)
	}
	return nil
}

func FromString(name string) (string, error) {
	return "", nil
}