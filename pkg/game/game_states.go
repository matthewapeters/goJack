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
	DetermineResults
	HandIsOver
	PlayerWantsToPlayAgain
	PlayerWantsToQuit
	GameOver
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
