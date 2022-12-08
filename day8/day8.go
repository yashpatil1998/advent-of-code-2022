package day8

import (
	"sort"
	"strings"
)

func Day8Logic() (int, int) {
	return day8Part1Logic(), day8Part2Logic()
}

func generateMatrix(input []string) [][]string {
	var response [][]string
	for i := 0; i < len(input); i++ {
		lineItem := strings.Split(input[i], "")
		response = append(response, lineItem)
	}
	return response
}

func day8Part1Logic() int {
	input := generateMatrix(strings.Split(input8, "\n"))
	ans := (2 * len(input[0])) + (2 * (len(input) - 2))
	row, col := 1, 1
	for row = 1; row < len(input[0])-1; row++ {
		incrementFlag := false
		for col = 1; col < len(input)-1; col++ {
			incrementFlag = false
			left, right, top, bottom := 0, len(input[0])-1, 0, len(input)-1
			for ; left <= col && !incrementFlag; left++ {
				if input[row][left] < input[row][col] {
					continue
				}
				if input[row][left] >= input[row][col] && left != col {
					break
				}
				if left == col {
					ans++
					incrementFlag = true
				}
			}
			for ; right >= col && !incrementFlag; right-- {
				if input[row][right] < input[row][col] {
					continue
				}
				if input[row][right] >= input[row][col] && right != col {
					break
				}
				if right == col {
					ans++
					incrementFlag = true
				}
			}
			for ; top <= row && !incrementFlag; top++ {
				if input[top][col] < input[row][col] {
					continue
				}
				if input[top][col] >= input[row][col] && top != row {
					break
				}
				if top == row {
					ans++
					incrementFlag = true
				}
			}
			for ; bottom >= row && !incrementFlag; bottom-- {
				if input[bottom][col] < input[row][col] {
					continue
				}
				if input[bottom][col] >= input[row][col] && bottom != row {
					break
				}
				if bottom == row {
					ans++
					incrementFlag = true
				}
			}
		}
	}
	return ans
}

func calculateTreeScore(row int, col int, input [][]string) int {

	h := input[row][col]
	sLeft, sRight, sUp, sDown := 0, 0, 0, 0

	for k := col - 1; k >= 0; k-- {
		if input[row][k] >= h || k == 0 {
			sLeft = col - k
			break
		}
	}

	for k := row - 1; k >= 0; k-- {
		if input[k][col] >= h || k == 0 {
			sUp = row - k
			break
		}
	}

	for k := row + 1; k < len(input); k++ {
		if input[k][col] >= h || k == len(input)-1 {
			sDown = k - row
			break
		}
	}

	for k := col + 1; k < len(input[0]); k++ {
		if input[row][k] >= h || k == len(input[0])-1 {
			sRight = k - col
			break
		}
	}

	return sUp * sDown * sLeft * sRight
}

func day8Part2Logic() int {
	input := generateMatrix(strings.Split(input8, "\n"))
	var scenicScores []int
	i, j := 1, 1
	for i = 1; i < len(input[0])-1; i++ {
		//sScore := 0
		for j = 1; j < len(input)-1; j++ {
			scenicScores = append(scenicScores, calculateTreeScore(i, j, input))
		}
		/* {
			sLeft, sRight, sTop, sBottom := 0, 0, 0, 0
			maxLeft, maxRight, maxTop, maxBottom := 0, len(input[0])-1, 0, len(input)-1
			for k := j - 1; k >= maxLeft; k-- {
				if input[i][j] >= input[i][k] {
					sLeft++
				} else {
					break
				}
			}
			for k := j + 1; k <= maxRight; k++ {
				if input[i][j] >= input[i][k] {
					sRight++
				} else {
					break
				}
			}
			for k := i - 1; k >= maxTop; k-- {
				if input[i][j] >= input[k][j] {
					sTop++
				} else {
					break
				}
			}
			for k := i + 1; k <= maxBottom; k++ {
				if input[i][j] >= input[k][j] {
					sBottom++
				} else {
					break
				}
			}
			sScore = sLeft * sRight * sTop * sBottom
			scenicScores = append(scenicScores, sScore)
		}*/
	}
	sort.Ints(scenicScores)
	return scenicScores[len(scenicScores)-1]
}
