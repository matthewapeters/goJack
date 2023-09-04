package card_test

import (
	"context"
	"fmt"

	"github.com/cucumber/godog"
	"github.com/matthewapeters/gojack/pkg/card"
)

func aCardWithSuitOfAndAFaceValueOf(suit, facevalue string) (ctx context.Context, err error) {
	var s card.Suit
	var fv card.FaceValue
	suits := map[string]card.Suit{
		"Hearts":   card.Hearts,
		"Spades":   card.Spades,
		"Diamonds": card.Diamonds,
		"Clubs":    card.Clubs}
	facevalues := map[string]card.FaceValue{"Two": card.Two, "Three": card.Three,
		"Four": card.Four, "Five": card.Five, "Six": card.Six,
		"Seven": card.Seven, "Eight": card.Eight, "Nice": card.Nine,
		"Ten": card.Ten, "Jack": card.Jack, "Queen": card.Queen, "King": card.King, "Ace": card.Ace}

	s = suits[suit]
	fv = facevalues[facevalue]
	crd := card.NewCard(s, fv)
	ctx = context.WithValue(context.Background(), "Card", crd)
	return
}

func itsStringValueWillBe(ctx context.Context, expected string) (err error) {
	card := ctx.Value("Card").(*card.Card)
	found := card.String()
	if found != expected {
		err = fmt.Errorf("expected %s but got %s", expected, found)
	}
	return
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^a card with Suit of "([^"]*)" and a face value of "([^"]*)"$`, aCardWithSuitOfAndAFaceValueOf)
	ctx.Step(`^its string value will be "([^"]*)""$`, itsStringValueWillBe)
}
