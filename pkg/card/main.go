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
	ValuePath map[int]*Card
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

	c = &Card{Suit: s, FaceValue: v, ValuePath: valuePaths}
	return
}

/*
*
addNextCard recursively traverses the hand's FirstCard ValuePath and
adds 'new' to the leaf of each path.  For all new cards but Aces, this
introduces a node that is a linked list.  For Aces, it adds a branching
node with possible values of 1 and 11.

Example One:

[King of Hearts]

	             10
			    [Ace of Spades]
			    1             			11
			    [Two of Diamonds]      [Two of Diamonds]

PossibleValues:

	13                     23

Example Two:
[Two of Diamonds]
2
[Ace of Hearts]
1             			11
[Ace of Spades]         [Ace of Spades]
1           11          1             11
PossibleValues:
4           14          14            24
*/
func (c *Card) AddNextCard(new *Card) (err error) {
	// No Circular References!
	if c == new {
		err = fmt.Errorf("Cannot add %s to %s", new, c)
		return err
	}
	var val int
	var crd *Card
	for val, crd = range c.ValuePath {
		if crd == nil {
			//fmt.Printf("%s next card (%d) is %s   %+v\n", c, val, new, c.ValuePath)
			c.ValuePath[val] = new
		}
	}
	if c.ValuePath[val] != new {
		c.ValuePath[val].AddNextCard(new)
	}
	return nil
}

func (c *Card) String() string {
	return fmt.Sprintf("%s%s", c.Suit.SuitSymbol, c.FaceValue.Name)
}

func (c *Card) BaseValue() (v int) {
	v = Values[c.FaceValue][0]
	return
}
func (c *Card) NextCard() (nxt *Card) {
	nxt = c.ValuePath[c.BaseValue()]
	return
}
