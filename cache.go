package toolformation

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/goark/errs"
)

const (
	cacheDir  = "toolformation"
	cacheFile = "toolformation-cache.yml"
)

func CacheDir() string {
	var path string
	if b := os.Getenv("APPDATA"); runtime.GOOS == "windows" && b != "" {
		path = filepath.Join(b, cacheDir)
	} else {
		c, _ := os.UserHomeDir()
		path = filepath.Join(c, ".config", cacheDir)
	}

	return path
}

func CreateCache(c *Config) error {
	p := CacheDir()
	if dirExists(p) {
		return nil
	}
	err := os.MkdirAll(p, os.ModePerm)
	if err != nil {
		return err
	}
	p = filepath.Join(p, cacheFile)
	err = WriteFile(c, p)
	if err != nil {
		return nil
	}
	return nil
}

func UpdateCache(c *Config) error {
	p := CacheDir()
	if !dirExists(p) {
		return nil
	}
	p = filepath.Join(p, cacheFile)
	err := WriteFile(c, p)
	if err != nil {
		return nil
	}
	return nil
}

func ReadCache() (*Config, error) {
	p := CacheDir()
	if !dirExists(p) {
		return nil, errs.New("Cache directory does not exist")
	}
	p = filepath.Join(p, cacheFile)
	return ReadFile(p)
}
