package cmd

import (
	"fmt"
	"io/ioutil"
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

func PrintBanner() {
	f, err := ioutil.ReadFile("./banner.txt")
	if err != nil {
		fmt.Println("banner.txt not found!")
		return
	}

	fmt.Println(string(f))
}
