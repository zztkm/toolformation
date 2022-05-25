package toolformation

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/goark/errs"
)

// Homebrew
type Homebrew struct {
	Formula []string `yaml:"formula"`
	Cask    []string `yaml:"cask"`
}

func InstallHomebrew() error {
	m := getUnameMachine()
	var c1, c2 string
	if m == "arm64" {
		c1 = `echo 'eval "$(/opt/homebrew/bin/brew shellenv)"' >> $HOME/.zprofile`
		c2 = `eval "$(/opt/homebrew/bin/brew shellenv)"`
	} else {
		c1 = `echo 'eval "$(/usr/local/bin/brew shellenv)"' >> $HOME/.zprofile`
		c2 = `eval "$(/usr/local/bin/brew shellenv)"`
	}
	cmds := []string{
		`/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"`,
		c1,
		c2,
	}
	for i := 0; i < len(cmds); i++ {
		err := RunCommand(cmds[i])
		if err != nil {
			color.Red("Failed to install homebrew")
			return err
		}
	}
	return nil
}

// Install homwbrew
func (h *Homebrew) Install() error {
	if code := check("brew"); code != 0 {
		fmt.Println("homebrew was not installed")
		err := InstallHomebrew()
		if err != nil {
			return err
		}
	}

	// Install formula
	for i := 0; i < len(h.Formula); i++ {
		cmd := fmt.Sprintf("brew install %s", h.Formula[i])
		err := RunCommand(cmd)
		if err != nil {
			fmt.Printf("Installation of %s[formula] failed\n", h.Formula[i])
			fmt.Println(err)
		}
	}

	// Install cask
	for i := 0; i < len(h.Cask); i++ {
		cmd := fmt.Sprintf("brew install --cask %s", h.Cask[i])
		err := RunCommand(cmd)
		if err != nil {
			fmt.Printf("Installation of %s[cask] failed\n", h.Cask[i])
			fmt.Println(err)
		}
	}
	return nil
}

// Uninstall homwbrew
func (h *Homebrew) Uninstall() error {
	if code := check("brew"); code != 0 {
		return errs.New("homebrew was not installed")
	}

	for i := 0; i < len(h.Formula); i++ {
		cmd := fmt.Sprintf("brew uninstall %s", h.Formula[i])
		err := RunCommand(cmd)
		if err != nil {
			fmt.Printf("Uninstallation of %s[formula] failed\n", h.Formula[i])
			fmt.Println(err)
		}
	}

	for i := 0; i < len(h.Cask); i++ {
		cmd := fmt.Sprintf("brew uninstall %s", h.Cask[i])
		err := RunCommand(cmd)
		if err != nil {
			fmt.Printf("Uninstallation of %s[cask] failed\n", h.Cask[i])
			fmt.Println(err)
		}
	}
	return nil
}
