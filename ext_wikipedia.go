package main

import (
	"fmt"
	"regexp"
	"strings"
)

func init() {
	var wikiSpecs = []spec{
		{
			func(inp string) bool {
				var lowerInp = strings.ToLower(inp)
				re := regexp.MustCompile(`^(?:define(?:\:)?\s+(?P<keyword>[\w\s]+))|(?:what\s+(?:is|does\smean)?\s+(?P<keyword>[\w\s]+))\??|(?:wiki\s+(?P<keyword>[\w\s]+))$`)
				return re.MatchString(lowerInp)
			},

			func(inp string) {
				var lowerInp = strings.ToLower(inp)
				re := regexp.MustCompile(`^(?:define(?:\:)?\s+(?P<keyword>[\w\s]+))|(?:what\s+(?:is|does\smean)?\s+(?P<keyword>[\w\s]+))\??|(?:wiki\s+(?P<keyword>[\w\s]+))$`)
				r2 := re.FindAllStringSubmatch(lowerInp, -1)[0]

				keyword := ""

				for i, _ := range r2 {
					if r2[i] != "" && i != 0 {
						keyword = r2[i]
						break
					}
				}

				command := []string{"dig +short txt ", keyword, ".wp.dg.cx"}
				execCmdAndPrintResult(strings.Join(command, ""))
			},

			"Get short info from Wikipedia about a keyword.",

			[]string{
				"define: <keyword>",
				"what is <keyword>?",
				"what does mean <keyword>?",
				"wiki <keyword>",
			},
		},
	}

	for _, spec := range wikiSpecs {
		specList = append(specList, spec)
	}
}
