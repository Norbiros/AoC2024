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
	var blockedCoordinates []utils.Pair

	inputLines := strings.Split(input, "\n")
	for _, coordinates := range inputLines {
		splitCoordinates := strings.Split(coordinates, ",")
		blockedCoordinates = append(blockedCoordinates, utils.Pair{First: utils.ToInt(splitCoordinates[0]), Second: utils.ToInt(splitCoordinates[1])})

	}

	fmt.Println("Solving \"Day 18: RAM Run\"...")
	fmt.Println("Part 1 result", partOne(blockedCoordinates, 1024))
	fmt.Println("Part 2 result", partTwo(blockedCoordinates))
}

func partOne(fullBlockedCoordinates []utils.Pair, maximumValue int) int {
	blockedCoordinates := fullBlockedCoordinates[:maximumValue]

	start := utils.Pair{First: 0, Second: 0}
	end := utils.Pair{First: 70, Second: 70}

	gridSize := 71

	directions := []utils.Pair{
		{First: 0, Second: 1},
		{First: 1, Second: 0},
		{First: 0, Second: -1},
		{First: -1, Second: 0},
	}

	grid := make([][]bool, gridSize)
	for i := range grid {
		grid[i] = make([]bool, gridSize)
	}
	for _, coordinate := range blockedCoordinates {
		if coordinate.First >= 0 && coordinate.First < gridSize && coordinate.Second >= 0 && coordinate.Second < gridSize {
			grid[coordinate.First][coordinate.Second] = true
		}
	}

	queue := []utils.Pair{start}
	visited := make(map[utils.Pair]bool)
	visited[start] = true
	distance := make(map[utils.Pair]int)
	distance[start] = 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == end {
			return distance[current]
		}

		for _, direction := range directions {
			next := utils.Pair{
				First:  current.First + direction.First,
				Second: current.Second + direction.Second,
			}

			if next.First >= 0 && next.First < gridSize &&
				next.Second >= 0 && next.Second < gridSize &&
				!grid[next.First][next.Second] && !visited[next] {

				visited[next] = true
				queue = append(queue, next)
				distance[next] = distance[current] + 1
			}
		}
	}

	return -1
}

func partTwo(blockedCoordinates []utils.Pair) string {
	i := 1024
	for {
		if partOne(blockedCoordinates, i) == -1 {
			coordinates := blockedCoordinates[i-1]
			return fmt.Sprintf("%d,%d", coordinates.First, coordinates.Second)
		}

		i += 1
	}
}
