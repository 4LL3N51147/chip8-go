package gba_go

type Word uint16

var (
	GameMemory     = make([]byte, 0xFFF)
	Registers      = make([]byte, 16)
	AddressI       Word
	ProgramCounter Word
	ScreenData     [][]byte
	Stack          = make([]Word, 0)
)
