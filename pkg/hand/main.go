package hand

import (
	"fmt"
	"strings"

	"github.com/matthewapeters/gojack/pkg/card"
)

const (
	MaxValue = 21
)

type Cards []*card.Card

func (crds *Cards) Append(v *card.Card) {
	*crds = append(*crds, v)
}

type PossibleValue struct {
	Value int
}

func (pv *PossibleValue) String() string {
	return fmt.Sprintf("%d", pv.Value)
}

type PossibleValues struct {
	Values []*PossibleValue
	Cards  *Cards
}
type Hand struct {
	TheScore []int
	TheCards *Cards
}

func NewHand() (h *Hand) {
	h = &Hand{
		TheCards: &Cards{},
		TheScore: []int{},
	}
	return
}
func (h *Hand) Takes(c *card.Card) {
	h.TheCards.Append(c)
	scores := []int{}
	for _, v := range h.values().Values {
		scores = append(scores, v.Value)
	}
	h.TheScore = scores
}

func (posVals *PossibleValues) evaluatePossibleValues(cardIndex int, p *PossibleValue) {
	c := (*posVals.Cards)[cardIndex]
	// If the card is face-down, the facevalue is not computed - just move on to the next card
	if c.FaceDown {
		if cardIndex < len(*posVals.Cards)-1 {
			posVals.evaluatePossibleValues(cardIndex+1, p)
		}
		return
	}

	// If the addition of this card exceeds MaxValue (21), remove the possible value from values and return
	if p.Value+card.Values[c.FaceValue][0] > MaxValue {
		l := len(posVals.Values)
		for i, v := range posVals.Values {
			if v == p {
				// Nil out the entry to allow the GC to collect the underylying array
				posVals.Values[i] = nil
				// There are several considerations for removing an element from an array
				if i == 0 {
					if l > 1 {
						// This is the first possible value, there are others
						posVals.Values = posVals.Values[1:]
						return
					}
					// This is the only possible value
					posVals.Values = []*PossibleValue{}
					return
				}
				if i == l-1 {
					//This is the last possible value, and there are others
					posVals.Values = posVals.Values[:i]
					return
				}
				posVals.Values = append(posVals.Values[:i], posVals.Values[i+1:]...)
				return
			}
		}
	}

	// Evaluate straight linked lists - for Aces, assume the One path by default
	startValue := p.Value
	p.Value += card.Values[c.FaceValue][0]

	// Determine if there are more cards to evaluate
	var nextCard *card.Card
	if cardIndex < len(*posVals.Cards)-1 {
		nextCard = (*posVals.Cards)[cardIndex+1]
	}

	branches := []*PossibleValue{p}

	// If there are possible branches (Aces, for example)
	for i := 1; i < len(card.Values[c.FaceValue]); i++ {
		if p.Value-card.Values[c.FaceValue][0]+card.Values[c.FaceValue][i] <= MaxValue {
			newPossibleValue := &PossibleValue{Value: startValue + card.Values[c.FaceValue][i]}
			branches = append(branches, newPossibleValue)
			posVals.Values = append(posVals.Values, newPossibleValue)
		}
	}
	if nextCard != nil {
		for _, pv := range branches {
			posVals.evaluatePossibleValues(cardIndex+1, pv)
		}
	}
}

func (h *Hand) values() (values *PossibleValues) {
	p := &PossibleValue{}
	values = &PossibleValues{Values: []*PossibleValue{p}, Cards: h.TheCards}
	if len(*h.TheCards) == 0 {
		return
	}
	values.evaluatePossibleValues(0, p)
	return
}

// Get the ith card from the hand as dealt, and re-link the remaining cards
func (h *Hand) GiveCard(cardIndex int) (c *card.Card) {
	c = (*h.TheCards)[cardIndex]
	newCards := &Cards{}
	for i, c := range *h.TheCards {
		if i != cardIndex {
			newCards.Append(c)
		}
	}
	h.TheCards = newCards
	h.values()
	return
}

func (h *Hand) String() (s string) {
	s = "Hand:\n"
	s0, s1, s2, s3, s4 := "", "", "", "", ""

	for _, k := range *h.TheCards {
		s0 += fmt.Sprintf("  ┌%s────┒", strings.Repeat("─", 3))
		if !k.FaceDown {
			s1 += fmt.Sprintf("  │%s    ┃", k)
			s2 += fmt.Sprintf("  │  %s  ┃", strings.Repeat(" ", 3))
			s3 += fmt.Sprintf("  │    %s┃", k)
		} else {
			s1 += fmt.Sprintf("  │ %s ┃", strings.Repeat("╬", 5))
			s2 += s1
			s3 += s1
		}
		s4 += fmt.Sprintf("  ┕━━━━%s┛", strings.Repeat("━", 3))
	}
	s0 += "\n"
	s1 += "\n"
	s2 += "\n"
	s3 += "\n"
	s4 += "\n"
	s += s0 + s1 + s2 + s2 + s3 + s4

	return
}

func (h *Hand) RevealFirstCard() {
	(*h.TheCards)[0].FaceDown = false
	scores := []int{}
	for _, v := range h.values().Values {
		scores = append(scores, v.Value)
	}
	h.TheScore = scores
}
