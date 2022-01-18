package basic

type ASTNode interface {
	GetType() string
	Express() string
}

type ExpressionNode struct {
	Operator string
	Right    *ASTNode
	Left     *ASTNode
}

func (e *ExpressionNode) Express() string {
	return "42"
}

func (e *ExpressionNode) GetType() string {
	return "expression"
}

type ValueNode struct {
	Value string
}

func (v *ValueNode) GetType() string {
	return "value"
}

func (v *ValueNode) Express() string {
	return v.Value
}

type VariableNode struct {
	Value string
}

func (v *VariableNode) GetType() string {
	return "variable"
}

func (v *VariableNode) Express() string {
	return v.Value
}

type FunctionNode struct {
	Name   string
	Params []ASTNode
}

func (v *FunctionNode) GetType() string {
	return "function"
}

func (v *FunctionNode) Express() string {
	return "42"
}
