Feature: Player May Have More Than One Hand


Scenario: Player With More Than One Hand
Given a player with "<NumberOfHands>"
Then HasMoreHands will respond with "<HasMoreHands>"
And the player's resulting hand index is "<HandIndex>"

Examples:
| NumberOfHands | HandIndex | HasMoreHands  |
| 1             | 0         | false         |
| 2             | 0         | true          |


Scenario: Player With Multiple Hands Can Access Next Hand
Given a player with "<NumberOfHands>"
And the player's current hand index is "<HandIndex>"
When the player invokes NextHand
Then the player's resulting hand index is "<NextHandIndex>"
And HasMoreHands will respond with "<HasMoreHands>"

Examples:
| NumberOfHands | HandIndex | NextHandIndex | HasMoreHands  |
| 2             | 0         | 1             | false         |
| 2             | 1         | 1             | false         |
| 1             | 0         | 0             | false         |
| 3             | 0         | 1             | true          |
| 3             | 1         | 2             | false         |
| 3             | 2         | 2             | false         |