package basic

import (
	"fmt"
)

type Parser struct {
	position int
	tokens   []Token
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Prepare(tokens []Token) (*Program, error) {
	p.tokens = tokens
	program := &Program{labelTable: map[string]int{}}

	commands := []Command{}
	for p.position < len(tokens)-1 {
		if tokens[p.position].Type == TokenTypeWord {
			word := tokens[p.position]
			p.position++
			if tokens[p.position].Type == TokenTypeEquals {
				p.position++
				commands = append(commands, Command{Cmd: "assign", Variable: word.Value, Value: p.getExpression()})
			}
		} else if tokens[p.position].Type == TokenTypeLabel {
			program.labelTable[tokens[p.position].Value] = len(commands)
			p.position++
		} else if tokens[p.position].Type == TokenTypeControl {
			var token = tokens[p.position]
			switch token.Value {
			case "print":
				p.position++
				var expression = p.getExpression()
				commands = append(commands, Command{Type: CmdTypeControl, Cmd: "print", Value: expression})

			case "goto":
				p.position++
				var expression = p.getExpression()
				commands = append(commands, Command{Type: CmdTypeControl, Cmd: "goto", Value: expression})
				p.position++

			case "if":
				p.position++
				var expression = p.getExpression()
				commands = append(commands, Command{Type: CmdTypeControl, Cmd: "if", Value: expression})

			case "endif":
				commands = append(commands, Command{Type: CmdTypeControl, Cmd: "endif"})
				p.position++
			case "else":
				commands = append(commands, Command{Type: CmdTypeControl, Cmd: "else"})
				p.position++
			default:
				p.position++
			}
		} else {
			return nil, fmt.Errorf("error, cannot compile: %d:%s", p.position, tokens[p.position].Value)
		}
	}

	program.commands = commands
	return program, nil
}

func (p *Parser) getExpression() *Expression {
	expression := p.getAtomic()
	for p.getCurrentToken().Type == TokenTypeOp || p.getCurrentToken().Type == TokenTypeEquals || p.getCurrentToken().Value == "or" || p.getCurrentToken().Value == "and" {
		var token = p.getCurrentToken()
		p.position++
		var right = p.getAtomic()
		expression = &Expression{Type: "op", Operator: token.Value, Right: right, Left: expression}
	}

	return expression
}

func (p *Parser) getCurrentToken() Token {
	return p.tokens[p.position]
}

func (p *Parser) getAtomic() *Expression {
	var value *Expression
	if p.getCurrentToken().Type == TokenTypeWord {
		value = &Expression{Type: "variable", Value: p.getCurrentToken().Value}
	} else if p.getCurrentToken().Type == TokenTypeString {
		value = &Expression{Type: "string", Value: p.getCurrentToken().Value}
	} else if p.getCurrentToken().Type == TokenTypeNumber {
		value = &Expression{Type: "number", Value: p.getCurrentToken().Value}
	} else if p.getCurrentToken().Type == TokenTypeFunction {
		var functionName = p.getCurrentToken().Value
		p.position++
		params := []*Expression{}
		for p.getCurrentToken().Value != ")" {
			if p.getCurrentToken().Value == "(" || p.getCurrentToken().Value == "," {
				p.position++
			} else {
				var expression = p.getExpression()
				if expression.Type != "" {
					params = append(params, expression)
				}
			}
		}
		value = &Expression{Type: "function", Name: functionName, Params: params}
	} else if p.getCurrentToken().Value == "(" {
		p.position++
		value = p.getExpression()
	} else if p.getCurrentToken().Value == "-" {
		p.position++
		value = &Expression{Type: "number", Value: "-" + p.getCurrentToken().Value}
	}

	p.position++
	return value
}
