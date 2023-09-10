package player

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cucumber/godog"
	"github.com/matthewapeters/gojack/pkg/hand"
)

func aPlayerWith(nbrHands string) (ctx context.Context, err error) {
	nbr, err := strconv.Atoi(nbrHands)
	if err != nil {
		return
	}
	p := NewPlayer(fmt.Sprintf("TestPlayerWith%sHands", nbrHands))
	for i := 1; i < nbr; i++ {
		p.Hands = append(p.Hands, hand.NewHand())
	}
	ctx = context.WithValue(context.Background(), "PLAYER", p)
	//fmt.Printf("expect %s hands, created %d hands", nbrHands, len(p.Hands))
	return
}

func hasMoreHandsWillRespondWith(ctx context.Context, expected string) (err error) {
	p := ctx.Value("PLAYER").(*Player)
	found := fmt.Sprintf("%t", p.HasMoreHands())
	if found != expected {
		err = fmt.Errorf("expected %s but got %s", expected, found)
		return
	}
	return
}

func thePlayerInvokesNextHand(ctx context.Context) (err error) {
	p := ctx.Value("PLAYER").(*Player)
	//prev := p.HandIndex()
	p.NextHand()
	//new := p.HandIndex()
	//fmt.Printf("previous hand index: %d new hand index %d\n", prev, new)
	return
}

func thePlayersCurrentHandIndexIs(ctx context.Context, idxString string) (err error) {
	p := ctx.Value("PLAYER").(*Player)
	idx, err := strconv.Atoi(idxString)
	if err != nil {
		return
	}
	p.Hand = p.Hands[idx]
	//fmt.Printf("expect index: %s actual: %d\n", idxString, p.HandIndex())
	return
}

func thePlayersResultingHandIndexIs(ctx context.Context, expIdx string) (err error) {
	p := ctx.Value("PLAYER").(*Player)
	found := fmt.Sprintf("%d", p.HandIndex())
	if found != expIdx {
		err = fmt.Errorf("expected index %s, got index %s", expIdx, found)
	}
	return
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^a player with "([^"]*)"$`, aPlayerWith)
	ctx.Step(`^HasMoreHands will respond with "([^"]*)"$`, hasMoreHandsWillRespondWith)
	ctx.Step(`^the player invokes NextHand$`, thePlayerInvokesNextHand)
	ctx.Step(`^the player\'s current hand index is "([^"]*)"$`, thePlayersCurrentHandIndexIs)
	ctx.Step(`^the player\'s resulting hand index is "([^"]*)"$`, thePlayersResultingHandIndexIs)
}
