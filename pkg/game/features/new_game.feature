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
