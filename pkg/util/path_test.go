package util_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/olzemal/go-emoji/pkg/util"
)

const (
	cacheFileName = "cache"
)

func TestSetup(t *testing.T) {
	tmpDirPath := SetupTestTmp()
	defer CleanUp(tmpDirPath)
	if !util.CanCreate(tmpDirPath + "/" + cacheFileName) {
		t.Errorf("TestSetup: cannot create `%s`", tmpDirPath+"/"+cacheFileName)
	}
}

func TestReadFrom(t *testing.T) {
	tmpDirPath := SetupTestTmp()
	defer CleanUp(tmpDirPath)
	f, err := os.Create(tmpDirPath + "/" + cacheFileName)
	if err != nil {
		t.Errorf("TestReadFrom: cannot create `%s`", tmpDirPath+"/"+cacheFileName)
	}
	defer f.Close()
	if !util.CanReadFrom(tmpDirPath + "/" + cacheFileName) {
		t.Errorf("TestReadFrom: cannot read from `%s`", tmpDirPath+"/"+cacheFileName)
	}
}

func SetupTestTmp() string {
	path, err := os.MkdirTemp("/tmp", "go-emoji_util_test_")
	if err != nil {
		panic(fmt.Errorf("SetupTestTemp: %v", err))
	}
	log.Println("Created: ", path)
	return path
}

func CleanUp(tmpDir string) {
	if err := os.RemoveAll(tmpDir); err != nil {
		panic(fmt.Errorf("CleanUp: %v", err))
	}
	log.Println("Removed: ", tmpDir)
}
