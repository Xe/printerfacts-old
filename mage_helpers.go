// +build mage

package main

import (
	"context"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/jtolds/qod"
	"github.com/pkg/errors"
)

var wd string

func init() {
	lwd, err := os.Getwd()
	qod.ANE(err)

	wd = lwd
}

func output(cmd string, args ...string) (string, error) {
	c := exec.Command(cmd, args...)
	c.Env = os.Environ()
	c.Stderr = os.Stderr
	b, err := c.Output()
	if err != nil {
		return "", errors.Wrapf(err, `failed to run %v %q`, cmd, args)
	}
	return string(b), nil
}

func gitTag() (string, error) {
	s, err := output("git", "describe", "--tags")
	if err != nil {
		ee, ok := errors.Cause(err).(*exec.ExitError)
		if ok && ee.Exited() {
			// probably no git tag
			return "dev", nil
		}
		return "", err
	}

	return strings.TrimSuffix(s, "\n"), nil
}

func shouldWork(ctx context.Context, env []string, dir string, cmdName string, args ...string) {
	loc, err := exec.LookPath(cmdName)
	qod.ANE(err)

	cmd := exec.CommandContext(ctx, loc, args...)
	cmd.Dir = dir
	cmd.Env = append(env, os.Environ()...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Printf("starting process, env: %v, pwd: %s, cmd: %s, args: %v", env, dir, loc, args)
	err = cmd.Run()
	qod.ANE(err)
}
