package main

import (
	_ "embed"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var processedInput [][]int

	for _, line := range strings.Split(input, "\n") {
		var lineNumbers = strings.Split(line, "   ")

		parsedNumber1, err := strconv.Atoi(lineNumbers[0])
		if err != nil {
			panic("Expected number")
		}

		parsedNumber2, err := strconv.Atoi(lineNumbers[1])
		if err != nil {
			panic("Expected number")
		}

		processedInput = append(processedInput, []int{parsedNumber1, parsedNumber2})
	}

	fmt.Println("Solving \"Day 1: Historian Hysteria\"...")
	fmt.Println("Part 1 result", partOne(processedInput))
	fmt.Println("Part 2 result", partTwo(processedInput))
}

func partOne(input [][]int) int {
	var leftColumn []int
	var rightColumn []int

	for _, numbers := range input {
		leftColumn = append(leftColumn, numbers[0])
		rightColumn = append(rightColumn, numbers[1])
	}

	sort.Ints(leftColumn)
	sort.Ints(rightColumn)

	var result = 0
	for i := 0; i < len(rightColumn); i++ {
		var diff = math.Abs(float64(leftColumn[i] - rightColumn[i]))
		result += int(diff)
	}

	return result
}

func partTwo(input [][]int) int {
	counts := make(map[int]int)
	for _, numbers := range input {
		counts[numbers[1]]++
	}

	var result int
	for _, numbers := range input {
		leftNumber := numbers[0]
		result += leftNumber * counts[leftNumber]
	}

	return result
}
