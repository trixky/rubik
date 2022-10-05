package main

import (
	"github.com/trixky/rubik/algo"
	"github.com/trixky/rubik/input"
	"github.com/trixky/rubik/server"
)

func main() {
	set, api_mode := input.Get()

	if api_mode {
		server.Start()
	} else {
		result := algo.Resolve(set)
		result.Print()
	}
}
