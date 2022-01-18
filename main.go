package main

import (
	"fmt"

	basic "github.com/jaeg/go-basic/basic"
)

var program = []string{
	"if test < 3.3 and test > 0",
	"print \"Hello world\"",
	"value = 4",
	"endif",
}

func main() {

	interpreter, err := basic.NewInterpreter(program)
	if err != nil {
		fmt.Println(err)
	} else {
		for k, v := range interpreter.Program.GetCommands() {
			fmt.Println(k, ":", v)
			node := v.Ast

			if node != nil {
				exp := node.(*basic.ExpressionNode)
				fmt.Println(exp)
			}
		}

		//fmt.Println("Solved", interpreter.Solve(*interpreter.Program.GetCommands()[0].Value.Left))
	}

	/*
		interpreter, err := basic.NewInterpreter(program)
		if err != nil {
			fmt.Println(err)
		} else {
			for k, v := range interpreter.Program.GetCommands() {
				fmt.Println(k, ":", v)
			}
		}*/
	/*
		interpreter.Environment["foo"] = 1
		fmt.Println(interpreter.Express("foo > 0"))
		interpreter.Environment["foo"] = -1
		fmt.Println(interpreter.Express("foo > 0"))

		fmt.Println("Program:", interpreter.GetTokens()) */
}
