package extensions

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var (
	clrscrCommands = []string{"clear", "clear screen", "clr", "clrscr"}
	exitCommands   = []string{"quit", "exit", "bye", "q"}
)

// echo function simply prints the user entered input as it is, and displays it
// back to the user on the the interactive-mode.
func echo(text string) { fmt.Println(text) }

// containsString is the function to lookup the slice of strings for an passd
// item; if the item is found in the slice, then it returns true, neither it
// returns false value.
func containsString(slice []string, item string) bool {
	for _, value := range slice {
		if value == item {
			return true
		}
	}
	return false
}

// execCmdAndPrintResult is used to execute system commands and prints the
// result of the command from stdio to the interactive session.
func execCmdAndPrintResult(cmdWithArg string) {
	var cmdFields = strings.Fields(cmdWithArg)
	cmd := exec.Command(cmdFields[0], cmdFields[1:]...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("Don't mind, something's broken!", "\n", err)
	}
	fmt.Println(strings.TrimSpace(out.String()))
}

func init() {
	var defaultSpecs = []spec{
		{
			func(inp string) bool {
				lowerInp := strings.ToLower(inp)
				return containsString(clrscrCommands, lowerInp)
			},

			func(inp string) {
				c := exec.Command("clear")
				c.Stdout = os.Stdout
				c.Run()
			},

			"Clear the screen.",

			clrscrCommands,
		},

		{
			func(inp string) bool {
				lowerInp := strings.ToLower(inp)
				return strings.HasPrefix(lowerInp, "echo")
			},

			func(inp string) {
				toEcho := inp[len("echo "):len(inp)]
				echo(toEcho)
			},

			"Echo provided string back to you.",

			[]string{
				"echo <something you want to be echoed back to you>",
			},
		},

		{
			func(inp string) bool {
				lowerInp := strings.ToLower(inp)
				return containsString(exitCommands, lowerInp)
			},

			func(inp string) {
				fmt.Println("Good bye!")
				os.Exit(0)
			},

			"Quit from the interactive session.",

			exitCommands,
		},

		{
			func(inp string) bool {
				lowerInp := strings.TrimSpace(strings.ToLower(inp))
				re := regexp.MustCompile(`^help(\s+me)?$`)
				return re.MatchString(lowerInp) || lowerInp == "?"
			},

			func(inp string) {
				for i, spec := range specList {
					fmt.Printf("%d.1 %s\n", i+1, spec.explanation)
					fmt.Printf("%d.2 Commands:\n", i+1)
					for _, command := range spec.commands {
						fmt.Printf("  * %s\n", command)
					}
				}
			},

			"Print nice help.",

			[]string{
				"help me",
				"help",
				"?",
			},
		},
	}

	for _, spec := range defaultSpecs {
		specList = append(specList, spec)
	}
}
