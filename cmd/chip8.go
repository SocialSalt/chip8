package main

import (
	"fmt"

	"github.com/socialsalt/chip8/internal/processor"
	userinterface "github.com/socialsalt/chip8/internal/user_interface"
)

func main() {
	fmt.Printf("hello world\n")
	userinterface.InitSDL(&processor.CHIP8{})
}
