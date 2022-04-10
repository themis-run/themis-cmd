package cmd

import (
	"errors"
	"sync"
)

type Command string
type CommandFunc func(args ...string) Result

var (
	commands map[Command]CommandFunc
	mu       sync.Mutex
)

var errorCommandNotFound = errors.New("command not found!")

var CommandNotFound = func(...string) Result {
	return NewErrorResult(errorCommandNotFound)
}

func Register(cmd Command, f CommandFunc) {
	mu.Lock()
	defer mu.Unlock()

	if commands == nil {
		commands = make(map[Command]CommandFunc)
	}

	commands[cmd] = f
}

func Get(cmd Command) CommandFunc {
	if c, ok := commands[cmd]; ok {
		return c
	}

	return CommandNotFound
}
