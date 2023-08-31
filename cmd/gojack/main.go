package main

import (
	"fmt"

	"github.com/matthewapeters/gojack/pkg/game"
)

func main() {
	game.ClearScreen()
	game.TitleArt()
	n := ""
	fmt.Print("What is your name? ")
	fmt.Scanln(&n)
	game.ClearScreen()
	game.Play(n)
}
