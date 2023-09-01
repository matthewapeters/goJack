Feature: a deck of 52 playing cards

Scenario: All Cards Accounted For
Given a deck of cards
Then there will be 4 cards of each  "<FaceValue>"
 | FaceValue   |
 | Ace         |
 | King        |
 | Queen       |
 | Jack        |
 | 10          |
 | 9           |
 | 8           |
 | 7           |
 | 6           |
 | 5           |
 | 4           |
 | 3           |
 | 2           |
    And each card in each value will be a unique Suit of "<Suit>"
    | Suit |
    | Hearts |
    | Diamonds|
    | Clubs |
    | Spades |


Scenario: All Cards Present in New Deck
Given a deck of cards
Then its string values will be
"""
Deck:
    ♠ A
    ♠ K
    ♠ Q
    ♠ J
    ♠10
    ♠ 9
    ♠ 8
    ♠ 7
    ♠ 6
    ♠ 5
    ♠ 4
    ♠ 3
    ♠ 2
    ♥ A
    ♥ K
    ♥ Q
    ♥ J
    ♥10
    ♥ 9
    ♥ 8
    ♥ 7
    ♥ 6
    ♥ 5
    ♥ 4
    ♥ 3
    ♥ 2
    ♦ A
    ♦ K
    ♦ Q
    ♦ J
    ♦10
    ♦ 9
    ♦ 8
    ♦ 7
    ♦ 6
    ♦ 5
    ♦ 4
    ♦ 3
    ♦ 2
    ♣ A
    ♣ K
    ♣ Q
    ♣ J
    ♣10
    ♣ 9
    ♣ 8
    ♣ 7
    ♣ 6
    ♣ 5
    ♣ 4
    ♣ 3
    ♣ 2

"""
