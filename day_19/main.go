package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

var DP map[string]int

func main() {
	fmt.Println("Solving \"Day 19: Linen Layout\"...")
	fmt.Println("Part 1 result:", partOne(input))
	fmt.Println("Part 2 result:", partTwo(input))
}

func partOne(input string) int {
	sections := strings.Split(input, "\n\n")
	words := strings.Split(sections[0], ", ")
	targets := strings.Split(sections[1], "\n")

	result := 0

	DP = make(map[string]int)
	for _, target := range targets {
		targetWays := ways(words, target)
		if targetWays > 0 {
			result++
		}
	}

	return result
}

func partTwo(input string) int {
	sections := strings.Split(input, "\n\n")
	words := strings.Split(sections[0], ", ")
	targets := strings.Split(sections[1], "\n")

	result := 0

	DP = make(map[string]int)
	for _, target := range targets {
		targetWays := ways(words, target)
		result += targetWays
	}

	return result
}

func ways(words []string, target string) int {
	if value, exists := DP[target]; exists {
		return value
	}
	if target == "" {
		return 1
	}

	result := 0
	for _, word := range words {
		if strings.HasPrefix(target, word) {
			result += ways(words, target[len(word):])
		}
	}
	DP[target] = result
	return result
}
