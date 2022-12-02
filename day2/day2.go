package day2

import (
	"strings"
)

var WinScoreMap = map[string]int{
	"W": 6,
	"D": 3,
	"L": 0,
}

var RPSMap = map[string]string{
	"A": "ROCK",
	"B": "PAPER",
	"C": "SCISSOR",
	"X": "ROCK",
	"Y": "PAPER",
	"Z": "SCISSOR",
}

var ActionScoreMap = map[string]int{
	"ROCK":    1,
	"PAPER":   2,
	"SCISSOR": 3,
}

func Day2LogicPart1() int64 {
	gamesList := strings.Split(input2, "\n")
	totalGameScore := int64(0)
	for _, game := range gamesList {
		switch game {
		case "A X":
			totalGameScore += 1 + 3
		case "A Y":
			totalGameScore += 2 + 6
		case "A Z":
			totalGameScore += 3 + 0
		case "B X":
			totalGameScore += 1 + 0
		case "B Y":
			totalGameScore += 2 + 3
		case "B Z":
			totalGameScore += 3 + 6
		case "C X":
			totalGameScore += 1 + 6
		case "C Y":
			totalGameScore += 2 + 0
		case "C Z":
			totalGameScore += 3 + 3
		}
	}
	return totalGameScore
}

func Day2LogicPart2() int64 {
	gamesList := strings.Split(input2, "\n")
	totalGameScore := int64(0)
	for _, game := range gamesList {
		switch game {
		case "A X":
			totalGameScore += 3 + 0
		case "A Y":
			totalGameScore += 1 + 3
		case "A Z":
			totalGameScore += 2 + 6
		case "B X":
			totalGameScore += 1 + 0
		case "B Y":
			totalGameScore += 2 + 3
		case "B Z":
			totalGameScore += 3 + 6
		case "C X":
			totalGameScore += 2 + 0
		case "C Y":
			totalGameScore += 3 + 3
		case "C Z":
			totalGameScore += 1 + 6
		}
	}
	return totalGameScore
}
