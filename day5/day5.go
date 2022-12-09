package day5

import (
	"fmt"
	"strings"
)

var inputOrganisation = func() map[int][]string {
	return map[int][]string{
		1: {"N", "R", "G", "P"},
		2: {"J", "T", "B", "L", "F", "G", "D", "C"},
		3: {"M", "S", "V"},
		4: {"L", "S", "R", "C", "Z", "P"},
		5: {"P", "S", "L", "V", "C", "W", "D", "Q"},
		6: {"C", "T", "N", "W", "D", "M", "S"},
		7: {"H", "D", "G", "W", "P"},
		8: {"Z", "L", "P", "H", "S", "C", "M", "V"},
		9: {"R", "P", "F", "L", "W", "G", "Z"},
	}
}

func shiftSingle(src, dst int, ogOrg map[int][]string) map[int][]string {
	item := ogOrg[src][len(ogOrg[src])-1]
	ogOrg[src] = ogOrg[src][:len(ogOrg[src])-1]
	ogOrg[dst] = append(ogOrg[dst], item)
	return ogOrg
}

func shiftMulti(src, dst, numCrates int, ogOrg map[int][]string) map[int][]string {
	items := ogOrg[src][len(ogOrg[src])-numCrates : len(ogOrg[src])]
	ogOrg[src] = ogOrg[src][:len(ogOrg[src])-numCrates]
	ogOrg[dst] = append(ogOrg[dst], items...)
	return ogOrg
}

func parseInputLine(input string) (int, int, int) {
	var count, from, to int
	fmt.Sscanf(input, "move %d from %d to %d", &count, &from, &to)
	return count, from, to
}

func giveTopItems(ogOrg map[int][]string) string {
	keys := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ans := ""
	for _, element := range keys {
		ans += ogOrg[element][len(ogOrg[element])-1]
	}
	return ans
}

func Day5Logic() (string, string) {
	return day5LogicPart1(), day5LogicPart2()
}

func day5LogicPart1() string {
	inputs := strings.Split(input5, "\n")
	structure := inputOrganisation()
	for _, move := range inputs {
		count, from, to := parseInputLine(move)
		for i := 0; i < count; i++ {
			structure = shiftSingle(from, to, structure)
		}
	}
	return giveTopItems(structure)
}

func day5LogicPart2() string {
	inputs := strings.Split(input5, "\n")
	structure := inputOrganisation()
	for _, move := range inputs {
		count, from, to := parseInputLine(move)
		structure = shiftMulti(from, to, count, structure)
	}
	return giveTopItems(structure)
}
