# goJack #

```bash
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

```

Copyright (c) Matthew Peters, 2023

Additional Copyrights and Credits:

_Tittle art generated at [https://patorjk.com/software/taag](https://patorjk.com/software/taag)_
_using font_
doh.flf _by Curtis Wanner (cwanner@acs.bu.edu)  (latest revision - 4/95)_

_All copyrights for patorjk.com and Curtis Wanner are held by respective authors or agents._

## Building goJack ##

GoJack is written in *Go version go1.21.0* and runs as a command-line application.

From the project directory, you can build goJack locally with

```bash
go build -o ./ ./...
```

## Playing goJack ##

goJack is the classic game of Black Jack or Twenty-One.  The game consists of a computerized dealer and you, the player.

The game starts with the dealer shuffling a standard 52 card pocker deck (Hearts, Diamonds, Clubs, Spades with values from 2-10 and Jack, Queen, King and Ace).  Number cards are worth the numbers shown.  Face cards are worth 10.  Aces are worth either 1 or 11.

The dealer deals two cards to itself and the player.  If a player has 21 points, the player wins.  Ties revert to the dealer.  If the player has more than 21 points, the player "Goes Bust" and loses the hand.  If neither player receives 21 points in a hand, the player with the highest score wins.  If neither player wins with the two initial cards, the game goes into rounds.  Rounds continue as long as either player chooses to "Hit" - that is, to receive one more card.  Once a player choses to "Stay", rounds continue but the player no longer receives cards.  The dealer will continue to hit until the dealer hand has 21, goes bust, or shows 17 or more.  When all players choose to Stay, rounds end, and the player with the highest score wins, with ties reverting to the dealer.

Suits do not affect a hand's points.

The winner of each hand is given a point, and scores are shown at the end of each hand.

## Programmer Notes ##

goJack is intended to demonstrate the use of Go and Go Modules and Behavioral Driven Development (BDD) using the Go Cucumber library "Godog."

Tests may periodically break as improvements are added to the game, such as for string representation of cards.  Each of the packages (./pkg/*) are testable using the godog command line from their respective directories.  Broken tests will be corrected when resources are available.

Packages are intended to illustrate compartmentalization of game components.

While Go does not explicitly support Object Oriented Prgramming (OOP) or class inheritence, it does support compositability, allowing for inheritence through inclusion.  Please refer to Go documentation for additional thoughts on this.

### Contributing, Forking, &c ###

Please contact me if you are interested in contributing to goJack.  GoJack is copyrighted content, although it is free to download and build, and play for individual entertainment purposes.  It may not be distributed in code or compiled form for monitary gain unless explicitly under license approved by author.  All rights are reserved, primarily because the content is ultimately intended for use as reference material in upcoming publication efforts.  Forking the repository is forbidden without express written permission from the author.  Portions or all of this repository may be released under license by the author at the author's discretion without notice in the form of other repositories.  Please contact author with any related questions.

## Contact Information ##

Author: Matthew Peters

email: [matthew@datadelve.net](matthew@datadelve.net)