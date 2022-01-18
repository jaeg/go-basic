package basic

type Program struct {
	labelTable map[string]int
	commands   []Command
}

func (p *Program) GetCommands() []Command {
	return p.commands
}
