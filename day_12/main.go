package main

import (
	_ "embed"
	"fmt"
	"github.com/Norbiros/AoC2024/utils"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	inputLines := strings.Split(input, "\n")

	fmt.Println("Solving \"Day 12: Garden Groups\"...")
	fmt.Println("Part 1 result:", solve(inputLines, false))
	fmt.Println("Part 2 result:", solve(inputLines, true))
}

func solve(inputLines []string, countSides bool) int {
	processedCoordinates := make(map[utils.Pair]bool)

	var result int
	for rowIndex, row := range inputLines {
		for colIndex := range row {
			coordinates := utils.Pair{First: colIndex, Second: rowIndex}
			if processedCoordinates[coordinates] {
				continue
			}
			initialProcessedCount := len(processedCoordinates)
			regionPerimeter, regionSides := calculateRegion(coordinates, inputLines, processedCoordinates)
			regionArea := len(processedCoordinates) - initialProcessedCount
			if countSides {
				result += regionArea * regionSides
			} else {
				result += regionArea * regionPerimeter
			}
		}
	}

	return result
}

func isInsideBounds(coordinates utils.Pair, inputLines []string) bool {
	return coordinates.First >= 0 && coordinates.First < len(inputLines) && coordinates.Second >= 0 && coordinates.Second < len(inputLines[coordinates.First])
}

func expandRegion(regionType rune, coordinates utils.Pair, outsideCount *int, region map[utils.Pair]bool, inputLines []string, processedCoordinates map[utils.Pair]bool) {
	if !isInsideBounds(coordinates, inputLines) || rune(inputLines[coordinates.First][coordinates.Second]) != regionType {
		*outsideCount++
		return
	}
	if processedCoordinates[coordinates] {
		return
	}
	processedCoordinates[coordinates] = true
	region[coordinates] = true
	for _, direction := range []utils.Pair{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
		expandRegion(regionType, utils.Pair{First: coordinates.First + direction.First, Second: coordinates.Second + direction.Second}, outsideCount, region, inputLines, processedCoordinates)
	}
}

func calculateRegion(startCoordinates utils.Pair, inputLines []string, processedCoordinates map[utils.Pair]bool) (int, int) {
	region := make(map[utils.Pair]bool)
	var outsideCount, sideCount int
	regionType := rune(inputLines[startCoordinates.First][startCoordinates.Second])
	expandRegion(regionType, startCoordinates, &outsideCount, region, inputLines, processedCoordinates)

	outsideDirections := make(map[utils.Pair][]utils.Pair)
	for rowIndex := range inputLines {
		for colIndex := range inputLines[rowIndex] {
			if !region[utils.Pair{First: rowIndex, Second: colIndex}] {
				continue
			}
			for _, direction := range []utils.Pair{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
				neighborCoordinates := utils.Pair{First: rowIndex + direction.First, Second: colIndex + direction.Second}
				if isInsideBounds(neighborCoordinates, inputLines) && rune(inputLines[neighborCoordinates.First][neighborCoordinates.Second]) == regionType {
					continue
				}
				leftDirection, rightDirection := []int{-direction.Second, direction.First}, []int{direction.Second, -direction.First}
				leftNeighbor := utils.Pair{First: rowIndex + leftDirection[0], Second: colIndex + leftDirection[1]}
				rightNeighbor := utils.Pair{First: rowIndex + rightDirection[0], Second: colIndex + rightDirection[1]}
				if !slices.Contains(outsideDirections[leftNeighbor], direction) && !slices.Contains(outsideDirections[rightNeighbor], direction) {
					sideCount++
				}
				outsideDirections[utils.Pair{First: rowIndex, Second: colIndex}] = append(outsideDirections[utils.Pair{First: rowIndex, Second: colIndex}], direction)
			}
		}
	}
	return outsideCount, sideCount
}
