package game

import (
	"github.com/matthewapeters/gojack/pkg/dealer"
	"github.com/matthewapeters/gojack/pkg/player"
)

var (
	Game = &game{
		Dealer:  dealer.NewDealer(),
		Players: player.Players{player.NewPlayer("PlayerOne")},
		Scores:  map[*player.Player]int{},
	}
)

type game struct {
	*dealer.Dealer
	player.Players
	Scores map[*player.Player]int
	Quit   bool
	//HasWinner bool
	Results string
	Names   []string
	State   GameState
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
	for _, p := range Game.Players {
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
	Game.Names = n
	for Game.State != GameOver {
		// Do the next thing the game requires based on state
		GameStateMachine[Game.State]()
	}
}
