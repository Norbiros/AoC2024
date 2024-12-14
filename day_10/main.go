package main

import (
	_ "embed"
	"fmt"
	"github.com/Norbiros/AoC2024/utils"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	inputLines := strings.Split(input, "\n")
	fmt.Println("Solving \"Day 10: Hoof It\"...")
	fmt.Println("Part 1 result", partOne(inputLines))
	fmt.Println("Part 2 result", partTwo(inputLines))
}

func partOne(input []string) int {
	var result int
	for y, row := range input {
		for x, char := range row {
			if char == '0' {
				visitedPlaces := make(map[utils.Pair]bool)
				result += findMore(input, x, y, visitedPlaces, -1, -1, false)
			}
		}
	}

	return result
}

func findMore(input []string, x, y int, visitedPlaces map[utils.Pair]bool, previousValue int, previousPreviousValue int, isPartTwo bool) int {
	if x < 0 || x >= len(input[0]) || y < 0 || y >= len(input) {
		return 0
	}

	if visitedPlaces[utils.Pair{First: x, Second: y}] {
		return 0
	}

	char := input[y][x]
	if char == '.' {
		return 0
	}

	currentValue := utils.ToInt(string(char))

	if previousValue != -1 && currentValue-previousValue != 1 {
		return 0
	}

	if isPartTwo && currentValue == previousPreviousValue {
		return 0
	}

	visitedPlaces[utils.Pair{First: x, Second: y}] = true

	result := 0
	if currentValue == 9 {
		result = 1
	}

	result += findMore(input, x+1, y, visitedPlaces, currentValue, previousValue, isPartTwo)
	result += findMore(input, x-1, y, visitedPlaces, currentValue, previousValue, isPartTwo)
	result += findMore(input, x, y+1, visitedPlaces, currentValue, previousValue, isPartTwo)
	result += findMore(input, x, y-1, visitedPlaces, currentValue, previousValue, isPartTwo)

	if isPartTwo {
		visitedPlaces[utils.Pair{First: x, Second: y}] = false
	}

	return result
}

func partTwo(input []string) int {
	var result int

	for y, row := range input {
		for x, char := range row {
			if char == '0' {
				visitedPlaces := make(map[utils.Pair]bool)
				result += findMore(input, x, y, visitedPlaces, -1, -1, true)
			}
		}
	}

	return result
}
