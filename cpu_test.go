package gba_go

import "testing"
import "github.com/stretchr/testify/assert"

func TestGetNextOpcode(t *testing.T) {
	GameMemory[0] = 0xAB
	GameMemory[1] = 0xCD
	ProgramCounter = 0
	assert.Equal(t, Word(0xABCD), GetNextOpcode())
	assert.Equal(t, Word(2), ProgramCounter)
}
