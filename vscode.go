package toolformation

import (
	"fmt"

	"github.com/goark/errs"
)

type VSCode struct {
	Extension []string `yaml:"extension"`
}

func (v VSCode) Install() error {
	if code := Check("code"); code != 0 {
		return errs.New("code was not installed")
	}

	// Install formula
	for i := 0; i < len(v.Extension); i++ {
		cmd := fmt.Sprintf("code --install-extension %s", v.Extension[i])
		err := RunCommand(cmd)
		if err != nil {
			fmt.Printf("Installation of %s[vscode extension] failed\n", v.Extension[i])
			fmt.Println(err)
		}
	}
	return nil
}

func (v VSCode) Uninstall() error { return nil }

func (v VSCode) Upgrade() error { return nil }
