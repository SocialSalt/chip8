package main

import (
	"time"

	"github.com/socialsalt/chip8/internal/processor"
	userinterface "github.com/socialsalt/chip8/internal/user_interface"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	// chip, err := processor.NewCHIP8("./roms/1-chip8-logo.ch8")
	chip, err := processor.NewCHIP8("./roms/2-ibm-logo.ch8")
	if err != nil {
		panic(err)
	}

	window, renderer, texture, err := userinterface.CreateSDL()
	defer sdl.Quit()
	defer window.Destroy()
	if err != nil {
		panic(err)
	}
	i := 0
Outer:
	for {
		switch userinterface.PollInput(&chip) {
		case false:
			break Outer
		default:
			if err := chip.Step(); err != nil {
				panic(err)
			}
			if i == 0 {
				time.Sleep(15 * time.Second)
				i++
			}
			userinterface.UpdateDisplay(chip.Display[:], renderer, texture, userinterface.VIDEO_PITCH)
		}
	}

}
