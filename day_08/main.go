package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Solving \"Day 8: Resonant Collinearity\"...")
	fmt.Println("Part 1 result", partOne(input))
	fmt.Println("Part 2 result", partTwo(input))
}

func partOne(input string) int {
	return solve(input, false)
}

func partTwo(input string) int {
	return solve(input, true)
}

func solve(input string, extended bool) int {
	inputLines := strings.Split(input, "\n")
	symbols := make(map[string][][]int)
	antidotes := make(map[string]bool)

	for y, line := range inputLines {
		for x, symbol := range line {
			symbolStr := string(symbol)
			if symbolStr == "." {
				continue
			}

			for _, anotherSymbol := range symbols[symbolStr] {
				xDiff := x - anotherSymbol[0]
				yDiff := y - anotherSymbol[1]

				markAntidotes(antidotes, x, y, xDiff, yDiff, inputLines, extended)
				markAntidotes(antidotes, anotherSymbol[0], anotherSymbol[1], -xDiff, -yDiff, inputLines, extended)

				if extended {
					antidotes[getKey(anotherSymbol[0], anotherSymbol[1])] = true
					antidotes[getKey(x, y)] = true
				}

			}
			symbols[symbolStr] = append(symbols[symbolStr], []int{x, y})
		}
	}

	return len(antidotes)
}

func markAntidotes(antidotes map[string]bool, startX, startY, xDiff, yDiff int, inputLines []string, extended bool) {
	x, y := startX, startY
	for {
		x += xDiff
		y += yDiff
		if x < 0 || x >= len(inputLines[0]) || y < 0 || y >= len(inputLines) {
			break
		}
		antidotes[getKey(x, y)] = true

		if !extended {
			break
		}
	}
}

func getKey(x int, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}
