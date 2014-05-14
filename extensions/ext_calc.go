package extensions

import (
	"fmt"
	"github.com/soniah/evaler"
	"strings"
)

func init() {
	var s = spec{
		func(inp string) bool {
			lowerInp := strings.ToLower(inp)
			return strings.HasPrefix(lowerInp, "calc")
		},

		func(inp string) {
			expression := inp[len("calc "):len(inp)]
			calc(expression)
		},

		"Evaluate the mathematical expression.",

		[]string{
			"calc 3+4",
			"calc 2**2",
		},
	}
	specList = append(specList, s)
}

// calc function takes an arithmetic expression in string-form, and using
// 'evaler.Eval', it evaluates that expression and displays the result of
// it to the user in the interactive-mode.
func calc(expression string) {
	result, err := evaler.Eval(expression)
	if err != nil {
		fmt.Println("Error while evaluating expression.", err)
		return
	}
	fmt.Printf("Result of expression: %f\n", evaler.BigratToFloat(result))
}
