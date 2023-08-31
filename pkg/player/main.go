package player

import (
	"github.com/matthewapeters/gojack/pkg/hand"
)

type Choice string

const (
	MIN  = "MIN"
	MAX  = "MAX"
	HIT  = Choice("HIT")
	STAY = Choice("STAY")
)

type Player struct {
	Name string
	Hand *hand.Hand
	Choice
}

func NewPlayer(name string) *Player {
	return &Player{
		Name:   name,
		Hand:   hand.NewHand(),
		Choice: HIT,
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
	p.Choice = HIT
}

func (p *Player) Scores() (scores map[string]int) {
	scores = map[string]int{MIN: 22, MAX: 0}
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
	switch string(input[0]) {
	case "S":
		p.Choice = STAY
		return true
	case "s":
		p.Choice = STAY
		return true
	case "H":
		p.Choice = HIT
		return true
	case "h":
		p.Choice = HIT
		return true
	}
	return false
}
