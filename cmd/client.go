package cmd

import (
	"bufio"
	"os"
	"strings"
)

type ThemisCmd struct {
}

func (c *ThemisCmd) Start() {
	for {
		PrintRowHead()
		var cmd string
		var err error
		if cmd, err = bufio.NewReader(os.Stdin).ReadString('\n'); err != nil {
			return
		}
		command, args := c.parseCommand(cmd)
		c.doCommand(command, args)
	}
}

var commandSeparator = " "

func (c *ThemisCmd) parseCommand(text string) (string, []string) {
	s := strings.Split(text, commandSeparator)
	if len(s) <= 0 {
		return "", nil
	}

	str := make([]string, 0)
	for _, v := range s {
		str = append(str, strings.TrimSpace(v))
	}

	return str[0], str[1:]
}

func (c *ThemisCmd) doCommand(command string, args []string) {
	f := Get(Command(command))
	result := f(args...)
	PrintResult(result)
}
