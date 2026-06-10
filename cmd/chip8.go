package main

import (
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/socialsalt/chip8/internal/processor"
	userinterface "github.com/socialsalt/chip8/internal/user_interface"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	// chip, err := processor.NewCHIP8("./roms/1-chip8-logo.ch8")
	var chip processor.CHIP8
	var err error
	if len(os.Args) > 1 {
		filename := os.Args[1]
		chip, err = processor.NewCHIP8(filename)
	} else {

		chip, err = processor.NewCHIP8("./roms/2-ibm-logo.ch8")
	}

	if err != nil {
		panic(err)
	}

	window, renderer, texture, err := userinterface.CreateSDL()
	defer sdl.Quit()
	defer window.Destroy()
	if err != nil {
		panic(errors.Wrap(err, "failed to create sdl context"))
	}

	var cycleDelay int64 = 12

	s := time.Now()
	var cpuFail chan bool
	// go func(chip *processor.CHIP8) {
	// 	for {
	// 		if err := chip.Step(); err != nil {
	// 			cpuFail <- true
	// 			break
	// 		}
	// 	}
	// }(&chip)

Outer:
	for {
		switch userinterface.PollInput(&chip) {
		case false:
			break Outer
		default:
			dt := time.Since(s).Milliseconds()
			if dt > cycleDelay {
				s = time.Now()
				fmt.Printf("Poll Time: %d\n", dt)
				if err := chip.Step(); err != nil {
					// cpuFail <- true
					panic("test")
				}
				select {
				case <-cpuFail:
					break Outer
				default:
				}
				userinterface.UpdateDisplay(chip.Display[:], renderer, texture, userinterface.VIDEO_PITCH)
			}
		}
	}
}
