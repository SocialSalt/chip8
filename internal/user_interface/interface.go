package userinterface

import (
	"fmt"
	"unsafe"

	"github.com/pkg/errors"
	"github.com/socialsalt/chip8/internal/processor"
	"github.com/veandco/go-sdl2/sdl"
)

var DISPLAY_SCALING int32 = 1
var VIDEO_PITCH = int(4 * processor.DISPLAY_WIDTH * DISPLAY_SCALING)

func CreateSDL() (*sdl.Window, *sdl.Renderer, *sdl.Texture, error) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return nil, nil, nil, errors.Wrap(err, "failed to init sdl")
	}
	window, err := sdl.CreateWindow(
		"CHIP8",
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		DISPLAY_SCALING*processor.DISPLAY_WIDTH,
		DISPLAY_SCALING*processor.DISPLAY_HEIGHT,
		sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE,
	)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "failed to create window")
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_TARGETTEXTURE)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "failed to get renderer")
	}
	if renderer == nil {
		return nil, nil, nil, fmt.Errorf("renderer was nil")
	}

	texture, err := renderer.CreateTexture(
		sdl.PIXELFORMAT_RGBA8888,
		sdl.TEXTUREACCESS_STREAMING,
		DISPLAY_SCALING*processor.DISPLAY_WIDTH,
		DISPLAY_SCALING*processor.DISPLAY_HEIGHT,
	)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "failed to create texture")
	}
	return window, renderer, texture, nil
}

func PollInput(c *processor.CHIP8) bool {
	event := sdl.PollEvent()
	switch t := event.(type) {
	case *sdl.KeyboardEvent:
		handleKeyboardEvent(t, c)
	case *sdl.QuitEvent:
		return false
	}
	return true
}

func UpdateDisplay(buffer []uint32, renderer *sdl.Renderer, texture *sdl.Texture, pitch int) error {
	rect := sdl.Rect{
		X: 0,
		Y: 0,
		W: DISPLAY_SCALING * processor.DISPLAY_WIDTH,
		H: DISPLAY_SCALING * processor.DISPLAY_HEIGHT,
	}
	ptr := unsafe.SliceData(buffer)
	unsafeptr := unsafe.Pointer(ptr)
	if err := texture.Update(&rect, unsafeptr, pitch); err != nil {
		return errors.Wrap(err, "failed to update texture")
	}
	if err := renderer.Clear(); err != nil {
		return errors.Wrap(err, "failed to clear the renderer")
	}

	if err := renderer.Copy(texture, nil, nil); err != nil {
		return errors.Wrap(err, "failed to copy the renderer")
	}
	renderer.Present()
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
		setKeypad(1, t.State, c.Keypad[:])
	case sdl.K_2:
		setKeypad(2, t.State, c.Keypad[:])
	case sdl.K_3:
		setKeypad(3, t.State, c.Keypad[:])
	case sdl.K_4:
		setKeypad(12, t.State, c.Keypad[:])
	case sdl.K_q:
		setKeypad(4, t.State, c.Keypad[:])
	case sdl.K_w:
		setKeypad(5, t.State, c.Keypad[:])
	case sdl.K_e:
		setKeypad(6, t.State, c.Keypad[:])
	case sdl.K_r:
		setKeypad(13, t.State, c.Keypad[:])
	case sdl.K_a:
		setKeypad(7, t.State, c.Keypad[:])
	case sdl.K_s:
		setKeypad(8, t.State, c.Keypad[:])
	case sdl.K_d:
		setKeypad(9, t.State, c.Keypad[:])
	case sdl.K_f:
		setKeypad(14, t.State, c.Keypad[:])
	case sdl.K_z:
		setKeypad(10, t.State, c.Keypad[:])
	case sdl.K_x:
		setKeypad(0, t.State, c.Keypad[:])
	case sdl.K_c:
		setKeypad(11, t.State, c.Keypad[:])
	case sdl.K_v:
		setKeypad(15, t.State, c.Keypad[:])
	}

	return nil
}
