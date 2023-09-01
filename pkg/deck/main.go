package deck

import (
	"fmt"
	"math/rand"
	"time"

	card "github.com/matthewapeters/gojack/pkg/card"
	hand "github.com/matthewapeters/gojack/pkg/hand"
)

type DeckOfCards struct {
	Cards *hand.Hand
}

var (
	Deck = NewDeck()
)

func NewDeck() (deck *DeckOfCards) {
	deck = &DeckOfCards{Cards: hand.NewHand()}
	for _, s := range card.Suits {
		for _, fv := range card.Cards {
			deck.Cards.Takes(card.NewCard(s, fv))
		}
	}
	return
}

func (deck *DeckOfCards) Shuffle() {
	var pickCard int
	src := rand.NewSource(time.Now().UTC().Unix())
	r := rand.New(src)
	// swap the existing hand with a new hand
	oldHand := deck.Cards
	deck.Cards = hand.NewHand()

	// remaining cards is counting down the available indexes of a 52-card deck (0-indexed)
	for remainingCards := 51; remainingCards >= 0; remainingCards-- {
		if remainingCards > 0 {
			pickCard = r.Intn(remainingCards)
		} else {
			pickCard = 0
		}
		//deal card from random pick of old hand into the new hand
		deck.Cards.Takes(oldHand.GiveCard(pickCard))
	}
}

func (d *DeckOfCards) String() (s string) {
	s = "Deck:\n"
	for _, c := range *d.Cards.TheCards {
		s += fmt.Sprintf("    %s\n", c)
	}
	return
}
