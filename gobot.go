/*
 *   ▄████  ▒█████   ▄▄▄▄    ▒█████  ▄▄▄█████▓
 *  ██▒ ▀█▒▒██▒  ██▒▓█████▄ ▒██▒  ██▒▓  ██▒ ▓▒
 * ▒██░▄▄▄░▒██░  ██▒▒██▒ ▄██▒██░  ██▒▒ ▓██░ ▒░
 * ░▓█  ██▓▒██   ██░▒██░█▀  ▒██   ██░░ ▓██▓ ░
 * ░▒▓███▀▒░ ████▓▒░░▓█  ▀█▓░ ████▓▒░  ▒██▒ ░
 * ░▒   ▒ ░ ▒░▒░▒░ ░▒▓███▀▒░ ▒░▒░▒░   ▒ ░░
 *   ░   ░   ░ ▒ ▒░ ▒░▒   ░   ░ ▒ ▒░     ░
 * ░ ░   ░ ░ ░ ░ ▒   ░    ░ ░ ░ ░ ▒    ░
 *       ░     ░ ░   ░          ░ ░
 *                        ░
 *
 * Interactive Bot, Gobot!
 *
 * Author : Vishal Telangre
 * Source : http://github.com/vishaltelangre/gobot
 * License: MIT
 *
 */

package main

import (
	"flag"
	"fmt"
	"github.com/vishaltelangre/gobot/readline"
	"os"
	"regexp"
	"strings"
)

// spec stores functions for condition and logic to invoke if that condition is
// true.
type spec struct {
	// 'matchCond' function checks the condition, returns boolean.
	matchCond func(string) bool

	// 'invoker' function contains logic to execute if 'matchCond' is true.
	invoker func(string)

	// 'explanation' is the description about the result when 'invoker' executes
	explanation string

	// 'commands' is a list of commands which can be used by the user
	commands []string
}

var (
	// input is the input string entered by the user on stdio.
	input string

	// id is the command-line option can be passed in with the 'Gobot'
	// command to set the name of the bot, which is displayed on stdio
	// while user interacts with it. Default name of 'Gobot' is set as
	// "Gobo".
	id = flag.String("id", "Gobo", "name your bot")

	// exitCommands are the different kinds of the commands which user can
	// enter in the interactive-mode to quit from it.
	exitCommands = []string{"quit", "exit", "bye", "q"}

	// specList is a slice of the 'spec' struts to store multiple
	// extension-specific list of condition and logic functions.
	specList = []spec{}

	// -i Interactive
	// -c Command string
	isInteractive    *bool = flag.Bool("i", true, "Specify an interactive shell")
	useCommandString *bool = flag.Bool("c", false, "Run commands from the command string")
)

// handler is the final stage function where all the logic is handled. It
// visualizes the REPL-like interactive-mode on stdio and facilitates user to
// specify the kind of commands. If the user command is known or recognised,
// then it's evaluated, needed processing is done in background, and if
// necessary then some textual information is displayed in the same interactive-
// mode as a response to the user's input. And, the cursor is handed-over back
// to the user, so (s)he can continue to interact with the 'Gobot' forever.
func interactiveQueryHandler() {
	for {
		// displays the bot's specified name on the left of stdio screen
		botName := fmt.Sprintf("\x1b[01;33m%s > \x1b[0m", *id)
		inpLine := readline.ReadLine(&botName)
		if inpLine == nil {
			fmt.Println("There seems some issue.")
			continue
		}
		input = strings.TrimSpace(*inpLine)
		if input != "" {
			readline.AddHistory(input)
		}
		for _, spec := range specList {
			if spec.matchCond(input) {
				spec.invoker(input)
				break
			}
		}
	}
}

func cliQueryHandler(input string) {
	for _, spec := range specList {
		if spec.matchCond(input) {
			spec.invoker(input)
			break
		}
	}
}

func init() {
	var defaultSpecs = []spec{
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

func main() {
	// "name" is another command-line option to allow user to set the bot's name.
	flag.StringVar(id, "name", "Gobo", "name your gobot")

	flag.Parse()

	if *useCommandString {
		inp := strings.Join(flag.Args(), " ")
		if trimmedInp := strings.TrimSpace(inp); trimmedInp != "" {
			cliQueryHandler(trimmedInp)
			return
		}
		flag.Usage()
	} else if *isInteractive {
		fmt.Printf("Howdy, greetings from %s.\n", *id)
		interactiveQueryHandler()
	} else {
		flag.Usage()
	}
}
