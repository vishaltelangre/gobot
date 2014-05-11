package main

import (
	"regexp"
	"strings"
)

func init() {
	var s = spec{

		func(inp string) bool {
			var lowerInp = strings.ToLower(inp)
			re1 := regexp.MustCompile(`^what\s+time\sis\sit\??$`)
			re2 := regexp.MustCompile(`^what\s+is\sthe\stime\??$`)
			return re1.MatchString(lowerInp) || re2.MatchString(lowerInp)
		},

		func(inp string) {
			execCmdAndPrintResult("date +\"%r\"")
		},

		"Gets the current time.",

		[]string{
			"what time is it?",
			"what is the time",
		},
	}
	specList = append(specList, s)
}
