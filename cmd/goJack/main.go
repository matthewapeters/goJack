package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/matthewapeters/gojack/pkg/game"
)

func main() {
	names := []string{}
	stdinReader := bufio.NewScanner(os.Stdin)
	game.ClearScreen()
	game.TitleArt()
	fmt.Print("List Players: ")
	stdinReader.Scan()
	input := stdinReader.Text()
	for _, n := range strings.Split(input, " ") {
		names = append(names, n)
	}
	game.ClearScreen()
	game.Play(names...)
}
