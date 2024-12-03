package main

import (
	_ "embed"
	"fmt"
	"github.com/Norbiros/AoC2024/utils"
	"regexp"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Solving \"Day 3: Mull It Over\"...")
	fmt.Println("Part 1 result", partOne(input))
	fmt.Println("Part 2 result", partTwo(input))
}

func partOne(input string) int {
	result := 0

	var regex = regexp.MustCompile(`(?m)mul\((\d\d?\d?),(\d\d?\d?)\)`)

	for _, match := range regex.FindAllStringSubmatch(input, -1) {
		result += utils.ToInt(match[1]) * utils.ToInt(match[2])
	}

	return result
}

func partTwo(input string) int {
	var regex = regexp.MustCompile(`^mul\((\d{1,3}),(\d{1,3})\)`)
	var result int
	isMulEnabled := true

	for len(input) > 0 {
		input = input[1:]

		if strings.HasPrefix(input, "do()") {
			isMulEnabled = true
		} else if strings.HasPrefix(input, "don't()") {
			isMulEnabled = false
		}

		if isMulEnabled {
			for _, match := range regex.FindAllStringSubmatch(input, -1) {
				result += utils.ToInt(match[1]) * utils.ToInt(match[2])
			}
		}
	}

	return result
}
