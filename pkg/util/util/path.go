package util

import (
	"io"
	"io/ioutil"
	"os"
)

func CanCreate(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return false
	}
	var d []byte
	if err := ioutil.WriteFile(path, d, 0644); err == nil {
		os.Remove(path)
		return true
	}

	return false
}

func CanReadFrom(path string) bool {
	file, err := os.Open(path)
	if err != nil {
		return false
	}
	if _, err := io.ReadAll(file); err != nil {
		return false
	}
	return true
}
