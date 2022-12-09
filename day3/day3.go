package day3

import (
	"strings"
)

func calPriority(s string) int {
	ascii := s[0]
	if ascii < 97 {
		return int(ascii - 65 + 27)
	} else {
		return int(ascii - 97 + 1)
	}
}

func day3LogicPart1() int {
	inputs := strings.Split(input3, "\n")
	totalSum := 0
	for _, entry := range inputs {
		r1 := entry[:len(entry)/2]
		r2 := entry[len(entry)/2:]
		for _, c := range r1 {
			if strings.Contains(r2, string(c)) {
				totalSum += calPriority(string(c))
				break
			}
		}
	}
	return totalSum
}

func day3LogicPart2() int {
	inputs := strings.Split(input3, "\n")
	totalSum := 0
	for i := 0; i < len(inputs); {
		r1 := inputs[i]
		i++
		r2 := inputs[i]
		i++
		r3 := inputs[i]
		i++
		for _, c := range r1 {
			if strings.Contains(r2, string(c)) && strings.Contains(r3, string(c)) {
				totalSum += calPriority(string(c))
				break
			}
		}
	}
	return totalSum
}

func Day3Logic() (int, int) {
	return day3LogicPart1(), day3LogicPart2()
}
