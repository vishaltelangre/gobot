package extensions

import (
	"regexp"
	"strings"
)

func init() {
	var wikiSearchRegex = `^(?:define(?:\:)?\s+(?P<keyword>[\w\s]+))|(?:what\s+(?:is|does\smean)?\s+(?P<keyword>[\w\s]+))\??|(?:wiki\s+(?P<keyword>[\w\s]+))$`

	var s = spec{
		func(inp string) bool {
			var lowerInp = strings.ToLower(inp)
			return regexp.MustCompile(wikiSearchRegex).MatchString(lowerInp)
		},

		func(inp string) {
			var lowerInp = strings.ToLower(inp)
			re := regexp.MustCompile(wikiSearchRegex)
			matches := re.FindAllStringSubmatch(lowerInp, -1)[0]

			keyword := ""

			for i, _ := range matches {
				if matches[i] != "" && i != 0 {
					keyword = matches[i]
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
	}

	specList = append(specList, s)
}
