package gba_go

import (
	"fmt"
	"os"
)

func CPUReset(romPath string) {
	AddressI = 0
	ProgramCounter = 0x200

	rom, err := os.ReadFile(romPath)
	if err != nil {
		fmt.Printf("Error reading rom file: %s, err: %v\n", romPath, err)
		return
	}

	copy(GameMemory[0x200:], rom)
	fmt.Printf("File loaded successfully\n")
}

func GetNextOpcode() Word {
	res := Word(0)
	res = Word(GameMemory[ProgramCounter])
	res = res << 8
	res = res | Word(GameMemory[ProgramCounter+1])
	ProgramCounter += 2
	return res
}

func DecodeOpcode(opcode Word) {
	switch opcode & 0xF000 {
	case 0x2000:
		Opcode2NNN(opcode)
	case 0x1000:
		Opcode1NNN(opcode)
	case 0x0000:
		switch opcode & 0x000F {
		case 0x0000:
			Opcode00E0(opcode)
		case 0x000E:
			Opcode00EE(opcode)
		}
	}
}

// Opcode1NNN Jumps to address NNN.
func Opcode1NNN(opcode Word) {
	ProgramCounter = opcode & 0x0FFF
}

// Opcode00E0 Clears the screen.
func Opcode00E0(opcode Word) {}

// Opcode00EE Returns from a subroutine.
func Opcode00EE(opcode Word) {}

// Opcode2NNN Call subroutine at NNN
func Opcode2NNN(opcode Word) {
	Stack = append(Stack, ProgramCounter)
	ProgramCounter = opcode & 0x0FFF
}
