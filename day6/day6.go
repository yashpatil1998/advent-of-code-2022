package day6

import (
	"sort"
	"strings"
)

func Day6Logic() (int, int) {
	return day6LogicPart1(), day6LogicPart2()
}

func sortString(input string) string {
	s := strings.Split(input, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func checkAllUnique(input string) bool {
	input = sortString(input)
	for i := 0; i < len(input)-1; i++ {
		if string(input[i]) == string(input[i+1]) {
			return false
		}
	}
	return true
}

func day6LogicPart1() int {
	input := input6
	for i := 3; i < len(input); i++ {
		subStr := input[i-3 : i+1]
		if checkAllUnique(subStr) {
			return i + 1
		}
	}
	return 0
}

func day6LogicPart2() int {
	input := input6
	for i := 13; i < len(input); i++ {
		subStr := input[i-13 : i+1]
		if checkAllUnique(subStr) {
			return i + 1
		}
	}
	return 0
}
