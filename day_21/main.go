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
	fmt.Println("Solving \"Day 21: Keypad Conundrum\"...")
	fmt.Println("Part 1 result", solve(2))
	fmt.Println("Part 2 result", solve(25))
}

type Cache struct {
	code  string
	robot int
}

func solve(robots int) int {
	var result int

	cache := make(map[Cache]int)
	for _, line := range strings.Split(input, "\n") {
		code := expandCodeNumericalKeypad(line)
		length := calcLengthAtLevel(code, cache, 0, robots)
		result += length * utils.ToInt(line[:len(line)-1])
	}

	return result
}

func calcLengthAtLevel(code string, cache map[Cache]int, robotNumBefore, maxRobots int) int {
	codeExtended := expandCodeDirectionKeypad(code)

	if robotNumBefore == maxRobots {
		return len(code)
	}

	length := 0
	for _, curr := range splitOnA(codeExtended) {
		cacheEntry := Cache{curr, robotNumBefore + 1}
		if _, ok := cache[cacheEntry]; ok {
			length += cache[cacheEntry]
			continue
		}
		count := calcLengthAtLevel(curr, cache, robotNumBefore+1, maxRobots)
		length += count
	}

	cache[Cache{code, robotNumBefore}] = length
	return length
}

func splitOnA(code string) []string {
	splits := make([]string, 0)
	start := 0
	for i, char := range code {
		if char == 'A' {
			splits = append(splits, code[start:i+1])
			start = i + 1
		}
	}
	return splits
}

func expandCodeNumericalKeypad(code string) string {
	expanded := ""
	curr := 'A'
	for i := 0; i < len(code); i++ {
		from := curr
		to := rune(code[i])
		expanded += numericalFromTo(from, to)
		expanded += "A"
		curr = to
	}

	return expanded
}

func numericalFromTo(from, to rune) string {
	coords := map[rune]utils.Pair{
		'7': {0, 0},
		'8': {1, 0},
		'9': {2, 0},
		'4': {0, 1},
		'5': {1, 1},
		'6': {2, 1},
		'1': {0, 2},
		'2': {1, 2},
		'3': {2, 2},
		'0': {1, 3},
		'A': {2, 3},
	}

	fromCoord := coords[from]
	toCoord := coords[to]

	xDiff := toCoord.First - fromCoord.First
	yDiff := toCoord.Second - fromCoord.Second

	vertical := ""
	for yDiff < 0 {
		vertical += "^"
		yDiff++
	}
	for yDiff > 0 {
		vertical += "v"
		yDiff--
	}

	horizontal := ""
	for xDiff < 0 {
		horizontal += "<"
		xDiff++
	}
	for xDiff > 0 {
		horizontal += ">"
		xDiff--
	}

	xDiff = toCoord.First - fromCoord.First

	if fromCoord.Second == 3 && toCoord.First == 0 {
		return vertical + horizontal
	} else if fromCoord.First == 0 && toCoord.Second == 3 {
		return horizontal + vertical
	} else if xDiff < 0 {
		return horizontal + vertical
	} else {
		return vertical + horizontal
	}
}

func expandCodeDirectionKeypad(code string) string {
	expanded := ""
	curr := 'A'
	for i := 0; i < len(code); i++ {
		from := curr
		to := rune(code[i])
		expanded += directionFromTo(from, to)
		expanded += "A"
		curr = to
	}

	return expanded
}

func directionFromTo(from, to rune) string {
	coords := map[rune]utils.Pair{
		'^': {1, 0},
		'A': {2, 0},
		'<': {0, 1},
		'v': {1, 1},
		'>': {2, 1},
	}

	fromCoord := coords[from]
	toCoord := coords[to]

	xDiff := toCoord.First - fromCoord.First
	yDiff := toCoord.Second - fromCoord.Second

	vertical := ""
	for yDiff < 0 {
		vertical += "^"
		yDiff++
	}
	for yDiff > 0 {
		vertical += "v"
		yDiff--
	}

	horizontal := ""
	for xDiff < 0 {
		horizontal += "<"
		xDiff++
	}
	for xDiff > 0 {
		horizontal += ">"
		xDiff--
	}

	xDiff = toCoord.First - fromCoord.First

	if fromCoord.First == 0 && toCoord.Second == 0 {
		return horizontal + vertical
	} else if fromCoord.Second == 0 && toCoord.First == 0 {
		return vertical + horizontal
	} else if xDiff < 0 {
		return horizontal + vertical
	} else {
		return vertical + horizontal
	}
}
