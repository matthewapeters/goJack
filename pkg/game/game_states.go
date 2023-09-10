package game

type GameState int
type Choice int

type StateMachine map[GameState]func()

const (
	NewGame = GameState(iota)
	Initialized
	NewHand
	DealToPlayer
	DealtARound
	DealToDealer
	PromptPlayer
	PlayerStays
	PlayerTakesCard
	PlayerSplitsHand
	PlayerGoesBust
	AllPlayersGoBust
	DealerGoesBust
	DetermineResults
	HandIsOver
	PlayerWantsToPlayAgain
	PlayerWantsToQuit
	GameOver
)

func (gs GameState) String() string {
	switch gs {
	case NewGame:
		return "NewGame"
	case Initialized:
		return "Initialized"
	case NewHand:
		return "NewHand"
	case DealToPlayer:
		return "DealToPlayer"
	case DealtARound:
		return "DealtARound"
	case DealToDealer:
		return "DealToDealer"
	case PromptPlayer:
		return "PromptPlayer"
	case PlayerStays:
		return "PlayerStays"
	case PlayerTakesCard:
		return "PlayerTakesCard"
	case PlayerSplitsHand:
		return "PlayerSplitsHand"
	case PlayerGoesBust:
		return "PlayerGoesBust"
	case AllPlayersGoBust:
		return "AllPlayersGoBust"
	case DealerGoesBust:
		return "DealerGoesBust"
	case DetermineResults:
		return "DetermineResults"
	case HandIsOver:
		return "HandIsOver"
	case PlayerWantsToPlayAgain:
		return "PlayerWantsToPlayAgain"
	case PlayerWantsToQuit:
		return "PlayerWantsToQuit"
	case GameOver:
		return "GameOver"
	}
	return "unknownState"
}

var (
	GameStateMachine = StateMachine{
		NewGame:                initializeGame,
		Initialized:            startNewHand,
		NewHand:                dealNewHand,
		DealToPlayer:           dealToPlayer,
		DealToDealer:           dealToDealer,
		PromptPlayer:           playerChooses,
		PlayerStays:            nextPlayersTurn,
		PlayerTakesCard:        playerTakesCard,
		PlayerSplitsHand:       playerSplitsHand,
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
