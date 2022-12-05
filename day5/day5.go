package day5

var ogOrg = map[int][]string{
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

func shift(src int, dst int, ogOrg map[int][]string) map[int][]string {
	item := ogOrg[src][len(ogOrg[src])-1]
	ogOrg[src] = ogOrg[src][:len(ogOrg[src])-2]
	ogOrg[dst] = append(ogOrg[dst], item)
	return ogOrg
}

func Day5Logic() (string, string) {
	return day5LogicPart1(), day5LogicPart2()
}

func day5LogicPart1() string {
	return ""
}

func day5LogicPart2() string {
	return ""
}
