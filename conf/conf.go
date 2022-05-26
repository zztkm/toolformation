package conf

import (
	"io/ioutil"

	"github.com/fatih/color"
	"github.com/goark/errs"
	"github.com/goccy/go-yaml"
	"github.com/zztkm/toolformation/pkgmgr"
)

// ToolFormation
type ToolFormation struct {
	PackageManagerName string `yaml:"package-manager"`
	pkgmgr.Homebrew    `yaml:"homebrew"`
}

func New(path string) (*ToolFormation, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return &ToolFormation{}, errs.Wrap(err)
	}
	var t ToolFormation
	err = yaml.Unmarshal(b, &t)
	if err != nil {
		return &ToolFormation{}, errs.Wrap(err)
	}
	return &t, nil
}

func (t *ToolFormation) Install() {
}

type PackageManager interface {
	Install() error
	Uninstall() error
	Upgrade() error
}

func (t *ToolFormation) NewPackageManager() PackageManager {
	if t.PackageManagerName == "homebrew" {
		hb := t.Homebrew
		return hb
	} else {
		color.HiBlue("Currently only `homebrew` is supported")
		color.HiBlue("Please send me an issue or pr!")
		color.HiBlue("https://github.com/zztkm/toolformation/issues")
		return nil
	}
}

func Install(p PackageManager) error {
	return p.Install()
}

func Uninstall(p PackageManager) error {
	return p.Uninstall()
}

func Upgrade(p PackageManager) error {
	return p.Upgrade()
}
