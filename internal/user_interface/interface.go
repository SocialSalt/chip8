package userinterface

import (
	"fmt"

	"github.com/socialsalt/chip8/internal/processor"
	"github.com/veandco/go-sdl2/sdl"
)

func InitSDL(c *processor.CHIP8) error {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()
	window, err := sdl.CreateWindow("test title", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 5*int32(processor.DISPLAY_WIDTH), 5*int32(processor.DISPLAY_HEIGHT), sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	surface.FillRect(nil, 0)
	rect := sdl.Rect{0, 0, 200, 200}
	colour := sdl.Color{R: 255, G: 0, B: 255, A: 255} // purple
	pixel := sdl.MapRGBA(surface.Format, colour.R, colour.G, colour.B, colour.A)
	surface.FillRect(&rect, pixel)
	window.UpdateSurface()

	running := true

	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.KeyboardEvent:
				handleKeyboardEvent(t, c)
			case *sdl.QuitEvent:
				running = false
			}
		}
		fmt.Printf("%+v\n", c.Keypad)
	}
	return nil
}

func setKeypad(indx int, state uint8, keypad []byte) {
	if state == 1 {
		keypad[indx] = 1
	} else {
		keypad[indx] = 0
	}
}

func handleKeyboardEvent(t *sdl.KeyboardEvent, c *processor.CHIP8) error {
	switch t.Keysym.Sym {
	case sdl.K_1:
		setKeypad(0, t.State, c.Keypad[:])
	case sdl.K_2:
		setKeypad(1, t.State, c.Keypad[:])
	case sdl.K_3:
		setKeypad(2, t.State, c.Keypad[:])
	case sdl.K_4:
		setKeypad(3, t.State, c.Keypad[:])
	case sdl.K_q:
		setKeypad(4, t.State, c.Keypad[:])
	case sdl.K_w:
		setKeypad(5, t.State, c.Keypad[:])
	case sdl.K_e:
		setKeypad(6, t.State, c.Keypad[:])
	case sdl.K_r:
		setKeypad(7, t.State, c.Keypad[:])
	case sdl.K_a:
		setKeypad(8, t.State, c.Keypad[:])
	case sdl.K_s:
		setKeypad(9, t.State, c.Keypad[:])
	case sdl.K_d:
		setKeypad(10, t.State, c.Keypad[:])
	case sdl.K_f:
		setKeypad(11, t.State, c.Keypad[:])
	case sdl.K_z:
		setKeypad(12, t.State, c.Keypad[:])
	case sdl.K_x:
		setKeypad(13, t.State, c.Keypad[:])
	case sdl.K_c:
		setKeypad(14, t.State, c.Keypad[:])
	case sdl.K_v:
		setKeypad(15, t.State, c.Keypad[:])
	}

	return nil
}
