package card_test

import (
	"context"
	"fmt"

	"github.com/cucumber/godog"
	"github.com/matthewapeters/gojack/pkg/card"
)

/*
	func aCardWithSuitOfClubsAndAFaceValueOfFour() (context.Context, error) {
		ctx := context.WithValue(context.Background(), "CARD", card.NewCard(card.Clubs, card.Four))
		return ctx, nil
	}

	func aCardWithSuitOfHeartsAndAFaceValueOfAce() (context.Context, error) {
		ctx := context.WithValue(context.Background(), "CARD", card.NewCard(card.Hearts, card.Ace))
		return ctx, nil
	}

	func aCardWithSuitOfSpadesAndAFaceValueOfThree() (context.Context, error) {
		ctx := context.WithValue(context.Background(), "CARD", card.NewCard(card.Spades, card.Three))
		return ctx, nil
	}

	func itsStringValueWillBeAceOfHearts(ctx context.Context) error {
		c := ctx.Value("CARD").(card.Card)
		expected := "Ace of Hearts"
		if c.String() != expected {
			return fmt.Errorf("EXPECTED %s got %s", expected, c.String())
		}
		return nil
	}

	func itsStringValueWillBeFourOfClubs(ctx context.Context) error {
		c := ctx.Value("CARD").(card.Card)
		expected := "Four of Clubs"
		if c.String() != expected {
			return fmt.Errorf("EXPECTED %s got %s", expected, c.String())
		}
		return nil

}

	func itsStringValueWillBeThreeOfSpades(ctx context.Context) error {
		c := ctx.Value("CARD").(card.Card)
		expected := "Three of Spades"
		if c.String() != expected {
			return fmt.Errorf("EXPECTED %s got %s", expected, c.String())
		}
		return nil

}

	func theCardIsStringified(ctx context.Context) error {
		return nil
	}

	func InitializeScenario(ctx *godog.ScenarioContext) {
		ctx.Step(`^the card is stringified$`, theCardIsStringified)
		ctx.Step(`^a card with Suit of Clubs and a face value of Four$`, aCardWithSuitOfClubsAndAFaceValueOfFour)
		ctx.Step(`^a card with Suit of Hearts and a face value of Ace$`, aCardWithSuitOfHeartsAndAFaceValueOfAce)
		ctx.Step(`^a card with Suit of Spades and a face value of Three$`, aCardWithSuitOfSpadesAndAFaceValueOfThree)
		ctx.Step(`^its string value will be Ace of Hearts$`, itsStringValueWillBeAceOfHearts)
		ctx.Step(`^its string value will be Four of Clubs$`, itsStringValueWillBeFourOfClubs)
		ctx.Step(`^its string value will be Three of Spades$`, itsStringValueWillBeThreeOfSpades)
	}
*/
func aCardWithSuitOfClubsAndAFaceValueOfFour() (context.Context, error) {
	ctx := context.WithValue(context.Background(), "CARD", card.NewCard(card.Clubs, card.Four))
	return ctx, nil
}

func aCardWithSuitOfHeartsAndAFaceValueOfAce() (context.Context, error) {
	ctx := context.WithValue(context.Background(), "CARD", card.NewCard(card.Hearts, card.Ace))
	return ctx, nil
}

func aCardWithSuitOfSpadesAndAFaceValueOfThree() (context.Context, error) {
	ctx := context.WithValue(context.Background(), "CARD", card.NewCard(card.Spades, card.Three))
	return ctx, nil
}

func itsStringValueWillBeAce(ctx context.Context) error {
	c := ctx.Value("CARD").(*card.Card)
	expected := fmt.Sprintf("%s%s", card.Hearts, card.Ace)
	if c.String() != expected {
		return fmt.Errorf("EXPECTED %s got %s", expected, c.String())
	}
	return nil
}

func itsStringValueWillBeFour(ctx context.Context) error {
	c := ctx.Value("CARD").(*card.Card)
	expected := fmt.Sprintf("%s%s", card.Clubs, card.Four)
	if c.String() != expected {
		return fmt.Errorf("EXPECTED %s got %s", expected, c.String())
	}
	return nil
}

func itsStringValueWillBeThree(ctx context.Context) error {
	c := ctx.Value("CARD").(*card.Card)
	expected := fmt.Sprintf("%s%s", card.Spades, card.Three)
	if c.String() != expected {
		return fmt.Errorf("EXPECTED %s got %s", expected, c.String())
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^a card with Suit of Clubs and a face value of Four$`, aCardWithSuitOfClubsAndAFaceValueOfFour)
	ctx.Step(`^a card with Suit of Hearts and a face value of Ace$`, aCardWithSuitOfHeartsAndAFaceValueOfAce)
	ctx.Step(`^a card with Suit of Spades and a face value of Three$`, aCardWithSuitOfSpadesAndAFaceValueOfThree)
	ctx.Step(`^its string value will be ♥ Ace$`, itsStringValueWillBeAce)
	ctx.Step(`^its string value will be ♣ Four$`, itsStringValueWillBeFour)
	ctx.Step(`^its string value will be ♠ Three$`, itsStringValueWillBeThree)
}
