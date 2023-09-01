package game

import (
	"fmt"
	"strings"
	"time"

	"github.com/matthewapeters/gojack/pkg/dealer"
	"github.com/matthewapeters/gojack/pkg/player"
)

var (
	Dealer    = dealer.NewDealer()
	PlayerOne = player.NewPlayer("PlayerOne")
	scores    = map[*player.Player]int{}
)

func ShowCards() {
	ClearScreen()

	//Dealer shows Hand
	fmt.Printf("Dealer's %s\n", Dealer.Player.Hand)

	fmt.Printf(strings.Repeat("\n", 4))

	//Show PlayerOne's Hand
	fmt.Printf("%s's %s\n", PlayerOne.Name, PlayerOne.Hand)
	fmt.Printf(strings.Repeat("\n", 4))
}

func Play(n string) {
	PlayerOne.Name = n
	Rounds := 0
	quit := false
	gameResult := ""
	ClearScreen()
	for quit != true {
		Rounds += 1
		winner := false
		// Dealer Shuffles the Deck
		Dealer.Shuffles()
		// Dealer deals 2 cards to each player
		for i := 0; i < 2; i++ {
			PlayerOne.Hand.Takes(Dealer.Deck.Cards.GiveCard(1))
			Dealer.Player.Hand.Takes(Dealer.Deck.Cards.GiveCard(1))
			// The first dealer's card is dealt face-down
			if i == 0 {
				(*Dealer.Player.Hand.TheCards)[0].FaceDown = true
			}
		}

		// Begin Play Loop
		for !winner {
			ShowCards()

			// If Dealer Hits 21, Dealer Wins
			if !winner && Dealer.Player.HitsTwentyOne() {
				gameResult = "Dealer Hits 21!"
				scores[Dealer.Player] += 1
				winner = true
			}
			if !winner && Dealer.Player.GoesBust() {
				gameResult = "Dealer Goes Bust!"
				scores[PlayerOne] += 1
				winner = true
			}

			if winner {
				break
			}

			// if Player has 21, player wins (unless dealer has already won)
			if !winner && PlayerOne.HitsTwentyOne() {
				gameResult = fmt.Sprintf("%s Hits 21!\n", PlayerOne.Name)
				scores[PlayerOne] += 1
				winner = true
			}

			// if Player goes bust, the dealer wins
			if !winner && PlayerOne.GoesBust() {
				gameResult = fmt.Sprintf("%s Goes Bust!\n", PlayerOne.Name)
				scores[Dealer.Player] += 1
				winner = true
			}

			if winner {
				break
			}
			// Player gets to choose next move first
			choiceMade := false
			for !choiceMade && PlayerOne.Choice == player.HIT {
				fmt.Printf("%s: (H)it or (S)tay? ", PlayerOne.Name)
				var choice string
				fmt.Scanln(&choice)
				choiceMade = PlayerOne.MakeChoice(choice)
			}
			if PlayerOne.Choice == player.HIT {
				PlayerOne.Hand.Takes(Dealer.Deck.Cards.GiveCard(1))
				ShowCards()
			}

			// Dealer decides to hit or stay
			if Dealer.Player.Scores()[player.MAX] < 17 {
				Dealer.Player.Choice = player.HIT
				fmt.Printf("Dealer takes a card.")
				Dealer.Player.Hand.Takes(Dealer.Deck.Cards.GiveCard(1))
			} else {
				fmt.Printf("Dealer Shows %d - cannot hit over 17!\n", Dealer.Player.Scores()[player.MAX])
				Dealer.Player.Choice = player.STAY
			}
			time.Sleep(3 * time.Second)

			// if both players have chosen to stay, rounds are over, determine player with highest score
			if Dealer.Player.Choice == player.STAY && PlayerOne.Choice == player.STAY {
				break
			}
		}
		Dealer.RevealFirstCard()
		ShowCards()
		// Evaluate final scores
		if Dealer.Player.Choice == player.STAY && PlayerOne.Choice == player.STAY {
			if Dealer.Player.Scores()[player.MAX] >= PlayerOne.Scores()[player.MAX] {
				scores[Dealer.Player] += 1
				gameResult = fmt.Sprintf("Dealer has %d, %s has %d.  Dealer Wins!\n",
					Dealer.Player.Scores()[player.MAX],
					PlayerOne.Name,
					PlayerOne.Scores()[player.MAX])
			} else {
				scores[PlayerOne] += 1
				gameResult = fmt.Sprintf("Dealer has %d, %s has %d.  %s Wins!\n",
					Dealer.Player.Scores()[player.MAX],
					PlayerOne.Name,
					PlayerOne.Scores()[player.MAX],
					PlayerOne.Name)
			}
		}
		// show final game result
		fmt.Println(gameResult)
		// show total game scores:
		fmt.Printf("PLAYER:\t\tSCORE:\n")
		for k, v := range scores {
			fmt.Printf("%s:\t\t%d\n", k.Name, v)
		}

		choice := "N"
		choiceMade := false
		for !choiceMade {
			fmt.Printf("\n%s: Play Again? (Y/N) ", PlayerOne.Name)
			fmt.Scanln(&choice)
			choice = strings.ToUpper(string(choice[0]))
			switch choice {
			case "Y":
				choiceMade = true
				Dealer.NewGame()
				PlayerOne.NewGame()
				ClearScreen()
			case "N":
				choiceMade = true
				quit = true
			default:
				fmt.Println("Dont be goof!")
				choiceMade = false
			}
		}
	}
	fmt.Printf("\nGoodbye, %s!  Play GoJack again soon!\n\n", PlayerOne.Name)
}

func ClearScreen() {
	// Actually, just CR until it should be blank
	fmt.Print(strings.Repeat("\n", 50))
}

func TitleArt() {
	// Tittle art generated at:
	// https://patorjk.com/software/taag
	// using
	// doh.flf by Curtis Wanner (cwanner@acs.bu.edu)
	// latest revision - 4/95
	art := `
                                                    JJJJJJJJJJJ          AAA                  CCCCCCCCCCCCCKKKKKKKKK    KKKKKKK
                                                    J:::::::::J         A:::A              CCC::::::::::::CK:::::::K    K:::::K
                                                    J:::::::::J        A:::::A           CC:::::::::::::::CK:::::::K    K:::::K
                                                    JJ:::::::JJ       A:::::::A         C:::::CCCCCCCC::::CK:::::::K   K::::::K
   ggggggggg   ggggg   ooooooooooo                    J:::::J        A:::::::::A       C:::::C       CCCCCCKK::::::K  K:::::KKK
  g:::::::::ggg::::g oo:::::::::::oo                  J:::::J       A:::::A:::::A     C:::::C                K:::::K K:::::K
 g:::::::::::::::::go:::::::::::::::o                 J:::::J      A:::::A A:::::A    C:::::C                K::::::K:::::K
g::::::ggggg::::::ggo:::::ooooo:::::o                 J:::::j     A:::::A   A:::::A   C:::::C                K:::::::::::K
g:::::g     g:::::g o::::o     o::::o                 J:::::J    A:::::A     A:::::A  C:::::C                K:::::::::::K
g:::::g     g:::::g o::::o     o::::o     JJJJJJJ     J:::::J   A:::::AAAAAAAAA:::::A C:::::C                K::::::K:::::K
g:::::g     g:::::g o::::o     o::::o     J:::::J     J:::::J  A:::::::::::::::::::::AC:::::C                K:::::K K:::::K
g::::::g    g:::::g o::::o     o::::o     J::::::J   J::::::J A:::::AAAAAAAAAAAAA:::::AC:::::C       CCCCCCKK::::::K  K:::::KKK
g:::::::ggggg:::::g o:::::ooooo:::::o     J:::::::JJJ:::::::JA:::::A             A:::::AC:::::CCCCCCCC::::CK:::::::K   K::::::K
 g::::::::::::::::g o:::::::::::::::o      JJ:::::::::::::JJA:::::A               A:::::ACC:::::::::::::::CK:::::::K    K:::::K
  gg::::::::::::::g  oo:::::::::::oo         JJ:::::::::JJ A:::::A                 A:::::A CCC::::::::::::CK:::::::K    K:::::K
    gggggggg::::::g    ooooooooooo             JJJJJJJJJ  AAAAAAA                   AAAAAAA   CCCCCCCCCCCCCKKKKKKKKK    KKKKKKK
            g:::::g
gggggg      g:::::g
g:::::gg   gg:::::g
 g::::::ggg:::::::g
  gg:::::::::::::g
    ggg::::::ggg
       gggggg

`
	fmt.Println(art)
}
