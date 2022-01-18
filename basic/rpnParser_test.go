package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRpnParserGiveANewParser(t *testing.T) {
	p := NewRpnParser([]Token{})
	assert.NotNil(t, p, "We got a parser back")
}
func TestNewRpnParserPrepareDoesNotError(t *testing.T) {
	p := NewRpnParser(tokensForStepTest)
	prog, err := p.Prepare()

	assert.NotNil(t, err, "No error")
	assert.Nil(t, prog, "We got a program back")
}

func TestNewRpnParserGetCurrentToken(t *testing.T) {
	p := NewRpnParser(tokensForStepTest)

	assert.NotNil(t, p.CurrentToken(), "We got a token")

	p.position = -1
	assert.Nil(t, p.CurrentToken(), "We didn't get a token because of the negative position")

	p.position = 100
	assert.Nil(t, p.CurrentToken(), "We didn't get a token because of the position exceeding the len")
}
func TestNewRpnParserGetNextToken(t *testing.T) {
	p := NewRpnParser(tokensForStepTest)

	assert.NotNil(t, p.NextToken(), "We got a token")
	assert.Equal(t, p.position, 1, "Position incremented")

	p.position = -2
	assert.Nil(t, p.NextToken(), "We didn't get a token because of the negative position")

	p.position = 100
	assert.Nil(t, p.NextToken(), "We didn't get a token because of the position exceeding the len")
}

func TestNewRpnParserGetPreviousToken(t *testing.T) {
	p := NewRpnParser(tokensForStepTest)

	p.position = 1
	assert.NotNil(t, p.PreviousToken(), "We got a token")
	assert.Equal(t, p.position, 0, "Position decremented")

	assert.Nil(t, p.PreviousToken(), "We didn't get a token because of the negative position")

	p.position = 100
	assert.Nil(t, p.PreviousToken(), "We didn't get a token because of the position exceeding the len")
}

func TestNewRpnParserPeekNextToken(t *testing.T) {
	p := NewRpnParser(tokensForStepTest)

	assert.NotNil(t, p.PeekNextToken(), "We got a token")

	p.position = -2
	assert.Nil(t, p.PeekNextToken(), "We didn't get a token because of the negative position")

	p.position = 100
	assert.Nil(t, p.PeekNextToken(), "We didn't get a token because of the position exceeding the len")
}

func TestNewRpnParserPeekPreviousToken(t *testing.T) {
	p := NewRpnParser(tokensForStepTest)

	p.position = 1
	assert.NotNil(t, p.PeekPreviousToken(), "We got a token")

	p.position = 0
	assert.Nil(t, p.PeekPreviousToken(), "We didn't get a token because of the negative position")

	p.position = 100
	assert.Nil(t, p.PeekPreviousToken(), "We didn't get a token because of the position exceeding the len")
}
