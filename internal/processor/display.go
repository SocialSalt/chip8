package processor

import "fmt"

var DISPLAY_HEIGHT int32 = 64
var DISPLAY_WIDTH int32 = 64

type Display [64 * 32]uint32

func (d *Display) Write(x byte, y byte, v uint32) error {
	index := uint32(x)*32 + uint32(y)
	if int32(index) > DISPLAY_HEIGHT*DISPLAY_WIDTH {
		return fmt.Errorf("attempt to write to array out of bounds. x: %v, y: %v, i: %v", x, y, index)
	}
	d[index] = v
	return nil
}
