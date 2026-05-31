package processor_test

import (
	"testing"

	"github.com/socialsalt/chip8/internal/processor"
	"github.com/stretchr/testify/require"
)

func TestDisplayWrite(t *testing.T) {
	d := processor.Display{}

	d.Write(0, 0, 1)
	require.Equal(t, uint32(1), d[0])

	d.Write(3, 6, 55)
	require.Equal(t, uint32(55), d[102])
}
