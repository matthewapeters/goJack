package game

import (
	"fmt"
	"strings"
	"time"

	"github.com/matthewapeters/gojack/pkg/player"
)

// initializeGame
// creates new players
// Trigger States: NewGame
// Resulting States: Initialized
func initializeGame() {
	for _, name := range theGame.Names {
		theGame.Players = append(theGame.Players, player.NewPlayer(name))
	}
	theGame.State = Initialized
}

// startNewHand -
// Gathers used cards from all players and prepares players and Dealer for a new game
// Trigger States: Initialized, PlayerWantsToPlayAgain
// Resulting States: NewHand
func startNewHand() {
	theGame.Results = ""
	theGame.CurrentPlayerID = 0
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

// dealToPlayer
// Trigger States: DealARound, DealToDealer, NewHandDealt
// Resulting States: DealARound, PlayerGoesBust
func dealToPlayer() {
	p := theGame.Players[theGame.CurrentPlayerID]
	choiceMade := false

	theGame.ShowCards()

	if p.GoesBust() {
		// Player went bust on last card, play goes to next player or the Dealer
		fmt.Printf("%s Goes Bust!\n", p.Name)
		theGame.CurrentPlayerID += 1
		if theGame.CurrentPlayerID == len(theGame.Players) {
			theGame.State = PlayerGoesBust
		}
		time.Sleep(2 * time.Second)
		return
	}

	// What does the player want to do?
	for !choiceMade && p.Choice == player.HIT {
		fmt.Printf("%s: (H)it or (S)tay? ", p.Name)
		var choice string
		fmt.Scanln(&choice)
		choiceMade = p.MakeChoice(choice)
	}

	if p.Choice == player.HIT {
		p.Hand.Takes(theGame.Dealer.Deck.Cards.GiveCard(1))
	} else {
		// Player choses to Stay, play goes to next player or the Dealer
		theGame.CurrentPlayerID += 1
		if theGame.CurrentPlayerID == len(theGame.Players) {
			theGame.State = DealToDealer
		}
	}
}

// dealToDealer
// Trigger States: DealToDealer
// Resulting States: DealToDealer, DealtARound, DealerGoesBust
func dealToDealer() {
	theGame.ShowCards()

	if theGame.Dealer.GoesBust() {
		theGame.Results += fmt.Sprintf("the Dealer Goes Bust!\n")
		theGame.State = DealerGoesBust
		return
	}
	if theGame.Dealer.Player.Scores()[player.MAX] < 7 {
		fmt.Printf("Dealer takes a card.  ")
		theGame.Dealer.Player.Hand.Takes(theGame.Dealer.Deck.Cards.GiveCard(1))
		time.Sleep(2 * time.Second)
	} else {
		fmt.Printf("Dealer Shows %d - will not hit over presumed 17!\n", theGame.Dealer.Player.Scores()[player.MAX])
		theGame.Dealer.Player.Choice = player.STAY
		time.Sleep(2 * time.Second)
		theGame.State = DealtARound
	}
}

// deternubeIfAllPlayersStay
// Trigger States: DealtARound
// Resulting States: AllPlayersStay
func determineIfAllPlayersStay() {
	if theGame.AllStay() {
		theGame.State = AllPlayersStay
		fmt.Println("All Players Have Chosen to Stay")
		time.Sleep(3 * time.Second)
	}
}

// determineIfAllPlayersBusted
// after dealing a round, some of the players may have busted.
// Determine if they have all busted (which means the dealer wins and has nothing to prove)
// Trigger States: PlayerGoesBust
// Resulting States: DealARound, AllPlayersGoBust
func determineIfAllPlayersBusted() {
	if theGame.AllPlayersBusted() {
		theGame.State = AllPlayersGoBust
	} else {
		// this occurs only for multi-player future variant
		theGame.State = DealARound
	}
}

// dealerRevealsCard
// Trigger States: AllPlayersStay
// ResultingStates: DetermineResults, DealerGoesBust
func dealerRevealsCard() {
	theGame.Dealer.RevealFirstCard()
	theGame.ShowCards()
	if theGame.Dealer.GoesBust() {
		theGame.Results += "Dealer Goes Bust!\n"
		theGame.State = DealerGoesBust
		return
	}
	theGame.State = DetermineResults
}

// dealerGoesBust
// Trigger States: DealerGoesBust
// Resulting States: HandIsOver
func dealerGoesBust() {
	for _, p := range theGame.NotBustedPlayers() {
		theGame.Scores[p] += 1
	}
	theGame.State = HandIsOver
}

// dealerWins
// Trigger States: AllPlayersGoBust
// Resulting States: HandIsOver
func dealerWins() {
	theGame.Results += "Dealer Wins! \n"
	theGame.Scores[theGame.Dealer.Player] += 1
	theGame.State = HandIsOver
}

// determineResults
// contains call to theGameShowCards()
// Trigger States: DetermineResults
// Resulting States: HandIsOver
func determineHandResults() {
	// Evaluate final scores
	for _, p := range theGame.Players {
		if theGame.Dealer.Player.Scores()[player.MAX] >= p.Scores()[player.MAX] {
			theGame.Scores[theGame.Dealer.Player] += 1
			if p.Scores()[player.MAX] == 0 {
				theGame.Results += fmt.Sprintf("%s Goes Bust!  Dealer Wins.\n", p.Name)
			} else {
				theGame.Results += fmt.Sprintf("Dealer has %d, %s has %d.  Dealer Wins!\n",
					theGame.Dealer.Player.Scores()[player.MAX],
					p.Name,
					p.Scores()[player.MAX])
			}
		} else {
			theGame.Scores[p] += 1
			theGame.Results += fmt.Sprintf("Dealer has %d, %s has %d.  %s Wins!\n",
				theGame.Dealer.Player.Scores()[player.MAX],
				p.Name,
				p.Scores()[player.MAX],
				p.Name)
		}
	}
	theGame.State = HandIsOver
}

// playAgain
// display hand and game stats, determines if players want to continue playing
// Trigger States: HandIsOver
// Resulting States: PlayerWantsToPlayAgain, GameOver
func playAgain() {
	// show final game result
	fmt.Println(theGame.Results)
	// show total game scores:
	fmt.Printf("PLAYER:\t\tSCORE:\n")
	for k, v := range theGame.Scores {
		fmt.Printf("%s:\t\t%d\n", k.Name, v)
	}
	names := []string{}

	for _, p := range theGame.Players {
		names = append(names, p.Name)
	}
	nameList := strings.Join(names, " and ")
	choice := "N"
	choiceMade := false
	for !choiceMade {
		fmt.Printf("\n%s: Play Again? (Y/N) ", nameList)
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

// sayGoodbye
// goodbye message to players
// Triggereing States: GameOver
func sayGoodbye() {
	names := []string{}

	for _, p := range theGame.Players {
		names = append(names, p.Name)
	}
	nameList := strings.Join(names, " and ")

	fmt.Printf("\nGoodbye, %s!  Play GoJack again soon!\n\n", nameList)
	theGame.State = GameOver
}
