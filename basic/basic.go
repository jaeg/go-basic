package basic

import (
	"github.com/antonmedv/expr"
)

var Environment = map[string]interface{}{}

func Express(input string) interface{} {
	program, err := expr.Compile(input, expr.Env(Environment))
	if err != nil {
		panic(err)
	}

	output, err := expr.Run(program, Environment)
	if err != nil {
		panic(err)
	}

	return output
}
