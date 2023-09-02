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
	for i, name := range theGame.Names {
		if theGame.Players[i] == nil {
			theGame.Players[i] = player.NewPlayer(name)
		} else {
			theGame.Players[i].Name = name
		}
	}
	theGame.State = NewHand
}

// startNewHand -
// Gathers used cards from all players and prepares players and Dealer for a new game
// Trigger States: Initialized, PlayerWantsToPlayAgain
// Resulting States: NewHand
func startNewHand() {
	theGame.Results = ""
	// Gather all of the cards from the players
	for _, p := range theGame.Players {
		theGame.Dealer.GatherPlayedCards(*p.Hand.TheCards)
		p.NewGame()
	}

	// Dealer initiates new game by returning played cards to deck and shuffling the deck, if necessary
	theGame.Dealer.NewGame()

	//Game.HasWinner = false
	theGame.State = NewHand
}

// dealNewHand
// Trigger States: NewHand
// Resulting States: NewHandDealt
func dealNewHand() {
	// Dealer deals 2 cards to each player
	for i := 0; i < 2; i++ {
		for _, player := range theGame.Players {
			player.Hand.Takes(theGame.Dealer.Deck.Cards.GiveCard(1))
		}
		theGame.Dealer.Player.Hand.Takes(theGame.Dealer.Deck.Cards.GiveCard(1))
		// The first dealer's card is dealt face-down
		if i == 0 {
			(*theGame.Dealer.Player.Hand.TheCards)[0].FaceDown = true
		}
	}
	theGame.State = NewHandDealt
}

// dealRound
// internal function called by dealRounds - would like to flatten these!
// contains call to theGameShowCards()
// Resulting States: DealtARound, DetermineResults, GameHasWinner, DealerGoesBust, PlayerGoesBust
func dealRound() {
	theGame.ShowCards()

	// If the dealer is showing 21 or more, then they must have busted - all other players win
	if theGame.State != GameHasWinner && theGame.Dealer.Player.Scores()[player.MIN] >= 21 {
		theGame.State = DealerGoesBust
		theGame.Results += fmt.Sprintf("The Dealer is showing %d - the Dealer Goes Bust!\n",
			theGame.Dealer.Player.Scores()[player.MIN])
		for _, p := range theGame.NotBustedPlayers() {
			theGame.Scores[p] += 1
		}
	}

	if theGame.State != GameHasWinner && theGame.Dealer.GoesBust() {
		theGame.Results = "Dealer Goes Bust!"
		for _, player := range theGame.Players {
			theGame.Scores[player] += 1
		}
		theGame.State = DealerGoesBust
	}

	if theGame.State == DealerGoesBust {
		return
	}

	// if Player goes bust, the dealer wins
	for _, p := range theGame.Players {
		if theGame.State != GameHasWinner && p.GoesBust() {
			theGame.Results += fmt.Sprintf("%s Goes Bust!\n", p.Name)
			theGame.State = PlayerGoesBust
		}
		if theGame.State == GameHasWinner {
			theGame.Scores[theGame.Dealer.Player] += 1
		}
	}

	if theGame.State == PlayerGoesBust {
		return
	}
	// Player gets to choose next move first
	for _, p := range theGame.NotBustedPlayers() {
		choiceMade := false
		for !choiceMade && p.Choice == player.HIT {
			fmt.Printf("%s: (H)it or (S)tay? ", p.Name)
			var choice string
			fmt.Scanln(&choice)
			choiceMade = p.MakeChoice(choice)
		}
		if p.Choice == player.HIT {
			p.Hand.Takes(theGame.Dealer.Deck.Cards.GiveCard(1))
			theGame.ShowCards()
		}
	}

	// Dealer decides to hit or stay
	if theGame.Dealer.Player.Scores()[player.MAX] < 17 {
		fmt.Printf("Dealer takes a card.")
		theGame.Dealer.Player.Hand.Takes(theGame.Dealer.Deck.Cards.GiveCard(1))
		time.Sleep(3 * time.Second)
	} else {
		fmt.Printf("Dealer Shows %d - cannot hit over 17!\n", theGame.Dealer.Player.Scores()[player.MAX])
		theGame.Dealer.Player.Choice = player.STAY
		time.Sleep(3 * time.Second)
	}
	theGame.State = DealtARound
}

// deternubeUfAllPlayersStay
// Triggering State:
// ResultingStates: AllPlayersStay, DealARound
func determineIfAllPlayersStay() {
	if theGame.AllStay() {
		theGame.State = AllPlayersStay
		fmt.Println("All Players Have Chosen to Stay")
		time.Sleep(3 * time.Second)
	} else {
		theGame.State = DealARound
	}
}

// determineResults
// contains call to theGameShowCards()
// Triggering States: PlayerGoesBust, DealerGoesBust, DetermineResults, AllPlayersStay
// Resulting States: HandIsOver
func determineHandResults() {
	// If the players have all busted, the dealer does not have to show card!
	if !theGame.AllPlayersBusted() {
		// If the dealer is showing a hard 21 without revealing, then dealer has busted
		if theGame.Dealer.GoesBust() {
			for _, p := range theGame.Players {
				theGame.Scores[p] += 1
			}
		} else {
			theGame.Dealer.RevealFirstCard()
			theGame.ShowCards()
		}
	}
	// Evaluate final scores
	if theGame.AllStay() {
		if theGame.Dealer.GoesBust() {
			theGame.Results += "Dealer Goes Bust!\n"
			for _, p := range theGame.Players {
				theGame.Scores[p] += 1
			}
		} else {
			for _, p := range theGame.Players {
				if theGame.Dealer.Player.Scores()[player.MAX] >= p.Scores()[player.MAX] {
					theGame.Scores[theGame.Dealer.Player] += 1
					theGame.Results += fmt.Sprintf("Dealer has %d, %s has %d.  Dealer Wins!\n",
						theGame.Dealer.Player.Scores()[player.MAX],
						p.Name,
						p.Scores()[player.MAX])
				} else {
					theGame.Scores[p] += 1
					theGame.Results += fmt.Sprintf("Dealer has %d, %s has %d.  %s Wins!\n",
						theGame.Dealer.Player.Scores()[player.MAX],
						p.Name,
						p.Scores()[player.MAX],
						p.Name)
				}
			}
		}
	}
	// show final game result
	fmt.Println(theGame.Results)
	// show total game scores:
	fmt.Printf("PLAYER:\t\tSCORE:\n")
	for k, v := range theGame.Scores {
		fmt.Printf("%s:\t\t%d\n", k.Name, v)
	}
	theGame.State = HandIsOver
}

// playAgain
// determines if players want to continue playing
// Triggering States:
// Resulting States: PalyerWantsToPlayAgain, GameOver
func playAgain() {
	for _, p := range theGame.Players {
		choice := "N"
		choiceMade := false
		for !choiceMade {
			fmt.Printf("\n%s: Play Again? (Y/N) ", p.Name)
			fmt.Scanln(&choice)
			choice = strings.ToUpper(string(choice[0]))
			switch choice {
			case "Y":
				choiceMade = true
				theGame.State = PlayerWantsToPlayAgain
			case "N":
				choiceMade = true
				theGame.State = PlayerWantsToQuit
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
	for _, p := range theGame.Players {
		fmt.Printf("\nGoodbye, %s!  Play GoJack again soon!\n\n", p.Name)
	}
	theGame.State = GameOver
}
