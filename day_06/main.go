package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	inputLines := strings.Split(input, "\n")

	fmt.Println("Solving \"Day 6: Guard Gallivant\"...")
	fmt.Println("Part 1 result", partOne(inputLines))
	fmt.Println("Part 2 result", partTwo(inputLines))
}

func partOne(input []string) int {
	var x, y, direction int

	for checkY, line := range input {
		if index := strings.Index(line, "^"); index >= 0 {
			x, y = index, checkY
			direction = 0
		}
	}

	directionOffsets := []struct{ dx, dy int }{
		{0, -1}, // Up
		{1, 0},  // Right
		{0, 1},  // Down
		{-1, 0}, // Left
	}

	// (x, y) -> [visited directions]
	visited := make(map[[2]int][]int)
	visited[[2]int{x, y}] = []int{direction}

	for {
		var previousX, previousY = x, y
		x += directionOffsets[direction].dx
		y += directionOffsets[direction].dy

		if y < 0 || y >= len(input) || x < 0 || x >= len(input[0]) {
			break
		}

		if input[y][x] == '#' {
			direction = (direction + 1) % 4
			x, y = previousX, previousY
			continue
		}

		for _, visitedDirection := range visited[[2]int{x, y}] {
			if visitedDirection == direction {
				return -1
			}
		}

		visited[[2]int{x, y}] = append(visited[[2]int{x, y}], direction)
	}

	return len(visited)
}

// I could bruteforce it in a smarter way, but this is AoC
func partTwo(input []string) int {
	var result int

	for y, line := range input {
		for x := 0; x < len(line); x++ {
			if line[x] != '.' {
				continue
			}

			input[y] = line[:x] + "#" + line[x+1:]

			if partOne(input) == -1 {
				result++
			}

			input[y] = line
		}
	}

	return result
}
