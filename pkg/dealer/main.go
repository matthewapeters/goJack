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

func (d *Dealer) NewGame() {
	// New Game returns played cards to deck and reshuffles the deck every third game
	d.GatherPlayedCards(*d.Player.Hand.TheCards)
	d.Deck.NewGame()
	d.Player.NewGame()
}

func (d *Dealer) RevealFirstCard() {
	//Reveal the dealer's first card
	d.Player.Hand.RevealFirstCard()
}

func (d *Dealer) GoesBust() (busted bool) {
	busted = ((*d.Deck.Cards.TheCards)[0].FaceDown && d.Player.Scores()[player.MIN] >= 21) ||
		d.Player.GoesBust()
	return
}

func (d *Dealer) GatherPlayedCards(crds hand.Cards) {
	d.Deck.PlayedCards(crds)
}
