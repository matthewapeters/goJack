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
// Resulting States: DealToPlayer
func dealNewHand() {
	// Dealer deals 2 cards to each player
	for i := 0; i < 2; i++ {
		for _, player := range theGame.Players {
			player.Hand.Takes(theGame.Dealer.Deck.Cards.GiveCard(1))
		}
		switch i {
		case 0:
			// The first dealer's card is dealt face-down
			theGame.Dealer.Player.Hand.Takes(theGame.Dealer.Deck.Cards.GiveCard(1).FacingDown())
		default:
			theGame.Dealer.Player.Hand.Takes(theGame.Dealer.Deck.Cards.GiveCard(1).FacingUp())
		}
	}
	theGame.State = DealToPlayer
}

// dealToPlayer
// Trigger States: DealToPlayer
// Resulting States: PlayerGoesBust, PromptPlayer
func dealToPlayer() {
	p := theGame.Players[theGame.CurrentPlayerID]

	theGame.ShowCards()

	if p.GoesBust() {
		// Player went bust on last card, play goes to next player or the Dealer
		if !theGame.SupressDisplay {
			fmt.Printf("%s Goes Bust!\n", p.Name)
		}
		theGame.CurrentPlayerID += 1
		if theGame.CurrentPlayerID >= len(theGame.Players) {
			theGame.State = PlayerGoesBust
		}
		time.Sleep(theGame.DisplayInterval)
		return
	}

	theGame.State = PromptPlayer
}

// playerChooses
// Trigger States: PromptPlayer
// Resulting States: PlayerTakesCard, PlayerStays
func playerChooses() {
	p := theGame.Players[theGame.CurrentPlayerID]

	choiceMade := false
	// What does the player want to do?
	for !choiceMade && p.Choice == player.HIT {
		fmt.Printf("%s: (H)it or (S)tay? ", p.Name)
		var choice string
		fmt.Scanln(&choice)
		choiceMade = p.MakeChoice(choice)
	}
	if p.Choice == player.HIT {
		theGame.State = PlayerTakesCard
	} else {
		theGame.State = PlayerStays
	}
}

// playerTakesCard
// Trigger States: PlayerTakesCard
// Resulting States: DealToPlayer
func playerTakesCard() {
	p := theGame.Players[theGame.CurrentPlayerID]
	p.Hand.Takes(theGame.Dealer.Deck.Cards.GiveCard(1))
	theGame.State = DealToPlayer
}

// playerStays
// Trigger States: PlayerStays
// Resulting States: DealToPlayer, DealToDealer
func nextPlayersTurn() {
	// Player choses to Stay, play goes to next player or the Dealer
	theGame.CurrentPlayerID += 1
	if theGame.CurrentPlayerID == len(theGame.Players) {
		theGame.CurrentPlayerID = 0
		theGame.State = DealToDealer
		return
	}
	theGame.State = DealToPlayer
}

// dealToDealer
// Trigger States: DealToDealer
// Resulting States: DealToDealer, DealtARound, DealerGoesBust
func dealToDealer() {
	theGame.Dealer.RevealFirstCard()
	theGame.ShowCards()
	if theGame.Dealer.GoesBust() {
		theGame.Results += "Dealer Goes Bust!\n"
		theGame.State = DealerGoesBust
		return
	}
	theGame.ShowCards()

	if theGame.Dealer.GoesBust() {
		theGame.Results += fmt.Sprintf("the Dealer Goes Bust!\n")
		theGame.State = DealerGoesBust
		return
	}
	if theGame.Dealer.Player.Scores()[player.MAX] < 17 {
		if !theGame.SupressDisplay {
			fmt.Printf("Dealer takes a card.  ")
		}
		theGame.Dealer.Player.Hand.Takes(theGame.Dealer.Deck.Cards.GiveCard(1))
		time.Sleep(theGame.DisplayInterval)
	} else {
		if !theGame.SupressDisplay {
			fmt.Printf("Dealer Shows %d - will not hit over 17!\n", theGame.Dealer.Player.Scores()[player.MAX])
		}
		theGame.Dealer.Player.Choice = player.STAY
		time.Sleep(theGame.DisplayInterval)
		theGame.State = DealtARound
	}
}

// determineIfAllPlayersBusted
// after dealing a round, some of the players may have busted.
// Determine if they have all busted (which means the dealer wins and has nothing to prove)
// Trigger States: PlayerGoesBust
// Resulting States: DealToPlayer, DealToDealer, AllPlayersGoBust
func determineIfAllPlayersBusted() {
	if theGame.AllPlayersBusted() {
		theGame.State = AllPlayersGoBust
	} else {
		if theGame.CurrentPlayerID == len(theGame.Players) {
			theGame.State = DealToDealer
			return
		}
		// this occurs only for multi-player future variant
		theGame.State = DealToPlayer
	}
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
// Trigger States: DealtARound
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
