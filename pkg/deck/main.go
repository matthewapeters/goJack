package deck

import (
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
	hand := hand.NewHand()
	for remainingCards := 52; remainingCards > 0; remainingCards-- {
		if remainingCards > 1 {
			pickCard = r.Intn(remainingCards) + 1
		} else {
			pickCard = 1
		}
		hand.Takes(deck.Cards.GiveCard(pickCard))
	}
	deck.Cards = hand

}
