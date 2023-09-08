Feature: Dealer Deals Cards To All Players, and One Card Face-Down to The Dealer

Scenario Outline: Game with One Player
Given a new game of goJack
And a player list of "<Players>"
When the Dealer deals a new hand
Then the deck will have "<RemainingCards>" in the deck
And the Dealer will have 1 card facing down
And all of the player cards will be facing up

Examples:
| Players       |       RemainingCards  |
| Matt          |                    48 |
| Matt Chris    |                    46 |
| Matt Chris v  |                    44 |

Scenario Outline: Dealer Hit Or Stay
Given a new game of goJack
And the Dealer has a hand with the cards "<DealersCards>"
And the game State is "<StartingGameState>"
When the Dealer must decide to hit or Stay
Then the resulting game state will be "<ResultingGameState>"

Examples:
| DealersCards              | StartingGameState   | ResultingGameState  |
| 2s                        | DealToDealer        | DealToDealer        |
| 2s,10S                    | DealToDealer        | DealToDealer        |
| 10S                       | DealToDealer        | DealToDealer        |
| 10S,10S,2S                | DealToDealer        | DealerGoesBust      |
| 10S,10S                   | DealToDealer        | DealtARound         |
| AH,10S                    | DealToDealer        | DealtARound         |