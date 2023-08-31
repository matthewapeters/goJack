Feature: A Hand Has Values

Scenario Outline: A Hand Can Have More Than One Value
Given a hand with cards "<Hand>"
Then it may have these "<Values>" below or equal to 21

Examples:
|    Hand     |        Values |
|    2H,10S   |            12 |
|    AS       |          1,11 |
|    AS,AC    |       2,12,12 |
|    QS,AS    |         11,21 |
|    2H,AS,QS |            13 |
|    QS,AS,2H |            13 |
|    AS,QS,2H |            13 |
| 2H,3S,4D,5C |            14 |
|    AD,AS,AC |    3,13,13,13 |
| AH,AD,AS,AC | 4,14,14,14,14 |
| QS,10S,AS   |            21 |
| QS,10S,3S   |               |