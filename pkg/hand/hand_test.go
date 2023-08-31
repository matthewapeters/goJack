package hand_test

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/cucumber/godog"
	"github.com/matthewapeters/gojack/pkg/card"
	"github.com/matthewapeters/gojack/pkg/hand"
)

func aHandWithCards(arg1 string) (ctx context.Context, err error) {
	newHand := hand.NewHand()
	ctx = context.WithValue(context.Background(), "Hand", newHand)
	cards := strings.Split(arg1, ",")
	for _, c := range cards {
		switch c {
		case "2H":
			newHand.Takes(card.NewCard(card.Hearts, card.Two))
		case "10S":
			newHand.Takes(card.NewCard(card.Spades, card.Ten))
		case "AS":
			newHand.Takes(card.NewCard(card.Spades, card.Ace))
		case "AC":
			newHand.Takes(card.NewCard(card.Clubs, card.Ace))
		case "QS":
			newHand.Takes(card.NewCard(card.Spades, card.Queen))
		case "3S":
			newHand.Takes(card.NewCard(card.Spades, card.Three))
		case "4D":
			newHand.Takes(card.NewCard(card.Diamonds, card.Four))
		case "5C":
			newHand.Takes(card.NewCard(card.Clubs, card.Five))
		case "AD":
			newHand.Takes(card.NewCard(card.Diamonds, card.Ace))
		case "AH":
			newHand.Takes(card.NewCard(card.Hearts, card.Ace))

		default:
			return ctx, fmt.Errorf("cannot support card %s", c)
		}
	}
	return
}

func itMayHaveTheseBelowOrEqualTo(ctx context.Context, expectedValueStr string, max int) error {
	expectedValues := []int{}
	for _, pv := range strings.Split(expectedValueStr, ",") {
		if pv != "" {
			v, err := strconv.Atoi(pv)
			if err != nil {
				return err
			}
			expectedValues = append(expectedValues, v)
		}
	}
	hand := ctx.Value("Hand").(*hand.Hand)
	foundValues := hand.Values().Values
	if len(foundValues) != len(expectedValues) {
		return fmt.Errorf("Expected %d possible values but got %d values (%+v)", len(expectedValues), len(foundValues), foundValues)
	}
	sort.IntSlice(expectedValues).Sort()
	values := []int{}
	for _, fv := range foundValues {
		values = append(values, fv.Value)
	}
	sort.IntSlice(values).Sort()
	for i := range expectedValues {
		if expectedValues[i] != values[i] {
			return fmt.Errorf("Expected value %d but got %d", expectedValues[i], values[i])
		}
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^a hand with cards "([^"]*)"$`, aHandWithCards)
	ctx.Step(`^it may have these "([^"]*)" below or equal to (\d+)$`, itMayHaveTheseBelowOrEqualTo)
}
