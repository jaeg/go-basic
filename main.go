package main

import (
	"fmt"
	basic "go-basic/basic"
)

var program = []string{
	"if test < 3",
	"write \"Hello\"",
	"endif",
}

func main() {
	basic.Environment["foo"] = 1
	fmt.Println(basic.Express("foo < 0"))
	basic.Environment["foo"] = -1
	fmt.Println(basic.Express("foo > 0"))

	fmt.Println("Program:", basic.Tokenize(program))
}
