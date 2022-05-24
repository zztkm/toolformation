package toolformation

import (
	"fmt"
	"io/ioutil"

	"github.com/goark/errs"
	"github.com/goccy/go-yaml"
)

// Homebrew
type Homebrew struct {
	Formula []string
	Cask    []string
}

// ToolFormation
type ToolFormation struct {
	PackageManager string `yaml:"package-manager"`
	Homebrew       `yaml:"homebrew"`
}

func (h *Homebrew) Install() {
	for i := 0; i < len(h.Formula); i++ {
		cmd := fmt.Sprintf("brew install %s", h.Formula[i])
		err := RunCommand(cmd)
		if err != nil {
			fmt.Printf("Installation of %s[formula] failed\n", h.Formula[i])
			fmt.Println(err)
		}
	}

	for i := 0; i < len(h.Cask); i++ {
		cmd := fmt.Sprintf("brew install %s", h.Cask[i])
		err := RunCommand(cmd)
		if err != nil {
			fmt.Printf("Installation of %s[cask] failed\n", h.Cask[i])
			fmt.Println(err)
		}
	}
}

//
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

	if t.PackageManager == "homebrew" {
		t.Homebrew.Install()
	}
}
