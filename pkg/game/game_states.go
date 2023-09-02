package game

type GameState int
type Choice int

type StateMachine map[GameState]func()

const (
	NewGame = GameState(iota)
	Initialized
	NewHand
	NewHandDealt
	DoingRounds
	PlayerGoesBust
	AllPlayerGoBust
	DealerGoesBust
	AllPlayersStay
	DetermineResults
	GameHasWinner // Dealer has gone bust
	HandIsOver
	PlayerWantsToPlayAgain
	GameOver

	PlayerChoiceHit = Choice(iota)
	DealerChoiceHit
	PlayerChoiceStay
	DealerChoicestay
	PlayerChoiceSplit
)

var (
	GameStateMachine = StateMachine{
		NewGame:                initializeGame,
		Initialized:            startNewHand,
		NewHand:                dealNewHand,
		NewHandDealt:           dealRounds,
		PlayerGoesBust:         determineHandResults,
		DealerGoesBust:         determineHandResults,
		DetermineResults:       determineHandResults,
		AllPlayersStay:         determineHandResults,
		HandIsOver:             playAgain,
		PlayerWantsToPlayAgain: startNewHand,
		GameOver:               sayGoodbye,
	}
)
