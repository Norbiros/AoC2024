package main

import (
	_ "embed"
	"fmt"
	"github.com/Norbiros/AoC2024/utils"
	"strings"
	"time"
)

//go:embed input.txt
var input string

func main() {
	inputLines := strings.Split(input, "\n")

	fmt.Println("Solving \"Day 14: Restroom Redoubt\"...")
	fmt.Println("Part 1 result", partOne(inputLines, 100))

	// For Part 2, I couldn't find a clean solution. Instead, I noticed some patterns:
	// After a few hundred iterations, the simulation started forming vertical and horizontal lines.
	// - Vertical: 1007, 1108
	// - Horizontal: 980, 1083
	// By experimenting, I started looping with a step of 101 and eventually found a Christmas tree pattern.
	// Uncomment the line below to observe the simulation visually:
	//fmt.Println("Part 2 result", partTwo(inputLines))
}

func parseCoordinates(input string) []int {
	values := strings.Split(input, "=")
	numbers := strings.Split(values[1], ",")
	return []int{utils.ToInt(numbers[0]), utils.ToInt(numbers[1])}
}

func partOne(input []string, n int) int {
	const width = 101
	const height = 103
	robots := make(map[utils.Pair]int)

	result := make(map[int]int, 4)
	for _, line := range input {
		values := strings.Split(line, " ")
		coordinates := parseCoordinates(values[0])
		velocity := parseCoordinates(values[1])

		for i := 0; i < n; i++ {
			coordinates[0] += velocity[0]
			coordinates[1] += velocity[1]

			coordinates[0] = (coordinates[0]%width + width) % width
			coordinates[1] = (coordinates[1]%height + height) % height
		}

		if coordinates[0] > (width/2) && coordinates[1] > (height/2) {
			result[0] += 1
		}

		if coordinates[0] > (width/2) && coordinates[1] < (height/2) {
			result[1] += 1
		}

		if coordinates[0] < (width/2) && coordinates[1] < (height/2) {
			result[2] += 1
		}

		if coordinates[0] < (width/2) && coordinates[1] > (height/2) {
			result[3] += 1
		}

		robots[utils.Pair{First: coordinates[0], Second: coordinates[1]}] += 1
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if robots[utils.Pair{First: x, Second: y}] != 0 {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	return result[0] * result[1] * result[2] * result[3]
}

func partTwo(input []string) int {
	for n := 1007; n < 10000000; n += 101 {
		fmt.Print("\033[H\033[2J")

		fmt.Printf("Iteration: %d\n", n)
		partOne(input, n)

		time.Sleep(200 * time.Millisecond)
	}

	return -1
}
