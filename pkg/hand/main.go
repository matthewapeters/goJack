package hand

import (
	"fmt"
	"strings"

	"github.com/matthewapeters/gojack/pkg/card"
)

const (
	MaxValue = 21
)

type PossibleValue struct {
	Cards []*card.Card
	Value int
}

func (pv *PossibleValue) String() string {
	return fmt.Sprintf("%d (%s)", pv.Value, pv.Cards)
}

type PossibleValues struct {
	Values []*PossibleValue
}
type Hand struct {
	//Cards     map[card.Suit]map[card.FaceValue]*card.Card
	FirstCard *card.Card
	TheScore  []int
	TheCards  []*card.Card
}

func NewHand() (h *Hand) {
	h = &Hand{
		FirstCard: nil,
		TheCards:  []*card.Card{},
		TheScore:  []int{},
	}
	return
}
func (h *Hand) Takes(c *card.Card) {
	if h.FirstCard == nil {
		h.FirstCard = c
	} else {
		h.FirstCard.AddNextCard(c)
	}
	h.TheCards = append(h.TheCards, c)
	scores := []int{}
	for _, v := range h.values().Values {
		scores = append(scores, v.Value)
	}
	h.TheScore = scores

}
func (posVals *PossibleValues) evaluatePossibleValues(c *card.Card, p *PossibleValue) {
	if c == nil {
		fmt.Printf("evaluatePossibleValues received nil for *card.Card\n")
		return
	}
	//fmt.Printf("evalutaing card %s\n", c)
	// If the addition of this card exceeds MaxValue (21), remove the possible value from values and return
	if p.Value+card.Values[c.FaceValue][0] > MaxValue {
		//fmt.Printf("next card %s exceeds %d\n", c, MaxValue)
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
	//fmt.Printf("DefaultPath: add %d to starting value %d %s\n", card.Values[c.FaceValue][0], startValue, c)
	p.Cards = append(p.Cards, c)
	nextCard := c.ValuePath[card.Values[c.FaceValue][0]]
	branches := []*PossibleValue{p}

	for i := 1; i < len(card.Values[c.FaceValue]); i++ {
		if p.Value-card.Values[c.FaceValue][0]+card.Values[c.FaceValue][i] <= MaxValue {
			//fmt.Printf("adding possible value (%d) %d to %d %s\n", i, card.Values[c.FaceValue][i], startValue, c)
			newCards := []*card.Card{}
			for _, crd := range p.Cards {
				newCards = append(newCards, crd)
			}
			newPossibleValue := &PossibleValue{Cards: newCards, Value: startValue + card.Values[c.FaceValue][i]}
			branches = append(branches, newPossibleValue)
			posVals.Values = append(posVals.Values, newPossibleValue)
		}
	}
	if nextCard != nil {
		for _, pv := range branches {
			posVals.evaluatePossibleValues(nextCard, pv)
		}
	} else {
		//println("no next card")
	}
}
func (h *Hand) values() (values *PossibleValues) {
	p := &PossibleValue{}
	values = &PossibleValues{Values: []*PossibleValue{p}}
	if h.FirstCard == nil {
		return
	}
	values.evaluatePossibleValues(h.FirstCard, p)
	return
}

// Get the ith card from the hand as dealt, and re-link the remaining cards
func (h *Hand) GiveCard(i int) (c *card.Card) {
	// Edge condition: asking for the first card
	if i == 1 {
		c = h.FirstCard
		next := c.NextCard()
		h.FirstCard = next
		for val := range c.ValuePath {
			c.ValuePath[val] = nil
		}
		return
	}
	//fmt.Printf("FirstCard: %+v  ValuePath: %+v  FaceValue: %+v\n", h.FirstCard, h.FirstCard.ValuePath, h.FirstCard.FaceValue)
	priorCard := h.FirstCard
	cardList := []*card.Card{h.FirstCard}
	//fmt.Printf("FirstCard is %s\n", h.FirstCard)
	theCard := priorCard.NextCard()
	for j := i - 1; j > 1; j-- {
		//fmt.Printf("J: %d  %s\n", j, theCard)
		priorCard = theCard
		cardList = append(cardList, priorCard)
		theCard = theCard.NextCard()
	}
	// Knitt the next card to the prior card
	nextCard := theCard.NextCard()
	for k := range priorCard.ValuePath {
		priorCard.ValuePath[k] = nextCard
	}

	// Clear the ValuePath of the card we are giving
	c = theCard
	for k := range c.ValuePath {
		c.ValuePath[k] = nil
	}
	// Remove the card from the sorted Cards
	//delete(h.Cards[c.Suit], c.FaceValue)
	//fmt.Printf("counted over %d cards %s, and gave %s\n", len(cardList), cardList, c)
	//fmt.Printf("The new next card after %s is %+v\n", cardList[len(cardList)-1], cardList[len(cardList)-1].ValuePath)
	return
}

func (h *Hand) String() (s string) {
	s = "Hand:\n"
	s0 := ""
	s1 := ""
	s2 := ""
	s3 := ""
	s4 := ""

	for _, k := range h.TheCards {
		s0 += fmt.Sprintf("\t┌%s────┒", strings.Repeat("─", 3))
		s1 += fmt.Sprintf("\t│%s    ┃", k)
		s2 += fmt.Sprintf("\t│  %s  ┃", strings.Repeat(" ", 3))
		s3 += fmt.Sprintf("\t│    %s┃", k)
		s4 += fmt.Sprintf("\t┕━━━━%s┛", strings.Repeat("━", 3))
	}
	s0 += "\n"
	s1 += "\n"
	s2 += "\n"
	s3 += "\n"
	s4 += "\n"
	s += s0 + s1 + s2 + s2 + s3 + s4

	return
}
