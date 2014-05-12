package main

import (
	"regexp"
	"strings"
)

func init() {
	var datetimeSpecs = []spec{
		{
			func(inp string) bool {
				var lowerInp = strings.ToLower(inp)
				re1 := regexp.MustCompile(`^what\s+time\sis\sit\??$`)
				re2 := regexp.MustCompile(`^what\s+is\sthe\stime\??$`)
				return re1.MatchString(lowerInp) || re2.MatchString(lowerInp)
			},

			func(inp string) {
				execCmdAndPrintResult("date +%r")
			},

			"Get the current time.",

			[]string{
				"what time is it?",
				"what is the time",
			},
		},

		{
			func(inp string) bool {
				var lowerInp = strings.ToLower(inp)
				re := regexp.MustCompile(`^what\s+is\s+((today'?s?|the)\s+)?date\??$`)
				return re.MatchString(lowerInp)
			},

			func(inp string) {
				execCmdAndPrintResult("date +%d/%m/%Y")
			},

			"Get the current date.",

			[]string{
				"what is today's date?",
				"what is the date",
			},
		},

		{
			func(inp string) bool {
				var lowerInp = strings.ToLower(inp)
				re := regexp.MustCompile(`^what\s+month\sis\sit\??$`)
				return re.MatchString(lowerInp)
			},

			func(inp string) {
				execCmdAndPrintResult("date +%B")
			},

			"Get the current month.",

			[]string{
				"what month is it?",
			},
		},

		{
			func(inp string) bool {
				var lowerInp = strings.ToLower(inp)
				re1 := regexp.MustCompile(`^what\s+day\s+(of\s+the\s+week\s+)?is\s+it\??$`)
				re2 := regexp.MustCompile(`^what\'?s?\s+today\??$`)
				return re1.MatchString(lowerInp) || re2.MatchString(lowerInp)
			},

			func(inp string) {
				execCmdAndPrintResult("date +%A")
			},

			"Get the day of the week.",

			[]string{
				"what day is it?",
				"what day of the week is it?",
				"what's today",
			},
		},
	}

	for _, spec := range datetimeSpecs {
		specList = append(specList, spec)
	}
}
