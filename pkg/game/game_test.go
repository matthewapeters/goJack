package game

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/cucumber/godog"
	"github.com/matthewapeters/gojack/pkg/card"
	"github.com/matthewapeters/gojack/pkg/dealer"
	"github.com/matthewapeters/gojack/pkg/player"
)

func aNewGameOfGoJack() (ctx context.Context, err error) {

	theGame = &game{
		Dealer:  dealer.NewDealer(),
		Players: player.Players{},
		Scores:  map[*player.Player]int{},
	}
	theGame.DisplayInterval = (0 * time.Second)
	ctx = context.WithValue(context.Background(), "GAME", theGame)
	return
}

func aPlayerListOf(ctxIn context.Context, listOfPlayers string) (ctx context.Context, err error) {
	ctx = ctxIn
	players := strings.Split(listOfPlayers, " ")
	g := ctx.Value("GAME").(*game)
	g.Names = players
	initializeGame()
	return
}

func allOfThePlayerCardsWillBeFacingUp(ctx context.Context) (err error) {
	g := ctx.Value("GAME").(*game)
	expectedPlayerCardsFaceDown := 0
	foundPlayerCardsFaceDown := 0
	for _, p := range g.Players {
		for _, c := range *p.Hand.TheCards {
			if c.FaceDown {
				foundPlayerCardsFaceDown += 1
			}
		}

	}
	if foundPlayerCardsFaceDown != expectedPlayerCardsFaceDown {
		err = fmt.Errorf("Expecetd %d player cards facing down, found %d facing down", expectedPlayerCardsFaceDown, foundPlayerCardsFaceDown)
	}
	return
}

func theDealerDealsANewHand(ctx context.Context) (err error) {
	theGame = ctx.Value("GAME").(*game)
	dealNewHand()
	return
}

func theDealerWillHaveCardFacingDown(ctx context.Context, expectedBNbrCardsFacingDown int) (err error) {
	g := ctx.Value("GAME").(*game)
	foundCount := 0
	for _, c := range *g.Dealer.Player.Hand.TheCards {
		if c.FaceDown {
			foundCount += 1
		}
	}
	if foundCount != expectedBNbrCardsFacingDown {
		err = fmt.Errorf("Expected %d cards facing down in dealer's hand, found %d cards", expectedBNbrCardsFacingDown, foundCount)
	}
	return
}

func theDeckWillHaveRemainingCardsInTheDeck(ctx context.Context, expectedNbrCardsInDeckString string) (err error) {
	expectedNbrCardsInDeck, err := strconv.Atoi(expectedNbrCardsInDeckString)
	if err != nil {
		err = fmt.Errorf(" could not convert '%s' to integer, %s", expectedNbrCardsInDeckString, err)
		return
	}
	g := ctx.Value("GAME").(*game)
	foundCardsInDect := len([]*card.Card(*g.Dealer.Deck.Cards.TheCards))
	if foundCardsInDect != expectedNbrCardsInDeck {
		err = fmt.Errorf("expected %d cards remaining in deck, found %d", expectedNbrCardsInDeck, foundCardsInDect)
	}
	return
}

func theDealerHasAHandWithTheCards(ctx context.Context, cards string) (ctxOut context.Context, err error) {
	ctxOut = ctx
	g := ctx.Value("GAME").(*game)

	for _, c := range strings.Split(cards, ",") {
		var crd *card.Card

		switch c {
		case "2s":
			crd = card.NewCard(card.Spades, card.Two)
			crd.FacingDown()
		case "2S":
			crd = card.NewCard(card.Spades, card.Two)
		case "10s":
			crd = card.NewCard(card.Spades, card.Ten)
			crd.FacingDown()
		case "10S":
			crd = card.NewCard(card.Spades, card.Ten)
		case "Ah":
			crd = card.NewCard(card.Hearts, card.Ace)
			crd.FacingDown()
		case "AH":
			crd = card.NewCard(card.Hearts, card.Ace)

		default:
			err = fmt.Errorf("cannot handle card %s", c)
			return

		}
		g.Dealer.Player.Hand.Takes(crd)
	}

	return
}

func theDealerMustDecideToHitOrStay(ctx context.Context) (err error) {
	theGame = ctx.Value("GAME").(*game)
	theGame.DisplayInterval = (0 * time.Second)
	theGame.SupressDisplay = true
	dealToDealer()
	return
}

func theGameStateIs(ctx context.Context, stateName string) (err error) {
	g := ctx.Value("GAME").(*game)
	switch stateName {
	case "DealToDealer":
		g.State = DealToDealer
	default:
		err = fmt.Errorf("state %s not handled", stateName)
	}
	return
}

func theResultingGameStateWillBe(ctx context.Context, stateName string) (err error) {
	g := ctx.Value("GAME").(*game)
	foundState := g.State
	var expectedState GameState

	switch stateName {
	case "DealToDealer":
		expectedState = DealToDealer
	case "DealerGoesBust":
		expectedState = DealerGoesBust
	case "DealtARound":
		expectedState = DealtARound
	default:
		err = fmt.Errorf("state %s not handled", stateName)
	}
	if foundState != expectedState {
		err = fmt.Errorf("expected state of %s, but found state is %d", stateName, foundState)
	}

	return
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^the Dealer has a hand with the cards "([^"]*)"$`, theDealerHasAHandWithTheCards)
	ctx.Step(`^the Dealer must decide to hit or Stay$`, theDealerMustDecideToHitOrStay)
	ctx.Step(`^the game State is "([^"]*)"$`, theGameStateIs)
	ctx.Step(`^the resulting game state will be "([^"]*)"$`, theResultingGameStateWillBe)
	ctx.Step(`^a new game of goJack$`, aNewGameOfGoJack)
	ctx.Step(`^a player list of "([^"]*)"$`, aPlayerListOf)
	ctx.Step(`^all of the player cards will be facing up$`, allOfThePlayerCardsWillBeFacingUp)
	ctx.Step(`^the Dealer deals a new hand$`, theDealerDealsANewHand)
	ctx.Step(`^the Dealer will have (\d+) card facing down$`, theDealerWillHaveCardFacingDown)
	ctx.Step(`^the deck will have "([^"]*)" in the deck$`, theDeckWillHaveRemainingCardsInTheDeck)
}
