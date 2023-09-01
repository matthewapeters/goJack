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

type Card struct {
	Suit
	FaceValue
	FaceDown bool
}

var (
	Ace   = FaceValue{" A", 0}
	King  = FaceValue{" K", 1}
	Queen = FaceValue{" Q", 2}
	Jack  = FaceValue{" J", 3}
	Ten   = FaceValue{"10", 4}
	Nine  = FaceValue{" 9", 5}
	Eight = FaceValue{" 8", 6}
	Seven = FaceValue{" 7", 7}
	Six   = FaceValue{" 6", 8}
	Five  = FaceValue{" 5", 9}
	Four  = FaceValue{" 4", 10}
	Three = FaceValue{" 3", 11}
	Two   = FaceValue{" 2", 12}

	Spades   = Suit{"Spades", "♠", 0}
	Hearts   = Suit{"Hearts", "♥", 1}
	Diamonds = Suit{"Diamonds", "♦", 2}
	Clubs    = Suit{"Clubs", "♣", 3}
	Suits    = []Suit{Spades, Hearts, Diamonds, Clubs}
	Cards    = []FaceValue{Ace, King, Queen, Jack, Ten, Nine, Eight, Seven, Six, Five, Four, Three, Two}
	Values   = map[FaceValue][]int{Ace: {1, 11}, King: {10}, Queen: {10}, Jack: {10}, Ten: {10}, Nine: {9},
		Eight: {8}, Seven: {7}, Six: {6}, Five: {5}, Four: {4}, Three: {3}, Two: {2}}
)

func FaceValueSort(fv1, fv2 FaceValue) bool {
	return fv1.Order < fv2.Order
}

func NewCard(s Suit, v FaceValue) (c *Card) {
	valuePaths := map[int]*Card{}
	valuePaths[Values[v][0]] = nil
	for i := 1; i < len(Values[v]); i++ {
		valuePaths[Values[v][i]] = nil
	}

	c = &Card{Suit: s, FaceValue: v} //ValuePath: valuePaths}
	return
}

func (c *Card) String() string {
	return fmt.Sprintf("%s%s", c.Suit.SuitSymbol, c.FaceValue.Name)
}

func (c *Card) BaseValue() (v int) {
	v = Values[c.FaceValue][0]
	return
}
