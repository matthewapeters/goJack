Feature: a card has a Suit and a Face Value

Scenario Outline: A card will print its face value and Suit
Given a card with Suit of "<Suit>" and a face value of "<FACEVALUE>"
Then its string value will be "<PROPERDISPLAY>""

Examples:
    |    Suit | FACEVALUE |   PROPERDISPLAY |
    |   Hearts |      Ace |             ♥ A |
    |   Spades |    Three |             ♠ 3 |
    |    Clubs |     Four |             ♣ 4 |
    | Diamonds |     Ten  |             ♦10 |

