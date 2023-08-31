package deck_test

import (
	"context"
	"fmt"
	"strings"

	"github.com/cucumber/godog"
	"github.com/matthewapeters/gojack/pkg/deck"
)

func aDeckOfCards() (ctx context.Context, err error) {
	d := deck.NewDeck()
	ctx = context.WithValue(context.Background(), "DECK", d)
	fmt.Printf("%s", d.Cards)
	return
}

func eachCardInEachValueWillBeAUniqueSuitOfSuit(arg1 *godog.Table) error {
	return godog.ErrPending
}

func thereWillBeCardsOfEachFaceValue(arg1 int, arg2 *godog.Table) error {
	return godog.ErrPending
}

func itsStringValuesWillBe(ctx context.Context, arg1 *godog.DocString) error {
	d := ctx.Value("DECK").(*deck.DeckOfCards)
	prntOut := d.Cards.String()

	expectedLines := strings.Split(arg1.Content, "\n")
	foundLines := strings.Split(prntOut, "\n")
	for line := 0; line < len(expectedLines); line += 1 {
		if expectedLines[line] != foundLines[line] {
			return fmt.Errorf("Expected '%s' on line %d but got '%s'", expectedLines[line], line, foundLines[line])
		}
	}
	if len(prntOut) != len(arg1.Content) {
		return fmt.Errorf("Expected length %d got length %d", len(arg1.Content), len(prntOut))
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^its string values will be$`, itsStringValuesWillBe)
	ctx.Step(`^a deck of cards$`, aDeckOfCards)
	ctx.Step(`^each card in each value will be a unique Suit of <Suit>$`, eachCardInEachValueWillBeAUniqueSuitOfSuit)
	ctx.Step(`^there will be (\d+) cards of each  <FaceValue>$`, thereWillBeCardsOfEachFaceValue)
}
