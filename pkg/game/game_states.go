package game

type GameState int
type Choice int

type StateMachine map[GameState]func()

const (
	NewGame = GameState(iota)
	Initialized
	NewHand
	//NewHandDealt
	DealToPlayer
	DealtARound
	DealToDealer
	PromptPlayer
	PlayerStays
	PlayerTakesCard
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
		NewGame:     initializeGame,
		Initialized: startNewHand,
		NewHand:     dealNewHand,
		//NewHandDealt:           dealToPlayer,
		DealToPlayer:           dealToPlayer,
		DealToDealer:           dealToDealer,
		PromptPlayer:           playerChooses,
		PlayerStays:            nextPlayersTurn,
		PlayerTakesCard:        playerTakesCard,
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
