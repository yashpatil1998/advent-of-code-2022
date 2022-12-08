package main

import (
	"AdventOfCode2022/day1"
	"AdventOfCode2022/day2"
	"AdventOfCode2022/day3"
	"AdventOfCode2022/day4"
	"AdventOfCode2022/day5"
	"AdventOfCode2022/day6"
	"AdventOfCode2022/day7"
	"log"
)

func main() {
	day1p1, day1p2 := day1.Day1Logic()
	log.Printf("day1 part 1 %v part2 %v", day1p1, day1p2)
	day2p1, day2p2 := day2.Day2Logic()
	log.Printf("day2 part 1 %v part2 %v", day2p1, day2p2)
	day3p1, day3p2 := day3.Day3Logic()
	log.Printf("day3 part 1 %v part2 %v", day3p1, day3p2)
	day4p1, day4p2 := day4.Day4Logic()
	log.Printf("day4 part 1 %v part 2 %v", day4p1, day4p2)
	day5p1, day5p2 := day5.Day5Logic()
	log.Printf("day5 part 1 %v part 2 %v", day5p1, day5p2)
	day6p1, day6p2 := day6.Day6Logic()
	log.Printf("day6 part 1 %v part 2 %v", day6p1, day6p2)
	day7p1, day7p2 := day7.Day7Logic()
	log.Printf("day7 part 1 %v part 2 %v", day7p1, day7p2)
}
