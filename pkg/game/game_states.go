package game

type GameState int

const (
	NewGame = GameState(iota)
	PlayerChoiceHit
	DealerChoiceHit
	PlayerChoiceStand
	DealerChoicestand
	PlayerChoiceSplit
	PlayerGoesBust
	DealerGoesBust
)
