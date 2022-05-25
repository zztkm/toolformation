package toolformation

import (
	"fmt"
	"io/ioutil"
	"os/exec"

	"github.com/fatih/color"
	"github.com/goark/errs"
	"github.com/goccy/go-yaml"
)

// ToolFormation
type ToolFormation struct {
	PackageManagerName string `yaml:"package-manager"`
	Homebrew           `yaml:"homebrew"`
}

// check return `type name` exit code
func check(name string) int {
	s := fmt.Sprintf("type %s", name)
	cmd := exec.Command("/bin/bash", "-c", s)
	err := cmd.Run()
	if err != nil {
		return 1
	}

	return cmd.ProcessState.ExitCode()
}

func getUnameMachine() string {
	s := fmt.Sprintf("/usr/bin/uname -m")
	out, err := exec.Command("/bin/bash", "-c", s).Output()
	if err != nil {
		return ""
	}
	return string(out)
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
		return t
	} else {
		color.HiBlue("Currently only `homebrew` is supported")
		color.HiBlue("Please send me an issue or pr!")
		color.HiBlue("https://github.com/zztkm/toolformation/issues")
		return
	}
}
