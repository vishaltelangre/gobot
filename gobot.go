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
	"github.com/vishaltelangre/gobot/extensions"
	"github.com/vishaltelangre/gobot/readline"
	"strings"
)

const VERSION = "0.0.1"

var (
	input string // input is the input string entered by the user on stdio.

	isInteractive    *bool   = flag.Bool("i", false, "Start Gobot in interactive mode")
	useCommandString *bool   = flag.Bool("c", true, "Make Gobot to execute commands directly from shell")
	id               *string = flag.String("name", "Gobo", "Name your bot")
	showVersion      *bool   = flag.Bool("version", false, "Show version info")
	showHelp         *bool   = flag.Bool("help", false, "Show help and usage of gobot command")
)

func printCommandUsage() {
	fmt.Println("gobot is a bot!\n")
	fmt.Println("Usage:")
	fmt.Println("\t gobot [options] [command string]\n")
	fmt.Println("The options are:")
	fmt.Println("\t -i        \t start gobot interactive-mode")
	fmt.Println("\t --name    \t name of gobot appear in interactive-mode")
	fmt.Println("\t --version \t displays version of gobot")
	fmt.Println("\t --help    \t show this information")
	fmt.Println("\n The command string can be any of following:\n")
	cliQueryHandler("help me")
}

func cliQueryHandler(input string) { extensions.EvalQuery(input) }

func interactiveQueryHandler() {
	for {
		// displays the bot's specified name on the left of stdio screen
		botName := fmt.Sprintf("\x1b[01;33m%s > \x1b[0m", *id)
		inpLine := readline.ReadLine(&botName)

		if inpLine == nil {
			fmt.Println("There seems some issue.")
			continue
		}

		if input = strings.TrimSpace(*inpLine); input != "" {
			readline.AddHistory(input)
		}

		extensions.EvalQuery(input)
	}
}

func main() {
	flag.Parse()

	if *showVersion {
		fmt.Printf("Gobot - v%s\n", VERSION)
		return
	}

	if *showHelp {
		printCommandUsage()
		return
	}

	if *isInteractive {
		fmt.Printf("Howdy, greetings from %s.\n", *id)
		interactiveQueryHandler()
		return
	}

	if *useCommandString {
		inp := strings.Join(flag.Args(), " ")
		if trimmedInp := strings.TrimSpace(inp); trimmedInp != "" {
			cliQueryHandler(trimmedInp)
			return
		}
	}

	printCommandUsage()
}
