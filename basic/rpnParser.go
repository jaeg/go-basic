package basic

import "fmt"

type RpnParser struct {
	position int
	tokens   []Token
}

func NewRpnParser(tokens []Token) *RpnParser {
	return &RpnParser{tokens: tokens, position: 0}
}

func (p *RpnParser) Prepare() (*Program, error) {

	program := &Program{labelTable: map[string]int{}}

	commands := []Command{}

	p.position = -1 //We immidiately increment with the NextToken call
	for p.NextToken() != nil {
		switch p.CurrentToken().Type {
		case TokenTypeWord:
			word := p.CurrentToken()
			tempToken := p.NextToken()
			if tempToken != nil && tempToken.Type == TokenTypeEquals {
				commands = append(commands, Command{Cmd: "assign", Variable: word.Value, RPNValue: p.GetExpression()})
			}

		case TokenTypeLabel:
			program.labelTable[p.CurrentToken().Value] = len(commands)

		case TokenTypeControl:
			switch p.CurrentToken().Value {
			case "print":
				expression := p.GetExpression()
				commands = append(commands, Command{Type: CmdTypeControl, Cmd: "print", RPNValue: expression})

			case "goto":
				expression := p.GetExpression()
				commands = append(commands, Command{Type: CmdTypeControl, Cmd: "goto", RPNValue: expression})
			case "if":
				expression := p.GetExpression()
				commands = append(commands, Command{Type: CmdTypeControl, Cmd: "if", RPNValue: expression})
			case "endif":
				commands = append(commands, Command{Type: CmdTypeControl, Cmd: "endif"})
			case "else":
				commands = append(commands, Command{Type: CmdTypeControl, Cmd: "else"})
			}
		case TokenTypeNewLine:

		default:
			return nil, fmt.Errorf("error, cannot compile: %d:%s", p.position, p.CurrentToken().Value)
		}

	}

	program.commands = commands
	return program, nil
}

func (p *RpnParser) GetExpression() []Token {
	output := make([]Token, 0)
	operators := make([]Token, 0)
	for p.NextToken() != nil {

		//Hit the end of the line, finalize.
		if p.CurrentToken().Type == TokenTypeNewLine {
			for i := range operators {
				output = append(output, operators[i])
			}
			break
		} else {
			output = append(output, *p.CurrentToken())
		}
	}
	return output
}

func (p *RpnParser) CurrentToken() *Token {
	if p.position < len(p.tokens) && p.position >= 0 {
		return &p.tokens[p.position]
	}

	return nil
}

func (p *RpnParser) NextToken() *Token {
	p.position++

	return p.CurrentToken()
}

func (p *RpnParser) PreviousToken() *Token {
	p.position--
	return p.CurrentToken()
}

func (p *RpnParser) PeekNextToken() *Token {
	tempP := p.position + 1
	if tempP < len(p.tokens) && tempP >= 0 {
		return &p.tokens[tempP]
	}

	return nil
}

func (p *RpnParser) PeekPreviousToken() *Token {
	tempP := p.position - 1
	if tempP < len(p.tokens) && tempP >= 0 {
		return &p.tokens[tempP]
	}

	return nil
}
