package toolformation

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os/exec"

	"github.com/goark/errs"
	"github.com/goccy/go-yaml"
)

type Homebrew struct {
	formula []string
	cask    []string
}

type ToolFormation struct {
	PackageManager string `yaml:"package-manager"`
	Homebrew       `yaml:"homebrew"`
}

func execute(cmd string) int {
	c := exec.Command("/bin/bash", "-c", cmd)
	stderr, _ := c.StderrPipe()
	c.Start()

	scanner := bufio.NewScanner(stderr)
	for scanner.Scan() {
		msg := scanner.Text()
		fmt.Println(msg)
	}
	c.Wait()

	return c.ProcessState.ExitCode()
}

func (h *Homebrew) Execute() {
	for i := 0; i < len(h.formula); i++ {
		cmd := fmt.Sprintf("brew install %s", h.formula[i])
		code := execute(cmd)
		if code != 0 {
			fmt.Printf("Installation of %s[formula] failed\n", h.formula[i])
		}
	}

	for i := 0; i < len(h.cask); i++ {
		cmd := fmt.Sprintf("brew install %s", h.cask[i])
		code := execute(cmd)
		if code != 0 {
			fmt.Printf("Installation of %s[cask] failed\n", h.formula[i])
		}
	}
}

func ReadConfig(path string) (*ToolFormation, error) {
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

func (t *ToolFormation) Execute() {

	if t.PackageManager == "homebrew" {
		t.Homebrew.Execute()
	}
}
