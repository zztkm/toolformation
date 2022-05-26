package toolformation

import (
	"context"
	"os"
	"strings"

	"mvdan.cc/sh/v3/expand"
	"mvdan.cc/sh/v3/interp"
	"mvdan.cc/sh/v3/syntax"
)

// RunCommand runs a shell command
func RunCommand(cmd string) error {
	p, err := syntax.NewParser().Parse(strings.NewReader(cmd), "")
	if err != nil {
		return err
	}

	env := os.Environ()

	r, err := interp.New(
		interp.Params("-e"),
		interp.Env(expand.ListEnviron(env...)),
		interp.StdIO(os.Stdin, os.Stdout, os.Stderr),
	)
	if err != nil {
		return err
	}

	return r.Run(context.TODO(), p)
}
