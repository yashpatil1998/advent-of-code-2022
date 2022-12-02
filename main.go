package main

import (
	"AdventOfCode2022/day1"
	"AdventOfCode2022/day2"
	"log"
)

func main() {
	day1p1, day1p2 := day1.Day1Logic()
	log.Printf("day1 part 1 %v part2 %v", day1p1, day1p2)
	day2p1, day2p2 := day2.Day2Logic()
	log.Printf("day2 part 1 %v part2 %v", day2p1, day2p2)
}
