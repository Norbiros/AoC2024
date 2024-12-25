package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Solving \"Day 25: Code Chronicle\"...")
	fmt.Println("Part 1 result", partOne(input))
	fmt.Println("Part 2 result: Thanks for participating in AoC 2024!")
}

func partOne(input string) int {
	shapes := strings.Split(input, "\n\n")
	var keys [][]string
	var locks [][]string

	for _, rawShape := range shapes {
		shape := strings.Split(rawShape, "\n")
		if shape[0][0] == '#' {
			keys = append(keys, shape)
		} else {
			locks = append(locks, shape)
		}
	}

	var result int
	for _, key := range keys {
		for _, lock := range locks {
			fits := true

			for y, row := range key {
				for x := range row {
					if key[y][x] == '#' && lock[y][x] == '#' {
						fits = false
						break
					}
				}
			}

			if fits {
				result += 1
			}
		}
	}

	return result
}
