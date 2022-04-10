package cmd

import (
	"fmt"
)

func PrintResult(res Result) {
	if res == nil {
		return
	}

	fmt.Println(res.String())
}

func PrintRowHead() {
	fmt.Print(">")
}

func PrintCmd(cmd string) {
	fmt.Print(cmd)
}
