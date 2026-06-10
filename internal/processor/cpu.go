package processor

import (
	"fmt"
	"math/rand/v2"
	"os"

	"github.com/pkg/errors"
)

type CHIP8 struct {
	Registers  [16]byte
	Memory     [4096]byte
	Index      uint16
	Stack      [16]uint16
	SP         byte
	PC         uint16
	RandByte   byte
	DelayTimer byte
	SoundTimer byte
	Keypad     [16]byte
	Display    Display
	Opcode     uint16
	DrawFlag   byte
}

var ROM_START_ADDRESS = 0x200

func NewCHIP8(romFile string) (CHIP8, error) {

	// load file
	fileInfo, err := os.Stat(romFile)
	if errors.Is(err, os.ErrNotExist) {
		return CHIP8{}, err
	}

	if fileInfo.Size() > (0xFFF - 0x200) {
		return CHIP8{}, fmt.Errorf("file contents too long %v, must be less than 3583 bytes", fileInfo.Size())
	}

	fp, err := os.Open(romFile)
	if err != nil {
		return CHIP8{}, err
	}

	data := make([]byte, fileInfo.Size())
	romSize, err := fp.Read(data)
	if err != nil {
		return CHIP8{}, err
	}

	// put rom into memory
	memory := [4096]byte{}
	for i := range romSize {
		memory[ROM_START_ADDRESS+i] = data[i]
	}

	// load font into memory
	for i := range fontset {
		memory[FONT_START_ADDRESS+i] = fontset[i]
	}

	randomByte := byte(rand.Uint32())

	return CHIP8{
		Memory:   memory,
		PC:       uint16(ROM_START_ADDRESS),
		RandByte: randomByte,
	}, nil
}

func (c *CHIP8) StackPush(v uint16) error {
	if c.SP > 16 {
		return fmt.Errorf("stack overflow")
	}

	c.Stack[c.SP] = v
	c.SP++
	return nil
}

func (c *CHIP8) Step() error {
	// get the opcode out of memory using pc
	opcode_1 := c.Memory[c.PC]
	opcode_2 := c.Memory[c.PC+1]
	c.Opcode = (uint16(opcode_1) << 8) | uint16(opcode_2)
	firstHex := c.Opcode >> 12
	c.PC += 2

	// look up and execute the opcode
	err := instructionPointers[firstHex](c)
	if err != nil {
		return errors.Wrapf(err, "encountered error while executing proc: %#X", c.Opcode)
	}

	if c.DelayTimer > 0 {
		c.DelayTimer -= 0
	}

	if c.SoundTimer > 0 {
		c.SoundTimer -= 0
	}

	return nil
}
