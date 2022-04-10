package main

import (
	"go.themis.run/themis-cmd/cmd"
	"go.themis.run/themis-cmd/conf"
)

func main() {
	conf.Setup()
	c := &cmd.ThemisCmd{}
	c.Start()
}
