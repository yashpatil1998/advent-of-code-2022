package day4

import (
	"strconv"
	"strings"
)

func Day4Logic() (int, int) {
	return day4LogicPart1(), day4LogicPart2()
}

func day4LogicPart1() int {
	totalCount := 0
	inputs := strings.Split(input4, "\n")
	for _, assignmentPairs := range inputs {
		assignments := strings.Split(assignmentPairs, ",")
		asmt12 := strings.Split(assignments[0], "-")
		asmt34 := strings.Split(assignments[1], "-")
		min1, _ := strconv.Atoi(asmt12[0])
		max1, _ := strconv.Atoi(asmt12[1])
		min2, _ := strconv.Atoi(asmt34[0])
		max2, _ := strconv.Atoi(asmt34[1])
		if (min1 <= min2 && max1 >= max2) || (min2 <= min1 && max2 >= max1) {
			totalCount++
		}
	}
	return totalCount
}

func day4LogicPart2() int {
	totalCount := 0
	inputs := strings.Split(input4, "\n")
	for _, assignmentPairs := range inputs {
		assignments := strings.Split(assignmentPairs, ",")
		asmt12 := strings.Split(assignments[0], "-")
		asmt34 := strings.Split(assignments[1], "-")
		min1, _ := strconv.Atoi(asmt12[0])
		max1, _ := strconv.Atoi(asmt12[1])
		min2, _ := strconv.Atoi(asmt34[0])
		max2, _ := strconv.Atoi(asmt34[1])
		if max1 >= min2 && max2 >= min1 {
			totalCount++
		}
	}
	return totalCount
}
