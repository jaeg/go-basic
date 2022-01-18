package basic

type CommandType int

const (
	CmdTypeOP CommandType = iota
	CmdTypeControl
)

type ValueType int

const (
	ValueTypeVariable ValueType = iota
	ValueTypeString
	ValueTypeNumber
	ValueTypeFunction
)

type CommandInterface interface {
	GetCommandType() CommandType
	Execute() error
}

type Command struct {
	Cmd      string
	Type     CommandType
	Variable string
	Value    *Expression
	Ast      ASTNode
	RPNValue []Token
}

type Expression struct {
	Type     string
	Operator string
	Right    *Expression
	Left     *Expression
	Value    string
	Name     string // Function name being called
	Params   []*Expression
}
