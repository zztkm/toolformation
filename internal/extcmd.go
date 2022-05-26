package internal

import (
	"fmt"
	"os/exec"
	"strings"
)

// check return `type name` exit code
func Check(name string) int {
	s := fmt.Sprintf("type %s", name)
	cmd := exec.Command("/bin/bash", "-c", s)
	err := cmd.Run()
	if err != nil {
		return 1
	}

	return cmd.ProcessState.ExitCode()
}

// GetUnameMachine
// not work on windows
func GetUnameMachine() string {
	s := fmt.Sprintf("/usr/bin/uname -m")
	out, err := exec.Command("/bin/bash", "-c", s).Output()
	if err != nil {
		return ""
	}
	return strings.Trim(string(out), "\n")
}
