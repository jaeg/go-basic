package basic

import "fmt"

type Parser_New struct {
	position int
	tokens   []Token
}

func NewParser_New(tokens []Token) *Parser_New {
	return &Parser_New{tokens: tokens, position: 0}
}

func (p *Parser_New) Prepare() (*Program, error) {

	program := &Program{labelTable: map[string]int{}}

	commands := []Command{}

	p.position = -1 //We immidiately increment with the NextToken call
	for p.NextToken() != nil {
		fmt.Println(p.position, p.CurrentToken().Value)
		switch p.CurrentToken().Type {
		case TokenTypeWord:
			word := p.CurrentToken()
			tempToken := p.NextToken()
			if tempToken != nil && tempToken.Type == TokenTypeEquals {
				commands = append(commands, Command{Cmd: "assign", Variable: word.Value, Ast: p.GetExpression()})
			}

		case TokenTypeLabel:
			program.labelTable[p.CurrentToken().Value] = len(commands)

		case TokenTypeControl:
			switch p.CurrentToken().Value {
			case "print":
				p.NextToken()
				var expression = p.GetExpression()
				commands = append(commands, Command{Type: CmdTypeControl, Cmd: "print", Ast: expression})

			case "goto":
				p.NextToken()
				var expression = p.GetExpression()
				commands = append(commands, Command{Type: CmdTypeControl, Cmd: "goto", Ast: expression})
			case "if":
				p.NextToken()
				var expression = p.GetExpression()
				commands = append(commands, Command{Type: CmdTypeControl, Cmd: "if", Ast: expression})
			case "endif":
				commands = append(commands, Command{Type: CmdTypeControl, Cmd: "endif"})
			case "else":
				commands = append(commands, Command{Type: CmdTypeControl, Cmd: "else"})
			}
		default:
			return nil, fmt.Errorf("error, cannot compile: %d:%s", p.position, p.CurrentToken().Value)

		}

	}

	program.commands = commands
	return program, nil
}

func (p *Parser_New) GetExpression() ASTNode {
	//Increment to the next token and get the expression starting there.
	expression := p.GetAtomic()
	if p.CurrentToken() != nil {
		for p.CurrentToken().Type == TokenTypeOp || p.CurrentToken().Type == TokenTypeEquals || p.CurrentToken().Value == "or" || p.CurrentToken().Value == "and" {

			token := p.CurrentToken()
			right := p.GetAtomic()

			expression = &ExpressionNode{Operator: token.Value, Right: &right, Left: &expression}

			p.NextToken()
		}
	}
	p.NextToken()
	return expression
}

func (p *Parser_New) GetAtomic() ASTNode {
	if p.CurrentToken() == nil {
		return nil
	}

	var value ASTNode
	if p.CurrentToken().Type == TokenTypeWord {
		value = &VariableNode{Value: p.CurrentToken().Value}
	} else if p.CurrentToken().Type == TokenTypeString {
		value = &ValueNode{Value: p.CurrentToken().Value}
	} else if p.CurrentToken().Type == TokenTypeNumber {
		value = &ValueNode{Value: p.CurrentToken().Value}
	} else if p.CurrentToken().Type == TokenTypeFunction {
		var functionName = p.CurrentToken().Value
		p.NextToken()
		params := []ASTNode{}
		for p.CurrentToken().Value != ")" {
			if p.CurrentToken().Value == "(" || p.CurrentToken().Value == "," {
				p.NextToken()
			} else {
				var expression = p.GetExpression()
				if expression != nil {
					params = append(params, expression)
				}
			}
		}
		value = &FunctionNode{Name: functionName, Params: params}
	} else if p.CurrentToken().Value == "(" {
		p.NextToken()
		return p.GetExpression()
	} else if p.CurrentToken().Value == "-" {
		p.NextToken()
		value = &ValueNode{Value: "-" + p.CurrentToken().Value}
	}

	p.NextToken()
	return value
}

func (p *Parser_New) CurrentToken() *Token {
	if p.position < len(p.tokens) && p.position >= 0 {
		return &p.tokens[p.position]
	}

	return nil
}

func (p *Parser_New) NextToken() *Token {
	p.position++

	return p.CurrentToken()
}

func (p *Parser_New) PreviousToken() *Token {
	p.position--
	return p.CurrentToken()
}

func (p *Parser_New) PeekNextToken() *Token {
	tempP := p.position + 1
	if tempP < len(p.tokens) && tempP >= 0 {
		return &p.tokens[tempP]
	}

	return nil
}

func (p *Parser_New) PeekPreviousToken() *Token {
	tempP := p.position - 1
	if tempP < len(p.tokens) && tempP >= 0 {
		return &p.tokens[tempP]
	}

	return nil
}
