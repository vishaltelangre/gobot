package main

import (
	"os"
	"os/exec"
	"strings"
)

// clrscrCommands is a set of commands, any of which user can enter to clear the
// stdio screen.
var clrscrCommands = []string{"clear", "clear screen", "clr", "clrscr"}

func init() {
	var s = spec{
		func(inp string) bool {
			lowerInp := strings.ToLower(inp)
			return containsString(clrscrCommands, lowerInp)
		},

		func(inp string) {
			// execute the os-specific clear screen command.
			c := exec.Command("clear")
			c.Stdout = os.Stdout
			c.Run()
		},

		"Clears the screen.",

		clrscrCommands,
	}
	specList = append(specList, s)
}
