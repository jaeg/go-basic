package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenizingSuccessfulProgram(t *testing.T) {
	program := []string{
		"if test < 3.3 and test > 0",
		"print \"Hello world\"",
		"endif",
	}

	tokens, err := Tokenize(program)
	assert.Nil(t, err, "No error tokenizing")

	assert.Equal(t, "if", tokens[0].Value, "should be the 'if' token")
}

func TestTokenizingFailsWithMisplacedWord(t *testing.T) {
	program := []string{
		"if test < 3.3 and test > 0",
		"print \"Hello world\"",
		"endif asdf",
	}

	_, err := Tokenize(program)
	assert.Error(t, err, "error while tokenizing")
}

func TestTokenizingFailsWithBadFloat(t *testing.T) {
	program := []string{
		"if test < 3.f and test > 0",
		"print \"Hello world\"",
		"endif",
	}

	_, err := Tokenize(program)
	assert.Error(t, err, "error while tokenizing")
}
