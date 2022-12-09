package day9

import (
	"log"
	"math"
	"strconv"
	"strings"
)

func Day9Logic() (int, int) {
	return day9Part1(), day9Part2()
}

var visitedTPart1 = map[[2]int]bool{}
var visitedTPart2 = map[[2]int]bool{}

func calcGap(head, tail [2]int) float64 {
	return math.Sqrt(math.Pow(math.Abs(float64(head[0]-tail[0])), 2) + math.Pow(math.Abs(float64(head[1]-tail[1])), 2))
}

func day9Part1() int {
	input := strings.Split(input9, "\n")
	knots := [2][2]int{{0, 0}, {0, 0}}
	head := knots[0]
	tail := knots[1]
	visitedTPart1[tail] = true
	for _, hMove := range input {
		dirAndSteps := strings.Split(hMove, " ")
		dir := dirAndSteps[0]
		steps, _ := strconv.Atoi(dirAndSteps[1])
		for i := 0; i < steps; i++ {
			switch dir {
			case "U":
				{
					head[1]++
					gap := calcGap(head, tail)
					if gap > math.Sqrt(2) {
						tail = [2]int{head[0], head[1] - 1}
					}
					visitedTPart1[tail] = true
				}
			case "R":
				{
					head[0]++
					gap := calcGap(head, tail)
					if gap > math.Sqrt(2) {
						tail = [2]int{head[0] - 1, head[1]}
					}
					visitedTPart1[tail] = true
				}
			case "D":
				{
					head[1]--
					gap := calcGap(head, tail)
					if gap > math.Sqrt(2) {
						tail = [2]int{head[0], head[1] + 1}
					}
					visitedTPart1[tail] = true
				}
			case "L":
				{
					head[0]--
					gap := calcGap(head, tail)
					if gap > math.Sqrt(2) {
						tail = [2]int{head[0] + 1, head[1]}
					}
					visitedTPart1[tail] = true
				}
			}
		}
	}
	return len(visitedTPart1)
}

func day9Part2() int {

	input := strings.Split(input9, "\n")
	knots := [10][2]int{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
	visitedTPart2[knots[8]] = true
	for _, hMove := range input {
		dirAndSteps := strings.Split(hMove, " ")
		dir := dirAndSteps[0]
		steps, _ := strconv.Atoi(dirAndSteps[1])
		log.Printf("move: %v", hMove)
		for i := 0; i < steps; i++ {
			switch dir {
			case "U":
				knots[0][1]++
			case "R":
				knots[0][0]++
			case "D":
				knots[0][1]--
			case "L":
				knots[0][0]--
			}
			for k := 1; k < 10; k++ {
				gap := calcGap(knots[k], knots[k-1])
				gapX := knots[k-1][0] - knots[k][0]
				gapY := knots[k-1][1] - knots[k][1]
				if gap > math.Sqrt(2) {
					knots[k] = [2]int{knots[k][0] + shiftDir(gapX), knots[k][1] + shiftDir(gapY)}
				}
			}
			visitedTPart2[knots[9]] = true
		}
	}
	return len(visitedTPart2)
}

func shiftDir(diff int) int {
	if diff > 0 {
		return 1
	}
	if diff == 0 {
		return 0
	}
	return -1
}
