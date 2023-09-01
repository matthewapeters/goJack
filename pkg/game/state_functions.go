package game

import (
	"fmt"
	"strings"
	"time"

	"github.com/matthewapeters/gojack/pkg/player"
)

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
				Game.State = GameOver
			default:
				fmt.Println("Dont be goof!")
				choiceMade = false
			}
		}
	}
}
func sayGoodbye() {
	for _, p := range Game.Players {
		fmt.Printf("\nGoodbye, %s!  Play GoJack again soon!\n\n", p.Name)

	}
}

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
			ShowCards()
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

func dealRounds() {
	for Game.State != DetermineResults && Game.State != AllPlayersStay {
		dealRound()
		// if both players have chosen to stay, rounds are over, determine player with highest score
		determineIfAllPlayersStay()
	}
}
func dealRound() {
	ShowCards()

	// If the dealer is showing 21 or more, then they must have busted - all other players win
	if !Game.HasWinner && Game.Dealer.Player.Scores()[player.MIN] >= 21 {
		Game.HasWinner = true
		Game.State = DetermineResults
		Game.Results += fmt.Sprintf("The Dealer is showing %d - the Dealer Goes Bust!\n",
			Game.Dealer.Player.Scores()[player.MIN])
		for _, p := range Game.NotBustedPlayers() {
			Game.Scores[p] += 1
		}
	}

	if !Game.HasWinner && Game.Dealer.GoesBust() {
		Game.Results = "Dealer Goes Bust!"
		for _, player := range Game.Players {
			Game.Scores[player] += 1
		}
		Game.State = DetermineResults
		Game.HasWinner = true
	}

	if Game.HasWinner {
		return
	}

	// if Player goes bust, the dealer wins
	for _, p := range Game.Players {
		if !Game.HasWinner && p.GoesBust() {
			Game.Results += fmt.Sprintf("%s Goes Bust!\n", p.Name)
			Game.HasWinner = true
			Game.State = DetermineResults
		}
		if Game.HasWinner {
			Game.Scores[Game.Dealer.Player] += 1
		}
	}

	if Game.HasWinner {
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
			ShowCards()
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
}
