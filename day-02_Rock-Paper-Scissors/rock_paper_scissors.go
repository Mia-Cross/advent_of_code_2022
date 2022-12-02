package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

/* Rules of the game :

- OutcomePts:	0 if you lost, 3 if the round was a draw, and 6 if you won
- BasePts:	1 for Rock, 2 for Paper, and 3 for Scissors
- Them:		A for Rock, B for Paper, and C for Scissors
- Us (pt1):	X for Rock, Y for Paper, and Z for Scissors
- Us (pt2):	X for lose, Y for draw, and Z for win

			X-Rock	Y-Paper	Z-Scis.     <- Us
Rock--A		[=|4]	[+|8]	[-|3]
Paper-B		[-|1]	[=|5]	[+|9]
Scis.-C		[+|7]	[-|2]	[=|6]
*/

func calculateScoreBasicRules(data string) int {
	score := 0
	outcome := map[string][]int{
		"X": {4, 1, 7},
		"Y": {8, 5, 2},
		"Z": {3, 9, 6},
	}

	for _, round := range strings.Split(data, "\n") {
		played := strings.Split(round, " ")
		us := played[1]
		them := played[0][0] - 'A'
		score += outcome[us][them]
	}
	return score
}

func calculateScoreNewRules(data string) int {
	score := 0
	us := map[string][]int{
		"A": {3, 4, 8},
		"B": {1, 5, 9},
		"C": {2, 6, 7},
	}
	for _, round := range strings.Split(data, "\n") {
		input := strings.Split(round, " ")
		them := input[0]
		outcome := input[1][0] - 'X'
		score += us[them][outcome]
	}
	return score
}

func main() {
	buffer, err := os.ReadFile("moves_list.txt")
	if err != nil {
		log.Fatalf("error parsing file: %v", err)
	}

	fmt.Printf("Total score part 1: %d\n", calculateScoreBasicRules(string(buffer)))
	fmt.Printf("Total score part 2: %d\n", calculateScoreNewRules(string(buffer)))
}
