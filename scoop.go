package toolformation

import "fmt"

type Scoop struct {
	Main   []string `yaml:"main"`
	Extras []string `yaml:"extras"`
}

func (s Scoop) Install() error {

	if len(s.Extras) != 0 {
		cmd := fmt.Sprintf("scoop bucket add extras")
		err := RunCommand(cmd)
		if err != nil {
			return err
		}
	}

	// Install Main
	for i := 0; i < len(s.Main); i++ {
		cmd := fmt.Sprintf("scoop install %s", s.Main[i])
		err := RunCommand(cmd)
		if err != nil {
			fmt.Printf("Installation of %s[formula] failed\n", s.Main[i])
			fmt.Println(err)
		}
	}

	// Install Extras
	for i := 0; i < len(s.Extras); i++ {
		cmd := fmt.Sprintf("scoop install %s", s.Extras[i])
		err := RunCommand(cmd)
		if err != nil {
			fmt.Printf("Installation of %s[cask] failed\n", s.Extras[i])
			fmt.Println(err)
		}
	}
	return nil
}

func (s Scoop) Uninstall() error { return nil }

func (s Scoop) Upgrade() error { return nil }
