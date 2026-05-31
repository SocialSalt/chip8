package processor_test

import (
	"testing"

	"github.com/socialsalt/chip8/internal/processor"
	"github.com/stretchr/testify/require"
)

func TestBCB(t *testing.T) {
	c := processor.CHIP8{
		Memory: [4096]byte{},
		PC:     uint16(processor.ROM_START_ADDRESS),
		Index:  0,
	}

	c.Registers[0] = 10
	c.Opcode = 0xF033

	require.NoError(t, processor.BCD(&c))

	require.Equal(t, byte(0), c.Memory[0])
	require.Equal(t, byte(1), c.Memory[1])
	require.Equal(t, byte(0), c.Memory[2])

	c.Registers[0xD] = 187
	c.Opcode = 0xFD33

	require.NoError(t, processor.BCD(&c))

	require.Equal(t, byte(1), c.Memory[0])
	require.Equal(t, byte(8), c.Memory[1])
	require.Equal(t, byte(7), c.Memory[2])
}
