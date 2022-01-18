package basic

type ParserInterface interface {
	Prepare() (*Program, error)
}
