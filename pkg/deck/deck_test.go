package deck_test

import (
	"context"
	"fmt"
	"strings"

	"github.com/cucumber/godog"
	"github.com/matthewapeters/gojack/pkg/deck"
)

func aDeckOfCards() (ctx context.Context, err error) {
	ctx = context.WithValue(context.Background(), "DECK", deck.NewDeck())
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
	prntOut := d.String()

	expectedLines := strings.Split(arg1.Content, "\n")
	foundLines := strings.Split(prntOut, "\n")
	for line := 0; line < len(expectedLines); line += 1 {
		if expectedLines[line] != foundLines[line] {
			return fmt.Errorf("Expected '%s' on line %d but got '%s'", expectedLines[line], line, foundLines[line])
		}
	}
	if len(prntOut) != len(arg1.Content) {
		for i := 0; i < max(len(prntOut), len(arg1.Content)); i++ {
			if i > len(prntOut)-1 {
				fmt.Printf("output has no %d-th character.  In expected, this is '%s'\n", i, string(arg1.Content[i]))
			} else {
				if i > len(arg1.Content)-1 {
					fmt.Printf("expected has no %d-th character.  In output, this is '%s'\n", i, string(prntOut[i]))
				} else {
					if prntOut[i] != arg1.Content[i] {
						fmt.Printf("character %d: Expected: '%s'  Got: '%s'\n", i, string(arg1.Content[i]), string(prntOut[i]))
					}
				}
			}

		}
		return fmt.Errorf("Expected length %d got length %d", len(arg1.Content), len(prntOut))
	}
	return nil
}

func eachCardInEachValueWillBeAUniqueSuitOf(arg1 string, arg2 *godog.Table) error {
	return godog.ErrPending
}

func thereWillBeCardsOfEach(arg1 int, arg2 string, arg3 *godog.Table) error {
	return godog.ErrPending
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^each card in each value will be a unique Suit of "([^"]*)"$`, eachCardInEachValueWillBeAUniqueSuitOf)
	ctx.Step(`^there will be (\d+) cards of each  "([^"]*)"$`, thereWillBeCardsOfEach)
	ctx.Step(`^its string values will be$`, itsStringValuesWillBe)
	ctx.Step(`^a deck of cards$`, aDeckOfCards)
	ctx.Step(`^each card in each value will be a unique Suit of <Suit>$`, eachCardInEachValueWillBeAUniqueSuitOfSuit)
	ctx.Step(`^there will be (\d+) cards of each  <FaceValue>$`, thereWillBeCardsOfEachFaceValue)
}
