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
	DealerGoesBust
	AllPlayersStay
	DetermineResults
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
		Initialized:            Game.NewHand,
		NewHand:                dealNewHand,
		NewHandDealt:           dealRounds,
		DetermineResults:       determineHandResults,
		AllPlayersStay:         determineHandResults,
		HandIsOver:             playAgain,
		PlayerWantsToPlayAgain: Game.NewHand,
		GameOver:               sayGoodbye,
	}
)
