package main

import (
	"regexp"
	"strings"
)

func init() {
	var s = spec{
		func(inp string) bool {
			var lowerInp = strings.ToLower(inp)
			re := regexp.MustCompile(`^(?:what(?:\'?s|\s+is)?\s+my(?:\s+public|\s+external)?\s+ip(?:\s+address)?)\??$`)
			return re.MatchString(lowerInp)
		},

		func(inp string) {
			execCmdAndPrintResult("dig +short myip.opendns.com @resolver1.opendns.com")
		},

		"Know your public IP address.",

		[]string{
			"what is my public ip",
			"what's my external IP address",
			"what's my IP?",
		},
	}

	specList = append(specList, s)
}
