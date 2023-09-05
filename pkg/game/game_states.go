package game

type GameState int
type Choice int

type StateMachine map[GameState]func()

const (
	NewGame = GameState(iota)
	Initialized
	NewHand
	NewHandDealt
	DealARound
	DealtARound
	DealToDealer
	PlayerGoesBust
	AllPlayersGoBust
	DealerGoesBust
	AllPlayersStay
	DetermineResults
	HandIsOver
	PlayerWantsToPlayAgain
	PlayerWantsToQuit
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
		NewHandDealt:           dealToPlayer,
		DealARound:             dealToPlayer,
		DealToDealer:           dealToDealer,
		DealtARound:            determineHandResults,
		PlayerGoesBust:         determineIfAllPlayersBusted,
		AllPlayersGoBust:       dealerWins,
		DealerGoesBust:         dealerGoesBust,
		DetermineResults:       determineHandResults,
		HandIsOver:             playAgain,
		PlayerWantsToPlayAgain: startNewHand,
		PlayerWantsToQuit:      sayGoodbye,
	}
)
