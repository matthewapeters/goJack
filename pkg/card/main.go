// Package card implements playing card objects (card.Card).  Cards have face-values and suits, which
// are both sortable.  A card may be face-down or not.  Face-Values may have more than one value - for example,
// Aces may count as 1 or 11 as best serves the player.
package card

import (
	"fmt"
)

type SuitName string
type SuitSymbol string
type FaceValue struct {
	Name  string
	Order int
}

type FaceValues []FaceValue

func (fv FaceValues) Less(i, j int) bool {
	return fv[i].Order < fv[j].Order
}

func (fv FaceValues) Len() int {
	return len(fv)
}

func (fv FaceValues) Swap(i, j int) {
	t := fv[i]
	fv[i] = fv[j]
	fv[j] = t
}

type Suit struct {
	SuitName
	SuitSymbol
	Order int
}

// Card provides a unique instance of a playing card.  Both the Suit and Facealue
// provide sorting capabilities.  FaceDown informs how the card is rendered
// as part of a hand, as well as the hand value.
type Card struct {
	Suit
	FaceValue
	FaceDown bool
}

var (
	Ace   = FaceValue{"A", 0}
	King  = FaceValue{"K", 1}
	Queen = FaceValue{"Q", 2}
	Jack  = FaceValue{"J", 3}
	Ten   = FaceValue{"10", 4}
	Nine  = FaceValue{"9", 5}
	Eight = FaceValue{"8", 6}
	Seven = FaceValue{"7", 7}
	Six   = FaceValue{"6", 8}
	Five  = FaceValue{"5", 9}
	Four  = FaceValue{"4", 10}
	Three = FaceValue{"3", 11}
	Two   = FaceValue{"2", 12}

	Spades   = Suit{"Spades", "♠", 0}
	Hearts   = Suit{"Hearts", "♥", 1}
	Diamonds = Suit{"Diamonds", "♦", 2}
	Clubs    = Suit{"Clubs", "♣", 3}
	Suits    = []Suit{Spades, Hearts, Diamonds, Clubs}
	Cards    = []FaceValue{Ace, King, Queen, Jack, Ten, Nine, Eight, Seven, Six, Five, Four, Three, Two}
	Values   = map[FaceValue][]int{Ace: {1, 11}, King: {10}, Queen: {10}, Jack: {10}, Ten: {10}, Nine: {9},
		Eight: {8}, Seven: {7}, Six: {6}, Five: {5}, Four: {4}, Three: {3}, Two: {2}}
)

// NewCard returns a pointer to a unique card instance.
func NewCard(s Suit, v FaceValue) (c *Card) {
	c = &Card{Suit: s, FaceValue: v}
	return
}

// Returns a string consisting of the Card's Suit Symbol and Face Value.  String is
// three characters long.
func (c *Card) String() string {
	return fmt.Sprintf("%s%s", c.Suit.SuitSymbol, c.FaceValue)
}

// Returns the first value associated with a card, based on its FaceValue.
// For example, will return 1 for Ace
func (c *Card) BaseValue() (v int) {
	v = Values[c.FaceValue][0]
	return
}

// Returns the suit's symbol
func (s Suit) String() string {
	return string(s.SuitSymbol)
}

// Returns a 2-character (left-padded) representation of the face value
func (fv FaceValue) String() string {
	return fmt.Sprintf("%2s", fv.Name)
}
