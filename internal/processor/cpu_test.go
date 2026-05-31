package processor_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/socialsalt/chip8/internal/processor"
	"github.com/stretchr/testify/require"
)

func TestCPU(t *testing.T) {
	dir, err := os.Getwd()
	fmt.Println(dir)
	require.NoError(t, err)

	chip, err := processor.NewCHIP8("../../roms/1-chip8-logo.ch8")
	require.NoError(t, err)
	for range 39 {
		chip.Step()
		// for i := range chip.Display {
		// 	if chip.Display[i] > 0 {
		// 		fmt.Printf("0")
		// 	} else {
		// 		fmt.Printf(" ")
		// 	}
		// 	if i%int(processor.DISPLAY_WIDTH) == 0 {
		// 		fmt.Printf("\n")
		// 	}
		// }
	}
}
