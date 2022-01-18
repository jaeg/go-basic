package basic

import (
	"fmt"

	"github.com/antonmedv/expr"
)

type Interpreter struct {
	parser      ParserInterface
	tokens      []Token
	Environment map[string]interface{}
	Program     *Program
	variables   map[string]string
}

func NewInterpreter(program []string) (*Interpreter, error) {
	i := &Interpreter{variables: make(map[string]string)}
	i.Environment = make(map[string]interface{})

	var err error
	i.tokens, err = Tokenize(program)
	if err != nil {
		return nil, err
	}
	fmt.Println("Tokens", i.tokens)
	i.parser = NewRpnParser(i.tokens)

	i.Program, err = i.parser.Prepare()
	if err != nil {
		return nil, err
	}

	return i, nil
}

func (interpreter *Interpreter) GetParser() ParserInterface {
	return interpreter.parser
}

func (interpreter *Interpreter) GetTokens() []Token {
	return interpreter.tokens
}

func (interpreter *Interpreter) Express(input string) interface{} {
	program, err := expr.Compile(input, expr.Env(interpreter.Environment))
	if err != nil {
		panic(err)
	}

	output, err := expr.Run(program, interpreter.Environment)
	if err != nil {
		panic(err)
	}

	return output
}

func (interpreter *Interpreter) Solve(expression Expression) interface{} {
	leftExp := expression.Left
	rightExp := expression.Right
	left := ""
	right := ""

	if leftExp == nil && rightExp == nil {
		//if this.functionTable[expression.name] != undefined {
		//	return this.functionTable[expression.name](expression.params)
		//}
		fmt.Println("Function call happened")
		return "function call happened"
	} else {
		if leftExp.Type == "function" {
			//left = interpreter.functionTable[leftExp.name](leftExp.params)
			left = "0"
		} else if leftExp.Type == "op" {
			fmt.Println("Funciton call left")
			left = interpreter.Solve(*leftExp).(string)
		} else if leftExp.Type == "variable" {
			left = interpreter.variables[leftExp.Value]
		} else {
			left = leftExp.Value
		}

		if rightExp.Type == "function" {
			//right = interpreter.functionTable[rightExp.name](rightExp.params)
			fmt.Println("function call right")
			right = "0"
		} else if rightExp.Type == "op" {
			right = interpreter.Solve(*rightExp).(string)
		} else if rightExp.Type == "variable" {
			right = interpreter.variables[rightExp.Value]
		} else {
			right = rightExp.Value
		}
		fmt.Println(left, expression.Operator, right)
		return interpreter.Express(left + expression.Operator + right)
	}

	return nil
}
