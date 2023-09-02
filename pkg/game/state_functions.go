package game

import (
	"fmt"
	"strings"
	"time"

	"github.com/matthewapeters/gojack/pkg/player"
)

// initializeGame
// creates new players
// Triggering States: NewGame
// Resulting States: NewHand
func initializeGame() {
	for i, name := range Game.Names {
		if Game.Players[i] == nil {
			Game.Players[i] = player.NewPlayer(name)
		} else {
			Game.Players[i].Name = name
		}
	}
	Game.State = NewHand
}

// startNewHand -
// Gathers used cards from all players and prepares players and Dealer for a new game
// Trigger States: Initialized, PlayerWantsToPlayAgain
// Resulting States: NewHand
func startNewHand() {
	Game.Results = ""
	// Gather all of the cards from the players
	for _, p := range Game.Players {
		Game.Dealer.GatherPlayedCards(*p.Hand.TheCards)
		p.NewGame()
	}

	// Dealer initiates new game by returning played cards to deck and shuffling the deck, if necessary
	Game.Dealer.NewGame()

	//Game.HasWinner = false
	Game.State = NewHand
}

// dealNewHand
// Trigger States: NewHand
// Resulting States: NewHandDealt
func dealNewHand() {
	// Dealer deals 2 cards to each player
	for i := 0; i < 2; i++ {
		for _, player := range Game.Players {
			player.Hand.Takes(Game.Dealer.Deck.Cards.GiveCard(1))
		}
		Game.Dealer.Player.Hand.Takes(Game.Dealer.Deck.Cards.GiveCard(1))
		// The first dealer's card is dealt face-down
		if i == 0 {
			(*Game.Dealer.Player.Hand.TheCards)[0].FaceDown = true
		}
	}
	Game.State = NewHandDealt
}

/*
// dealRounds
// Deals new round of one card per player until players stay or go broke
// Trigger States: NewHandDealt
// Resulting States: AllPlayersStay, DetermineResults
func dealRounds() {
	keepDealing := true
	for keepDealing {
		dealRound()
		switch Game.State {
		case DetermineResults:
			keepDealing = false
		case AllPlayersStay:
			keepDealing = false
		case GameHasWinner:
			keepDealing = false
		case DealerGoesBust:
			keepDealing = false
		case PlayerGoesBust:
			keepDealing = false
		case AllPlayerGoBust:
			keepDealing = false
		}
		if !keepDealing {
			return
		}
		// if both players have chosen to stay, rounds are over, determine player with highest score
		determineIfAllPlayersStay()
	}
}
*/

// dealRound
// internal function called by dealRounds - would like to flatten these!
// contains call to Game.ShowCards()
// Resulting States: DetermineResults, GameHasWinner, DealerGoesBust, PlayerGoesBust
func dealRound() {
	Game.ShowCards()

	// If the dealer is showing 21 or more, then they must have busted - all other players win
	if Game.State != GameHasWinner && Game.Dealer.Player.Scores()[player.MIN] >= 21 {
		Game.State = DealerGoesBust
		Game.Results += fmt.Sprintf("The Dealer is showing %d - the Dealer Goes Bust!\n",
			Game.Dealer.Player.Scores()[player.MIN])
		for _, p := range Game.NotBustedPlayers() {
			Game.Scores[p] += 1
		}
	}

	if Game.State != GameHasWinner && Game.Dealer.GoesBust() {
		Game.Results = "Dealer Goes Bust!"
		for _, player := range Game.Players {
			Game.Scores[player] += 1
		}
		Game.State = DealerGoesBust
	}

	if Game.State == DealerGoesBust {
		return
	}

	// if Player goes bust, the dealer wins
	for _, p := range Game.Players {
		if Game.State != GameHasWinner && p.GoesBust() {
			Game.Results += fmt.Sprintf("%s Goes Bust!\n", p.Name)
			Game.State = PlayerGoesBust
		}
		if Game.State == GameHasWinner {
			Game.Scores[Game.Dealer.Player] += 1
		}
	}

	if Game.State == PlayerGoesBust {
		return
	}
	// Player gets to choose next move first
	for _, p := range Game.NotBustedPlayers() {
		choiceMade := false
		for !choiceMade && p.Choice == player.HIT {
			fmt.Printf("%s: (H)it or (S)tay? ", p.Name)
			var choice string
			fmt.Scanln(&choice)
			choiceMade = p.MakeChoice(choice)
		}
		if p.Choice == player.HIT {
			p.Hand.Takes(Game.Dealer.Deck.Cards.GiveCard(1))
			Game.ShowCards()
		}
	}

	// Dealer decides to hit or stay
	if Game.Dealer.Player.Scores()[player.MAX] < 17 {
		fmt.Printf("Dealer takes a card.")
		Game.Dealer.Player.Hand.Takes(Game.Dealer.Deck.Cards.GiveCard(1))
		time.Sleep(3 * time.Second)
	} else {
		fmt.Printf("Dealer Shows %d - cannot hit over 17!\n", Game.Dealer.Player.Scores()[player.MAX])
		Game.Dealer.Player.Choice = player.STAY
		time.Sleep(3 * time.Second)
	}
	Game.State = DealtARound
}

// deternubeUfAllPlayersStay
// Triggering State:
// ResultingStates: AllPlayersStay, DealARound
func determineIfAllPlayersStay() {
	if Game.AllStay() {
		Game.State = AllPlayersStay
		fmt.Println("All Players Have Chosen to Stay")
		time.Sleep(3 * time.Second)
	} else {
		Game.State = DealARound
	}
}

// determineResults
// contains call to Game.ShowCards()
// Triggering States: PlayerGoesBust, DealerGoesBust, DetermineResults, AllPlayersStay
// Resulting States: HandIsOver
func determineHandResults() {
	// If the players have all busted, the dealer does not have to show card!
	if !Game.AllPlayersBusted() {
		// If the dealer is showing a hard 21 without revealing, then dealer has busted
		if Game.Dealer.GoesBust() {
			for _, p := range Game.Players {
				Game.Scores[p] += 1
			}
		} else {
			Game.Dealer.RevealFirstCard()
			Game.ShowCards()
		}
	}
	// Evaluate final scores
	if Game.AllStay() {
		if Game.Dealer.GoesBust() {
			Game.Results += "Dealer Goes Bust!\n"
			for _, p := range Game.Players {
				Game.Scores[p] += 1
			}
		} else {
			for _, p := range Game.Players {
				if Game.Dealer.Player.Scores()[player.MAX] >= p.Scores()[player.MAX] {
					Game.Scores[Game.Dealer.Player] += 1
					Game.Results += fmt.Sprintf("Dealer has %d, %s has %d.  Dealer Wins!\n",
						Game.Dealer.Player.Scores()[player.MAX],
						p.Name,
						p.Scores()[player.MAX])
				} else {
					Game.Scores[p] += 1
					Game.Results += fmt.Sprintf("Dealer has %d, %s has %d.  %s Wins!\n",
						Game.Dealer.Player.Scores()[player.MAX],
						p.Name,
						p.Scores()[player.MAX],
						p.Name)
				}
			}
		}
	}
	// show final game result
	fmt.Println(Game.Results)
	// show total game scores:
	fmt.Printf("PLAYER:\t\tSCORE:\n")
	for k, v := range Game.Scores {
		fmt.Printf("%s:\t\t%d\n", k.Name, v)
	}
	Game.State = HandIsOver
}

// playAgain
// determines if players want to continue playing
// Triggering States:
// Resulting States: PalyerWantsToPlayAgain, GameOver
func playAgain() {
	for _, p := range Game.Players {
		choice := "N"
		choiceMade := false
		for !choiceMade {
			fmt.Printf("\n%s: Play Again? (Y/N) ", p.Name)
			fmt.Scanln(&choice)
			choice = strings.ToUpper(string(choice[0]))
			switch choice {
			case "Y":
				choiceMade = true
				Game.State = PlayerWantsToPlayAgain
			case "N":
				choiceMade = true
				Game.State = PlayerWantsToQuit
			default:
				fmt.Println("Dont be goof!")
				choiceMade = false
			}
		}
	}
}

// sayGoodbye
// goodbye message to players
// Triggereing States: GameOver
func sayGoodbye() {
	for _, p := range Game.Players {
		fmt.Printf("\nGoodbye, %s!  Play GoJack again soon!\n\n", p.Name)
	}
	Game.State = GameOver
}
