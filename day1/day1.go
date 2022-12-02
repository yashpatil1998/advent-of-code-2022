package day1

import (
	"sort"
	"strconv"
	"strings"
)

func Day1Logic() (int, int) {
	var response []int
	listOfItems := strings.Split(input1, "\n\n")
	for _, items := range listOfItems {
		caloriesList := strings.Split(items, "\n")
		sumOfCalories := 0
		for _, cal := range caloriesList {
			calInt, _ := strconv.Atoi(cal)
			sumOfCalories += calInt
		}
		response = append(response, sumOfCalories)
	}
	sort.Ints(response)
	return response[len(response)-1], response[len(response)-1] + response[len(response)-2] + response[len(response)-3]
}
