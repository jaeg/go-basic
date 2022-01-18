package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tokensForStepTest = []Token{{Value: "1"}, {Value: "2"}, {Value: "3"}}

func TestNewParser_NewGiveANewParser(t *testing.T) {
	p := NewParser_New([]Token{})
	assert.NotNil(t, p, "We got a parser back")
}
func TestNewParser_NewPrepareDoesError(t *testing.T) {
	p := NewParser_New(tokensForStepTest)
	prog, err := p.Prepare()

	assert.NotNil(t, err, "No error")
	assert.Nil(t, prog, "We got a program back")
}

func TestNewParser_NewGetCurrentToken(t *testing.T) {
	p := NewParser_New(tokensForStepTest)

	assert.NotNil(t, p.CurrentToken(), "We got a token")

	p.position = -1
	assert.Nil(t, p.CurrentToken(), "We didn't get a token because of the negative position")

	p.position = 100
	assert.Nil(t, p.CurrentToken(), "We didn't get a token because of the position exceeding the len")
}
func TestNewParser_NewGetNextToken(t *testing.T) {
	p := NewParser_New(tokensForStepTest)

	assert.NotNil(t, p.NextToken(), "We got a token")
	assert.Equal(t, p.position, 1, "Position incremented")

	p.position = -2
	assert.Nil(t, p.NextToken(), "We didn't get a token because of the negative position")

	p.position = 100
	assert.Nil(t, p.NextToken(), "We didn't get a token because of the position exceeding the len")
}

func TestNewParser_NewGetPreviousToken(t *testing.T) {
	p := NewParser_New(tokensForStepTest)

	p.position = 1
	assert.NotNil(t, p.PreviousToken(), "We got a token")
	assert.Equal(t, p.position, 0, "Position decremented")

	assert.Nil(t, p.PreviousToken(), "We didn't get a token because of the negative position")

	p.position = 100
	assert.Nil(t, p.PreviousToken(), "We didn't get a token because of the position exceeding the len")
}

func TestNewParser_NewPeekNextToken(t *testing.T) {
	p := NewParser_New(tokensForStepTest)

	assert.NotNil(t, p.PeekNextToken(), "We got a token")

	p.position = -2
	assert.Nil(t, p.PeekNextToken(), "We didn't get a token because of the negative position")

	p.position = 100
	assert.Nil(t, p.PeekNextToken(), "We didn't get a token because of the position exceeding the len")
}

func TestNewParser_NewPeekPreviousToken(t *testing.T) {
	p := NewParser_New(tokensForStepTest)

	p.position = 1
	assert.NotNil(t, p.PeekPreviousToken(), "We got a token")

	p.position = 0
	assert.Nil(t, p.PeekPreviousToken(), "We didn't get a token because of the negative position")

	p.position = 100
	assert.Nil(t, p.PeekPreviousToken(), "We didn't get a token because of the position exceeding the len")
}
