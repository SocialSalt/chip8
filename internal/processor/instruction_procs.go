package processor

import "fmt"

type InstructionProc func(*CHIP8) error

func Proc0Collector(c *CHIP8) error {
	switch c.Opcode {
	case 0x00E0:
		return CLS(c)
	case 0x00EE:
		return RET(c)
	default:
		return JP(c)
	}
}

func Proc8Collector(c *CHIP8) error {
	finalHex := c.Opcode & 0x000F
	return proc8Array[finalHex](c)
}

func ProcECollector(c *CHIP8) error {
	finalHexes := c.Opcode & 0x00FF
	switch finalHexes {
	case 0x9E:
		return SKP(c)
	case 0xA1:
		return SKNP(c)
	default:
		return fmt.Errorf("unrecognized opcode: %#X", c.Opcode)
	}
}

func ProcFCollector(c *CHIP8) error {
	finalHexes := c.Opcode & 0x00FF
	return procFArray[finalHexes](c)
}

// 00E0
func CLS(c *CHIP8) error {
	for i := range c.Display {
		c.Display[i] = 0
	}
	return nil
}

// 00EE
func RET(c *CHIP8) error {
	c.SP -= 1
	c.PC = c.Stack[c.SP]
	return nil
}

// 0NNN
func JP(c *CHIP8) error {
	c.PC = c.Opcode & 0x0FFF
	return nil
}

// 2NNN
func CALL(c *CHIP8) error {
	if err := c.StackPush(c.PC); err != nil {
		return err
	}
	c.PC = c.Opcode & 0x0FFF
	return nil
}

// 3XNN
// Skip if eq to register
func SER(c *CHIP8) error {
	registerIdx := (c.Opcode & 0x0F00) >> 8
	v := byte(c.Opcode & 0x00FF)

	if c.Registers[registerIdx] == v {
		c.PC += 2
	}
	return nil
}

// 4XNN
// Skip if NE to register
func SNER(c *CHIP8) error {
	registerIdx := (c.Opcode & 0x0F00) >> 8
	v := byte(c.Opcode & 0x00FF)

	if c.Registers[registerIdx] != v {
		c.PC += 2
	}
	return nil
}

// 5XY0
// skip if registers equal
func SERR(c *CHIP8) error {
	registerX := (c.Opcode & 0x0F00) >> 8
	registerY := (c.Opcode & 0x00F0) >> 4

	if c.Registers[registerX] == c.Registers[registerY] {
		c.PC += 2
	}
	return nil
}

// 6XNN
func LDR(c *CHIP8) error {
	registerIdx := (c.Opcode & 0x0F00) >> 8
	v := byte(c.Opcode)

	c.Registers[registerIdx] = v
	return nil
}

// 7XNN
func ADDRN(c *CHIP8) error {
	registerX := (c.Opcode & 0x0F00) >> 8
	v := byte(c.Opcode & 0x00FF)

	c.Registers[registerX] += v
	return nil
}

// 8XY0 Load register y into x
func LDRR(c *CHIP8) error {
	registerX := (c.Opcode & 0x0F00) >> 8
	registerY := (c.Opcode & 0x00F0) >> 4

	c.Registers[registerX] = c.Registers[registerY]
	return nil
}

// 8XY1 x = x | y
func ORRR(c *CHIP8) error {
	registerX := (c.Opcode & 0x0F00) >> 8
	registerY := (c.Opcode & 0x00F0) >> 4

	c.Registers[registerX] = c.Registers[registerX] | c.Registers[registerY]
	return nil
}

// 8XY2 x = x & y
func ANDRR(c *CHIP8) error {
	registerX := (c.Opcode & 0x0F00) >> 8
	registerY := (c.Opcode & 0x00F0) >> 4

	c.Registers[registerX] = c.Registers[registerX] & c.Registers[registerY]
	return nil
}

// 8XY3 x = x ^ y
func XORRR(c *CHIP8) error {
	registerX := (c.Opcode & 0x0F00) >> 8
	registerY := (c.Opcode & 0x00F0) >> 4

	c.Registers[registerX] = c.Registers[registerX] ^ c.Registers[registerY]
	return nil
}

// 8XY4 add with carry
func ADDRRC(c *CHIP8) error {
	registerX := (c.Opcode & 0x0F00) >> 8
	registerY := (c.Opcode & 0x00F0) >> 4

	sum := uint16(c.Registers[registerX]) + uint16(c.Registers[registerY])
	if sum > 0xFF {
		c.Registers[0xF] = 1
	} else {
		c.Registers[0xF] = 0
	}
	c.Registers[registerX] = byte(sum)
	return nil
}

// 8XY5 subtract with borrow
func SUBXYB(c *CHIP8) error {
	registerX := (c.Opcode & 0x0F00) >> 8
	registerY := (c.Opcode & 0x00F0) >> 4

	if c.Registers[registerX] > c.Registers[registerY] {
		c.Registers[0xF] = 1
	} else {
		c.Registers[0xF] = 0
	}
	c.Registers[registerX] = c.Registers[registerX] - c.Registers[registerY]
	return nil
}

// TODO:
// derermine if this is wrong, the blog says to only use registerX
// other documentation says use x and y
// 8XY6 x = y >> 1; save smallest bit in VF
func SHIFTR(c *CHIP8) error {
	registerX := (c.Opcode & 0x0F00) >> 8
	registerY := (c.Opcode & 0x00F0) >> 4

	smallestBit := c.Registers[registerY] & (0x01)
	c.Registers[registerX] = c.Registers[registerY] >> 1
	c.Registers[0xF] = smallestBit
	return nil
}

// 8XY7 x = y-x
func SUBYXB(c *CHIP8) error {
	registerX := (c.Opcode & 0x0F00) >> 8
	registerY := (c.Opcode & 0x00F0) >> 4

	if c.Registers[registerY] > c.Registers[registerX] {
		c.Registers[0xF] = 1
	} else {
		c.Registers[0xF] = 0
	}
	c.Registers[registerX] = c.Registers[registerY] - c.Registers[registerX]
	return nil

}

// 8XYE
func SHIFTL(c *CHIP8) error {
	registerX := (c.Opcode & 0x0F00) >> 8
	registerY := (c.Opcode & 0x00F0) >> 4

	// 0x80 = 0b1000_0000
	biggestBit := c.Registers[registerY] & (0x80)
	c.Registers[registerX] = c.Registers[registerY] << 1
	c.Registers[0xF] = biggestBit
	return nil
}

// 9XY0 skip if registers are not equal
func SNERR(c *CHIP8) error {
	registerX := (c.Opcode & 0x0F00) >> 8
	registerY := (c.Opcode & 0x00F0) >> 4

	if c.Registers[registerX] != c.Registers[registerY] {
		c.PC += 2
	}
	return nil
}

// ANNN store Memory address NNN into register I
func STOREMI(c *CHIP8) error {
	addr := (c.Opcode & 0x0FFF)
	c.Index = addr
	return nil
}

// BNNN Jump to NNN + V0
func JUMP(c *CHIP8) error {
	addr := (c.Opcode & 0x0FFF)
	c.PC = addr + uint16(c.Registers[0])

	return nil
}

// CXNN set VX to random number masked with NN
func RANDR(c *CHIP8) error {
	registerX := (c.Opcode & 0x0F00) >> 8
	mask := byte(c.Opcode & 0x00FF)

	c.Registers[registerX] = c.RandByte & mask
	return nil
}

// DXYN draw a sprite
// all sprites are 8px wide and 1-15 px tall
func DRAW(c *CHIP8) error {
	registerX := ((c.Opcode & 0x0F00) >> 8)
	registerY := ((c.Opcode & 0x00F0) >> 4)
	numBytes := (c.Opcode & 0x000F)

	x := c.Registers[registerX] % 0x3F
	y := c.Registers[registerY] % 0x1F
	c.Registers[0xF] = 0

	for row := range numBytes {
		byte := c.Memory[c.Index+row]
		startIndex := uint32(x) + uint32(y+uint8(row))*64

		// read the byte left to right
		// 0x80
		// 0b10000000
		for pixelOffset := range uint32(8) {
			val := c.Display[startIndex+pixelOffset]
			setVal := uint32(((byte) & (0x80 >> pixelOffset)))
			if setVal > 0 {
				if val == 0xFFFFFFFF {
					c.Registers[0xF] = 1
				}
				c.Display[startIndex+pixelOffset] ^= 0xFFFFFFFF
			}
		}
	}
	return nil
}

// EX9E
func SKP(c *CHIP8) error {
	registerX := ((c.Opcode & 0x0F00) >> 8)
	val := c.Registers[registerX]
	if c.Keypad[val] != 0 {
		c.PC += 2
	}
	return nil
}

// EXA1
func SKNP(c *CHIP8) error {
	registerX := ((c.Opcode & 0x0F00) >> 8)
	val := c.Registers[registerX]
	if c.Keypad[val] == 0 {
		c.PC += 2
	}
	return nil
}

// FX07
func StoreDelay(c *CHIP8) error {
	registerX := ((c.Opcode & 0x0F00) >> 8)
	c.Registers[registerX] = c.DelayTimer
	return nil
}

// FX0A
func WAIT(c *CHIP8) error {
	registerX := ((c.Opcode & 0x0F00) >> 8)
	set := false
	for i, val := range c.Keypad {
		if val != 0 {
			c.Registers[registerX] = byte(i)
			set = true
			break
		}
	}
	if !set {
		c.PC -= 2
	}
	return nil
}

// FX15
func SetDelay(c *CHIP8) error {
	registerX := ((c.Opcode & 0x0F00) >> 8)
	c.DelayTimer = c.Registers[registerX]
	return nil
}

// FX18
func SetSound(c *CHIP8) error {
	registerX := ((c.Opcode & 0x0F00) >> 8)
	c.SoundTimer = c.Registers[registerX]
	return nil
}

// FX1E
func ADDRI(c *CHIP8) error {
	registerX := ((c.Opcode & 0x0F00) >> 8)
	c.Index += uint16(c.Registers[registerX])
	return nil
}

// FX29
func LDRI(c *CHIP8) error {
	registerX := ((c.Opcode & 0x0F00) >> 8)
	c.Index = uint16(FONT_START_ADDRESS) + 5*uint16(c.Registers[registerX])
	return nil
}

// FX33
func BCD(c *CHIP8) error {
	registerX := ((c.Opcode & 0x0F00) >> 8)
	val := c.Registers[registerX]
	ones := val % 10
	val = val / 10
	tens := val % 10
	val = val / 10
	hundos := val % 10

	c.Memory[c.Index] = hundos
	c.Memory[c.Index+1] = tens
	c.Memory[c.Index+2] = ones
	return nil
}

// FX55
func WriteMRX(c *CHIP8) error {
	registerX := ((c.Opcode & 0x0F00) >> 8)
	for i := range registerX + 1 {
		c.Memory[c.Index+i] = c.Registers[i]
	}
	return nil
}

// FX65
func WriteRXM(c *CHIP8) error {
	registerX := ((c.Opcode & 0x0F00) >> 8)
	for i := range registerX + 1 {
		c.Registers[i] = c.Memory[c.Index+i]
	}
	return nil
}
