package day2

import (
	"strings"
)

func day2LogicPart1() int64 {
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

func day2LogicPart2() int64 {
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

func Day2Logic() (int64, int64) {
	return day2LogicPart1(), day2LogicPart2()
}
