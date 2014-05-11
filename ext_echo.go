package main

import (
	"fmt"
	"strings"
)

func init() {
	var s = spec{
		func(inp string) bool {
			lowerInp := strings.ToLower(inp)
			return strings.HasPrefix(lowerInp, "echo")
		},

		func(inp string) {
			toEcho := inp[len("echo "):len(inp)]
			echo(toEcho)
		},

		"Echoes provided string back to you.",

		[]string{
			"echo <something you want to be echoed back to you>",
		},
	}
	specList = append(specList, s)
}

// echo function simply prints the user entered input as it is, and displays it
// back to the user on the the interactive-mode.
func echo(text string) {
	fmt.Println(text)
}
