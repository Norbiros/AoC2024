package main

import (
	_ "embed"
	"fmt"
	"github.com/Norbiros/AoC2024/utils"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Solving \"Day 11: Plutonian Pebbles\"...")
	fmt.Println("Part 1 result:", partOne(input))
	fmt.Println("Part 2 result:", partTwo(input))
}

var cache = make(map[string]int)

func calculateCount(number int, left int) int {
	cacheKey := strconv.Itoa(number) + ":" + strconv.Itoa(left)
	if cachedResult, exists := cache[cacheKey]; exists {
		return cachedResult
	}

	if left == 0 {
		cache[cacheKey] = 1
		return 1
	}

	numberStr := strconv.Itoa(number)
	var count int

	if number == 0 {
		count = calculateCount(1, left-1)
	} else if len(numberStr)%2 == 0 {
		middle := len(numberStr) / 2
		firstNumber := utils.ToInt(numberStr[:middle])
		secondNumber := utils.ToInt(numberStr[middle:])

		count += calculateCount(firstNumber, left-1)
		count += calculateCount(secondNumber, left-1)
	} else {
		count += calculateCount(number*2024, left-1)
	}

	cache[cacheKey] = count
	return count
}

func partOne(input string) int {
	var result int
	for _, char := range strings.Split(input, " ") {
		result += calculateCount(utils.ToInt(char), 25)
	}

	return result
}

func partTwo(input string) int {
	var result int
	for _, char := range strings.Split(input, " ") {
		result += calculateCount(utils.ToInt(char), 75)
	}

	return result
}
