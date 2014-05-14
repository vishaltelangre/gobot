package extensions

// spec stores functions for condition and logic to invoke if that condition is
// true.
type spec struct {
	matchCond   func(string) bool
	invoker     func(string)
	explanation string
	commands    []string
}

// specList is a slice of the 'spec' struts to store multiple extension-specific
// list of condition and logic functions.
var specList = []spec{}

func EvalQuery(q string) {
	for _, spec := range specList {
		if spec.matchCond(q) {
			spec.invoker(q)
			break
		}
	}
}
