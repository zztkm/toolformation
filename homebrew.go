package toolformation

import "fmt"

// Homebrew
type Homebrew struct {
	Formula []string `yaml:"formula"`
	Cask    []string `yaml:"cask"`
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
