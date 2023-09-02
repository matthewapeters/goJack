package deck

import (
	"fmt"
	"math/rand"
	"time"

	card "github.com/matthewapeters/gojack/pkg/card"
	hand "github.com/matthewapeters/gojack/pkg/hand"
)

type DeckOfCards struct {
	Cards         *hand.Hand
	GameCount     int
	ReturnedCards hand.Cards
}

var (
	Deck = NewDeck()
)

func NewDeck() (deck *DeckOfCards) {
	deck = &DeckOfCards{
		Cards:         hand.NewHand(),
		ReturnedCards: hand.Cards{},
	}
	for _, s := range card.Suits {
		for _, fv := range card.Cards {
			deck.Cards.Takes(card.NewCard(s, fv))
		}
	}
	return
}

func (deck *DeckOfCards) Shuffle() {
	var pickCard int
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
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

func (deck *DeckOfCards) NewGame() {
	if deck.GameCount%3 == 0 {
		deck.gatherAllCards()
		fmt.Println("Dealer Is Shuffling Deck")
		deck.Shuffle()
		time.Sleep(1 * time.Second)
		deck.GameCount = 1
	}
	deck.GameCount += 1
}

func (deck *DeckOfCards) gatherAllCards() {
	//fmt.Printf("PRE: Number of Cards in Deck: %d  Number of Cards Returned: %d\n", len(*deck.Cards.TheCards), len(deck.ReturnedCards))
	deck.Cards.TheCards.Append(deck.ReturnedCards...)
	deck.ReturnedCards = hand.Cards{}
	//fmt.Printf("POST: Number of Cards in Deck: %d  Number of Cards Returned: %d\n", len(*deck.Cards.TheCards), len(deck.ReturnedCards))
}

func (deck *DeckOfCards) PlayedCards(c hand.Cards) {
	deck.ReturnedCards.Append(c...)
}

func (d *DeckOfCards) String() (s string) {
	s = "Deck:\n"
	for _, c := range *d.Cards.TheCards {
		s += fmt.Sprintf("    %s\n", c)
	}
	return
}
