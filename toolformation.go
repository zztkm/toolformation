package toolformation

import (
	"fmt"
	"io/ioutil"
	"os/exec"

	"github.com/fatih/color"
	"github.com/goark/errs"
	"github.com/goccy/go-yaml"
)

// Homebrew
type Homebrew struct {
	Formula []string `yaml:"formula"`
	Cask    []string `yaml:"cask"`
}

// ToolFormation
type ToolFormation struct {
	PackageManager string `yaml:"package-manager"`
	Homebrew       `yaml:"homebrew"`
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

// Install homwbrew
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
		cmd := fmt.Sprintf("brew install --cask %s", h.Cask[i])
		err := RunCommand(cmd)
		if err != nil {
			fmt.Printf("Installation of %s[cask] failed\n", h.Cask[i])
			fmt.Println(err)
		}
	}
}

// Uninstall homwbrew
func (h *Homebrew) Uninstall() {
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
	if t.PackageManager == "homebrew" {
		if code := check("brew"); code != 0 {
			fmt.Println("homebrew was not installed")
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
					return
				}
			}
		}
		t.Homebrew.Install()
	} else {
		color.HiBlue("Currently only `homebrew` is supported")
		color.HiBlue("Please send me an issue or pr!")
		color.HiBlue("https://github.com/zztkm/toolformation/issues")
		return
	}
}
