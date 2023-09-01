package game

import (
	"fmt"
	"strings"
	"time"

	"github.com/matthewapeters/gojack/pkg/dealer"
	"github.com/matthewapeters/gojack/pkg/player"
)

var (
	Game = &game{
		Dealer:  dealer.NewDealer(),
		Players: player.Players{player.NewPlayer("PlayerOne")},
		Scores:  map[*player.Player]int{},
	}
)

type game struct {
	*dealer.Dealer
	player.Players
	Scores    map[*player.Player]int
	Quit      bool
	HasWinner bool
	Results   string
}

func (g *game) AllStay() bool {
	allStay := false

	if g.Dealer.Player.Choice == player.STAY {
		allStay = true
		for _, p := range g.NotBustedPlayers() {
			allStay = allStay && p.Choice == player.STAY
		}
	}
	return allStay
}

func (g *game) AllPlayersBusted() (allBusted bool) {
	allBusted = true
	for _, p := range Game.Players {
		allBusted = allBusted && p.GoesBust()
	}
	return
}

func (g *game) NotBustedPlayers() (players []*player.Player) {
	players = []*player.Player{}
	for _, p := range g.Players {
		if !p.GoesBust() {
			players = append(players, p)
		}
	}
	return
}

func (g *game) NewHand() {
	g.Results = ""
	// Gather all of the cards from the players
	for _, p := range g.Players {
		Game.Dealer.GatherPlayedCards(*p.Hand.TheCards)
		p.NewGame()
	}

	// Dealer initiates new game by returning played cards to deck and shuffling the deck, if necessary
	g.Dealer.NewGame()

	g.HasWinner = false

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
}

func ShowCards() {
	ClearScreen()

	//Dealer shows Hand
	fmt.Printf("Dealer's %s%s", Game.Dealer.Player.Hand, strings.Repeat("\n", 5))

	//Show PlayerOne's Hand
	for _, player := range Game.Players {
		fmt.Printf("%s's %s%s", player.Name, player.Hand, strings.Repeat("\n", 5))
	}
}

func Play(n ...string) {
	for i, name := range n {
		if Game.Players[i] == nil {
			Game.Players[i] = player.NewPlayer(name)
		} else {
			Game.Players[i].Name = name
		}
	}
	Rounds := 0
	ClearScreen()
	for Game.Quit != true {
		Game.NewHand()
		Rounds += 1

		// Begin Play Loop
		for !Game.HasWinner {
			ShowCards()
			time.Sleep(2 * time.Second)

			// If the dealer is showing 21 or more, then they must have busted - all other players win
			if !Game.HasWinner && Game.Dealer.Player.Scores()[player.MIN] >= 21 {
				Game.HasWinner = true
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
				Game.HasWinner = true
			}

			if Game.HasWinner {
				break
			}

			// if Player goes bust, the dealer wins
			for _, p := range Game.Players {
				if !Game.HasWinner && p.GoesBust() {
					Game.Results += fmt.Sprintf("%s Goes Bust!\n", p.Name)
					Game.HasWinner = true
				}
				if Game.HasWinner {
					Game.Scores[Game.Dealer.Player] += 1
				}
			}

			if Game.HasWinner {
				break
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

			// if both players have chosen to stay, rounds are over, determine player with highest score

			if Game.AllStay() {
				fmt.Println("All Players Have Chosen to Stay - Dealer Reveals First Card")
				time.Sleep(3 * time.Second)
				break
			}
		}

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
				case "N":
					choiceMade = true
					Game.Quit = true
				default:
					fmt.Println("Dont be goof!")
					choiceMade = false
				}
			}
		}
	}
	for _, p := range Game.Players {
		fmt.Printf("\nGoodbye, %s!  Play GoJack again soon!\n\n", p.Name)

	}
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
