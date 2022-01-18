package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCommands(t *testing.T) {
	p := &Program{commands: []Command{{Type: CmdTypeControl}}}

	assert.Equal(t, 1, len(p.GetCommands()), "should be equal")
}
