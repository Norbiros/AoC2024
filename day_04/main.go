package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var input = strings.Split(input, "\n")

	fmt.Println("Solving \"Day 4: Ceres Search\"...")
	fmt.Println("Part 1 result", partOne(input))
	fmt.Println("Part 2 result", partTwo(input))
}

func partOne(input []string) int {
	var result = 0

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			var stringsList []string

			// Vertical
			if y < len(input)-3 {
				stringsList = append(stringsList, string(input[y][x])+string(input[y+1][x])+string(input[y+2][x])+string(input[y+3][x]))
			}
			// Horizontal
			if x < len(input)-3 {
				stringsList = append(stringsList, string(input[y][x])+string(input[y][x+1])+string(input[y][x+2])+string(input[y][x+3]))
			}
			// Diagonal
			if x < len(input)-3 && y < len(input)-3 {
				stringsList = append(stringsList, string(input[y][x])+string(input[y+1][x+1])+string(input[y+2][x+2])+string(input[y+3][x+3]))
				stringsList = append(stringsList, string(input[y][x+3])+string(input[y+1][x+2])+string(input[y+2][x+1])+string(input[y+3][x+0]))

			}

			for _, str := range stringsList {
				if str == "XMAS" || str == "SAMX" {
					result++
				}
			}
		}
	}

	return result
}

func partTwo(input []string) int {
	var result = 0

	for y := 0; y < len(input)-2; y++ {
		for x := 0; x < len(input[y])-2; x++ {
			bottomLeft := string(input[y][x]) + string(input[y+1][x+1]) + string(input[y+2][x+2])
			topRight := string(input[y][x+2]) + string(input[y+1][x+1]) + string(input[y+2][x])

			if (bottomLeft == "MAS" || bottomLeft == "SAM") && (topRight == "MAS" || topRight == "SAM") {
				result++
			}
		}
	}

	return result
}
