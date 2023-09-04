// goJack is an implementation of multi-player Black Jack or Twentyone
package main

import (
	"github.com/matthewapeters/gojack/pkg/game"
)

func main() {
	game.ClearScreen()
	game.TitleArt()
	game.ClearScreen()
	game.Play(game.RequestPlayerNames()...)
}
