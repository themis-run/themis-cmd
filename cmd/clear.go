package cmd

import (
	"errors"
	"os"
	"os/exec"
	"runtime"
)

var clear map[string]*exec.Cmd

func init() {
	clear = make(map[string]*exec.Cmd)
	clear["linux"] = exec.Command("clear")
	clear["windows"] = exec.Command("cmd", "/c", "cls")
	clear["darwin"] = exec.Command("clear")

	Register("clear", Clear)
}

type ClearResult struct {
	err error
}

func NewClearResult(err error) *ClearResult {
	return &ClearResult{
		err: err,
	}
}

func (c *ClearResult) String() string {
	return c.err.Error()
}

var platformUnSupport = errors.New("Your platform is unsupported! I can't clear terminal screen :(")

func Clear(args ...string) Result {
	c, ok := clear[runtime.GOOS]
	if !ok {
		return NewClearResult(platformUnSupport)
	}

	c.Stdout = os.Stdout
	if err := c.Run(); err != nil {
		return NewClearResult(err)
	}

	return nil
}
