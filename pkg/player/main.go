package player

import (
	"strings"

	"github.com/matthewapeters/gojack/pkg/hand"
	"github.com/matthewapeters/gojack/pkg/purse"
)

type Choice string

const (
	MIN   = "MIN"
	MAX   = "MAX"
	HIT   = Choice("HIT")
	STAY  = Choice("STAY")
	SPLIT = Choice("SPLIT")
)

type Player struct {
	Name  string
	Hands []*hand.Hand
	Hand  *hand.Hand
	Choice
	Purse purse.Purse
}

func NewPlayer(name string) *Player {
	h := hand.NewHand()
	return &Player{
		Name:   name,
		Hands:  []*hand.Hand{h},
		Hand:   h,
		Choice: HIT,
		Purse:  purse.NewPurse(1000),
	}
}

func (d *Player) GoesBust() bool {
	if len(d.Hand.TheScore) == 0 {
		return true
	}
	return false
}

func (p *Player) HitsTwentyOne() bool {
	return p.Scores()[MAX] == hand.MaxValue
}

func (p *Player) NewGame() {
	p.Hand = hand.NewHand()
	p.Hands = []*hand.Hand{p.Hand}
	p.Choice = HIT
}

func (p *Player) Scores() (scores map[string]int) {
	scores = map[string]int{MIN: 999, MAX: 0}
	for _, v := range p.Hand.TheScore {
		if v < scores[MIN] {
			scores[MIN] = v
		}
		if v > scores[MAX] {
			scores[MAX] = v
		}
	}
	return
}

func (p *Player) MakeChoice(input string) bool {
	switch strings.ToUpper(string(input[0])) {
	case "S":
		p.Choice = STAY
		return true
	case "H":
		p.Choice = HIT
		return true
	case "P":
		p.Choice = SPLIT
		return true
	}
	return false
}

func (p *Player) HasMoreHands() (answer bool) {
	answer = p.HandIndex() < len(p.Hands)-1

	return
}

func (p *Player) NextHand() {
	for i := 0; i < len(p.Hands)-1; i++ {
		if p.Hand == p.Hands[i] {
			p.Hand = p.Hands[i+1]
			return
		}
	}
}

func (p *Player) HandIndex() int {
	for i := 0; i < len(p.Hands); i++ {
		if p.Hand == p.Hands[i] {
			return i
		}
	}
	return -1
}

type Players []*Player
