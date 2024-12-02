package main

import (
	_ "embed"
	"fmt"
	"github.com/Norbiros/AoC2024/utils"
	"math"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var processedInput [][]int
	for _, line := range strings.Split(input, "\n") {
		var report []int
		for _, number := range strings.Fields(line) {
			report = append(report, utils.ToInt(number))
		}
		processedInput = append(processedInput, report)
	}

	fmt.Println("Solving \"Day 2: Red-Nosed Reports\"...")
	fmt.Println("Part 1 result", partOne(processedInput))
	fmt.Println("Part 2 result", partTwo(processedInput))
}

func isSafe(values []int, startingValue int) int {
	previousValue := values[0]
	isIncreasing := values[startingValue] > values[0]

	errors := 0

	for _, value := range values[startingValue:] {
		difference := int(math.Abs(float64(value - previousValue)))
		if difference > 3 || difference < 1 {
			errors += 1
			continue
		}

		if isIncreasing != (value > previousValue) {
			errors += 1
			continue
		}

		previousValue = value
	}

	return errors
}

func partOne(input [][]int) int {
	var result int

	for _, report := range input {
		if isSafe(report, 1) == 0 {
			result += 1
		}
	}

	return result
}

func partTwo(input [][]int) int {
	var result int

	for _, report := range input {
		// Probably there's a better solution to this,
		// But I just implemented specific checks for edge cases 0 and 1
		isSafeWhenZeroRemoved := isSafe(report, 2) == 0
		isSafeWhenOneRemoved := isSafe(report[1:], 1) == 0
		isSafeInOtherCases := isSafe(report, 1) <= 1

		if isSafeWhenZeroRemoved || isSafeWhenOneRemoved || isSafeInOtherCases {
			result += 1
		}
	}

	return result
}
