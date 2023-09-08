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

func getCard(cardString string) (crd *card.Card, err error) {
	switch cardString {
	case "10H":
		crd = card.NewCard(card.Hearts, card.Ten)
	case "8S":
		crd = card.NewCard(card.Spades, card.Eight)
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
	case "4C":
		crd = card.NewCard(card.Clubs, card.Four)

	default:
		err = fmt.Errorf("cannot handle card %s", cardString)
	}
	return
}

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
	var crd *card.Card
	ctxOut = ctx
	g := ctx.Value("GAME").(*game)

	for _, c := range strings.Split(cards, ",") {
		crd, err = getCard(c)
		if err != nil {
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

func itIsThePlayersTurn(ctx context.Context) (err error) {
	g := ctx.Value("GAME").(*game)
	g.SupressDisplay = true
	g.DisplayInterval = (0 * time.Second)
	theGame = g

	dealToPlayer()
	return
}

func theDealerDealsAToThePlayer(ctx context.Context, cardString string) (err error) {
	g := ctx.Value("GAME").(*game)
	p := g.Players[g.CurrentPlayerID]
	var crd *card.Card
	crd, err = getCard(cardString)
	if err != nil {
		return
	}
	p.Hand.Takes(crd)
	g.State = DealARound
	return
}

func theGameStateWillBe(ctx context.Context, expectedState string) (err error) {
	g := ctx.Value("GAME").(*game)
	var expState GameState
	switch expectedState {
	case "PromptPlayer":
		expState = PromptPlayer
	case "PlayerGoesBust":
		expState = PlayerGoesBust
	default:
		err = fmt.Errorf("game state %s not handled yet", expectedState)
		return
	}
	if g.State != expState {
		err = fmt.Errorf("Expected State of %s, but got state %d", expectedState, g.State)
	}
	return
}

func thePlayerChoosesToHit(ctx context.Context) (err error) {
	g := ctx.Value("GAME").(*game)
	g.State = PlayerTakesCard
	return
}

func thePlayerHasAHandWithTheCards(ctx context.Context, cards string) (err error) {
	g := ctx.Value("GAME").(*game)
	p := g.Players[theGame.CurrentPlayerID]
	var crd *card.Card
	for _, c := range strings.Split(cards, ",") {
		crd, err = getCard(c)
		if err != nil {
			return
		}
		p.Hand.Takes(crd)
	}
	g.State = DealARound
	return
}

func thePlayersHandIsCounted(ctx context.Context) (err error) {
	g := ctx.Value("GAME").(*game)
	theGame = g
	dealToPlayer()
	return
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^it is the player\'s turn$`, itIsThePlayersTurn)
	ctx.Step(`^the dealer deals a "([^"]*)" to the player$`, theDealerDealsAToThePlayer)
	ctx.Step(`^the game state will be "([^"]*)"$`, theGameStateWillBe)
	ctx.Step(`^the player chooses to Hit$`, thePlayerChoosesToHit)
	ctx.Step(`^the player has a hand with the cards "([^"]*)"$`, thePlayerHasAHandWithTheCards)
	ctx.Step(`^the player\'s hand is counted$`, thePlayersHandIsCounted)

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
