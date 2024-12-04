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

// I know it's overkill for this task,
// but i wanted to solve it in O(N),
// so I used *dynamic programming*!
func dpCheck(height []int) int {
	n := len(height)

	increasing := make([]int, n)
	decreasing := make([]int, n)
	for i := range increasing {
		increasing[i] = math.MaxInt32
		decreasing[i] = math.MaxInt32
	}
	increasing[0], decreasing[0] = 0, 0

	for i := 1; i < n; i++ {
		if diff := utils.Abs(height[i] - height[i-1]); diff >= 1 && diff <= 3 {
			if height[i] > height[i-1] {
				increasing[i] = min(increasing[i], increasing[i-1])
			} else if height[i] < height[i-1] {
				decreasing[i] = min(decreasing[i], decreasing[i-1])
			}
		}

		if i > 1 {
			if diff := utils.Abs(height[i] - height[i-2]); diff < 1 || diff > 3 {
				continue
			}

			if height[i] > height[i-2] {
				increasing[i] = min(increasing[i], increasing[i-2]+1)
			} else if height[i] < height[i-2] {
				decreasing[i] = min(decreasing[i], decreasing[i-2]+1)
			}
		}
	}

	return min(increasing[n-1], decreasing[n-1])
}

func partOne(input [][]int) int {
	var result int

	for _, report := range input {
		if dpCheck(report) == 0 {
			result += 1
		}
	}

	return result
}

func partTwo(input [][]int) int {
	var result int

	for _, report := range input {
		// Probably there's a better solution to this,
		// But I just implemented specific checks for edge cases
		isSafeWhenFirstValueRemoved := dpCheck(report[1:]) == 0
		isSafeWhenLastValueRemoved := dpCheck(report[:len(report)-1]) == 0
		otherCasesResult := dpCheck(report)
		isSafeInOtherCases := otherCasesResult == 0 || otherCasesResult == 1

		if isSafeWhenFirstValueRemoved || isSafeWhenLastValueRemoved || isSafeInOtherCases {
			result += 1
		}
	}

	return result
}
