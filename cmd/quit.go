package cmd

import "os"

func init() {
	Register("quit", Quit)
	Register("q", Quit)
}

func Quit(args ...string) Result {
	if len(args) != 0 {
		return CommandNotFound()
	}

	os.Exit(1)
	return nil
}
