package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/matthewapeters/gojack/pkg/dealer"
	"github.com/matthewapeters/gojack/pkg/player"
)

var (
	theGame = &game{
		Dealer:  dealer.NewDealer(),
		Players: player.Players{},
		Scores:  map[*player.Player]int{},
	}
)

type game struct {
	*dealer.Dealer
	player.Players
	Scores          map[*player.Player]int
	Results         string
	Names           []string
	State           GameState
	CurrentPlayerID int
}

func (g *game) AllStay() bool {
	allStay := false

	if g.Dealer.Player.Choice == player.STAY {
		allStay = true
		for _, p := range g.NotBustedPlayers() {
			allStay = allStay && p.Choice == player.STAY
		}
	}
	return allStay
}

func (g *game) AllPlayersBusted() (allBusted bool) {
	allBusted = true
	for _, p := range theGame.Players {
		allBusted = allBusted && p.GoesBust()
	}
	return
}

func (g *game) NotBustedPlayers() (players []*player.Player) {
	players = []*player.Player{}
	for _, p := range g.Players {
		if !p.GoesBust() {
			players = append(players, p)
		}
	}
	return
}

func Play(n ...string) {
	theGame.Names = n
	for theGame.State != GameOver {
		// Do the next thing the game requires based on state
		GameStateMachine[theGame.State]()
	}
}

func RequestPlayerNames() (names []string) {
	names = []string{}
	stdinReader := bufio.NewScanner(os.Stdin)
	fmt.Print("Who Wants to Play? (space-separated list): ")
	stdinReader.Scan()
	input := stdinReader.Text()
	for _, n := range strings.Split(input, " ") {
		names = append(names, n)
	}
	return
}
