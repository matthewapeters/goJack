package game

import (
	"fmt"
	"strings"

	"github.com/matthewapeters/gojack/pkg/dealer"
	"github.com/matthewapeters/gojack/pkg/player"
)

var (
	Dealer    = dealer.NewDealer()
	PlayerOne = player.NewPlayer("PlayerOne")
)

func Play(n string) {
	PlayerOne.Name = n
	Rounds := 0
	quit := false
	scores := map[*player.Player]int{}
	ClearScreen()
	for quit != true {
		Rounds += 1
		winner := false
		// Dealer Shuffles the Deck
		Dealer.Shuffles()
		// Dealer takes two cards
		Dealer.Player.Hand.Takes(Dealer.Deck.Cards.GiveCard(1))
		Dealer.Player.Hand.Takes(Dealer.Deck.Cards.GiveCard(1))
		// Player takes first two cards
		PlayerOne.Hand.Takes(Dealer.Deck.Cards.GiveCard(1))
		PlayerOne.Hand.Takes(Dealer.Deck.Cards.GiveCard(1))

		// Begin Play Loop
		for !winner {
			//Dealer shows Hand
			fmt.Printf("Dealer's %s\n", Dealer.Player.Hand)
			fmt.Printf("Dealer's Hand Worth: %v\n", Dealer.Player.Scores())

			// If Dealer Hits 21, Dealer Wins
			if Dealer.Player.HitsTwentyOne() {
				fmt.Println("Dealer Hits 21!")
				scores[Dealer.Player] += 1
				winner = true
			}
			if Dealer.Player.GoesBust() {
				fmt.Println("Dealer Goes Bust!")
				scores[PlayerOne] += 1
				winner = true
			}

			//Show PlayerOne's Hand
			fmt.Printf("%s's %s\n", PlayerOne.Name, PlayerOne.Hand)
			fmt.Printf("%s Hand Worth: %v\n", PlayerOne.Name, PlayerOne.Scores())

			// if Player has 21, player wins (unless dealer has already won)
			if PlayerOne.HitsTwentyOne() && !winner {
				fmt.Printf("%s Hits 21!\n", PlayerOne.Name)
				scores[PlayerOne] += 1
				winner = true
			}
			// if Player goes bust, the dealer wins
			if PlayerOne.GoesBust() {
				fmt.Printf("%s Goes Bust!\n", PlayerOne.Name)
				scores[Dealer.Player] += 1
				winner = true
			}
			if winner {
				break
			}
			choiceMade := false
			for !choiceMade && PlayerOne.Choice == player.HIT {
				fmt.Printf("%s: (H)it or (S)tay? ", PlayerOne.Name)
				var choice string
				fmt.Scanln(&choice)
				choiceMade = PlayerOne.MakeChoice(choice)
			}
			if PlayerOne.Choice == player.HIT {
				PlayerOne.Hand.Takes(Dealer.Deck.Cards.GiveCard(1))
			}

			if Dealer.Player.Scores()[player.MAX] < 17 {
				Dealer.Player.Choice = player.HIT
				Dealer.Player.Hand.Takes(Dealer.Deck.Cards.GiveCard(1))
			} else {
				fmt.Printf("Dealer Shows %d - cannot hit over 17!\n", Dealer.Player.Scores()[player.MAX])
				Dealer.Player.Choice = player.STAY
			}

			if Dealer.Player.Choice == player.STAY && PlayerOne.Choice == player.STAY {
				winner = true
				if Dealer.Player.Scores()[player.MAX] >= PlayerOne.Scores()[player.MAX] {
					scores[Dealer.Player] += 1
					fmt.Printf("Dealer Wins!\n")
				} else {
					scores[PlayerOne] += 1
					fmt.Printf("%s Wins!\n", PlayerOne.Name)
				}
			}
			// Clear Screen
			ClearScreen()
		}

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
	fmt.Print(strings.Repeat("\n", 20))
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
