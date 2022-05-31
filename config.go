package toolformation

import (
	"bytes"
	"io"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/goark/errs"
	"github.com/goccy/go-yaml"
)

var defaultFileNames = []string{"ToolFormation.yml", "ToolFormation.yaml"}

type Config struct {
	PackageManagerName string `yaml:"package-manager"`
	Homebrew           `yaml:"homebrew"`
	Scoop              `yaml:"scoop"`
}

func Write(c *Config, w io.Writer) error {
	bytes, err := yaml.Marshal(&c)
	if err != nil {
		return err
	}
	_, err = w.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}

// Read
func Read(r io.Reader) (*Config, error) {
	buf := bytes.Buffer{}
	_, err := buf.ReadFrom(r)
	if err != nil {
		return nil, errs.Wrap(err)
	}
	b := buf.Bytes()
	var t Config
	err = yaml.Unmarshal(b, &t)
	if err != nil {
		return nil, errs.Wrap(err)
	}
	return &t, nil
}

func WriteFile(c *Config, path string) error {
	f, err := createFile(path)
	defer f.Close()
	if err != nil {
		return errs.Wrap(err)
	}
	return Write(c, f)
}

// ReadFile
func ReadFile(path string) (*Config, error) {
	f, err := openFile(path)
	defer f.Close()
	if err != nil {
		return nil, errs.Wrap(err)
	}
	return Read(f)
}

// DefaultConfigPath
func DefaultConfigPath() (string, error) {

	for i := 0; i < len(defaultFileNames); i++ {
		d, err := os.Getwd()
		if err != nil {
			return "", err
		}
		defaultFileNames[i] = filepath.Join(d, defaultFileNames[i])
	}

	var fileCheck func(filenames []string) string
	fileCheck = func(filenames []string) string {
		for i := 0; i < len(filenames); i++ {
			if fileExists(filenames[i]) {
				return filenames[i]
			}
		}
		return ""
	}
	file := fileCheck(defaultFileNames)

	return file, nil
}

func ParseDefaultConfigFile() (*Config, error) {
	p, err := DefaultConfigPath()
	if err != nil {
		return nil, err
	}
	return ReadFile(p)
}

type PackageManager interface {
	Install() error
	Uninstall() error
	Upgrade() error
}

func (c *Config) NewPackageManager() PackageManager {
	if c.PackageManagerName == "homebrew" {
		return c.Homebrew
	} else if c.PackageManagerName == "scoop" {
		return c.Scoop
	} else {
		color.HiBlue("Currently only `homebrew` is supported")
		color.HiBlue("Please send me an issue or pr!")
		color.HiBlue("https://github.com/zztkm/toolformation/issues")
		return nil
	}
}

// 未実装
func (c *Config) DifferenceCheck(cc *Config) {}

func Install(p PackageManager) error {
	return p.Install()
}

func Uninstall(p PackageManager) error {
	return p.Uninstall()
}

func Upgrade(p PackageManager) error {
	return p.Upgrade()
}

func createFile(path string) (*os.File, error) {
	return os.Create(path)
}

func openFile(path string) (*os.File, error) {
	return os.Open(path)
}

func dirExists(path string) bool {
	f, err := os.Stat(path)
	return err == nil && f.IsDir()
}

func fileExists(path string) bool {
	f, err := os.Stat(path)
	return err == nil && !f.IsDir()
}
