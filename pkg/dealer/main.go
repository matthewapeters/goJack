package dealer

import (
	"github.com/matthewapeters/gojack/pkg/deck"
	"github.com/matthewapeters/gojack/pkg/hand"
	"github.com/matthewapeters/gojack/pkg/player"
)

type Dealer struct {
	Deck   *deck.DeckOfCards
	Player *player.Player
}

func NewDealer() *Dealer {
	return &Dealer{
		Deck:   deck.NewDeck(),
		Player: player.NewPlayer("Dealer"),
	}
}

func (d *Dealer) DealsACard(p *player.Player) {
	p.Hand.Takes(d.Deck.Cards.GiveCard(1))
}

func (d *Dealer) Shuffles() {
	d.Deck.Shuffle()
}

func (d *Dealer) NewGame() {
	d.Deck = deck.NewDeck()
	d.Player.Hand = hand.NewHand()
}
