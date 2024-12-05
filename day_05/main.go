package main

import (
	_ "embed"
	"fmt"
	"github.com/Norbiros/AoC2024/utils"
	"strings"
)

//go:embed input.txt
var input string

// I know the code quality could be better,
// but I spent an unusually long time on this task.
func main() {
	var loadingConstrains = true
	var processedInput [][]int

	dependencies := map[int][]int{}
	dependents := map[int][]int{}

	for _, line := range strings.Split(input, "\n") {
		if loadingConstrains && !strings.Contains(line, "|") {
			loadingConstrains = false
			continue
		}

		if loadingConstrains {
			lineText := strings.Split(line, "|")

			index := utils.ToInt(lineText[1])
			dependencies[index] = append(dependencies[index], utils.ToInt(lineText[0]))

			index2 := utils.ToInt(lineText[0])
			dependents[index2] = append(dependents[index2], utils.ToInt(lineText[1]))
		} else {
			lineText := strings.Split(line, ",")

			var numbers []int
			for _, str := range lineText {
				numbers = append(numbers, utils.ToInt(str))
			}

			processedInput = append(processedInput, numbers)
		}
	}

	fmt.Println("Solving \"Day 5: Print Queue\"...")
	fmt.Println("Part 1 result", partOne(dependencies, processedInput))
	fmt.Println("Part 2 result", partTwo(dependencies, dependents, processedInput))
}

func contains(slice []int, element int) bool {
	for _, value := range slice {
		if value == element {
			return true
		}
	}
	return false
}

func checkConstraints(line []int, dependencies map[int][]int) bool {
	isOkay := true
	var invalidConstrains []int
	var previousElements []int

	for _, el := range line {
		for _, constraint := range dependencies[el] {
			if contains(line, constraint) && !contains(previousElements, constraint) {
				isOkay = false
				invalidConstrains = append(invalidConstrains, constraint)
				break
			}
		}
		previousElements = append(previousElements, el)
	}

	return isOkay
}

func partOne(dependencies map[int][]int, processedInput [][]int) int {
	var result = 0

	for _, line := range processedInput {
		isOkay := checkConstraints(line, dependencies)

		if isOkay {
			result += line[len(line)/2]
		}
	}

	return result
}

func partTwo(dependencies map[int][]int, dependents map[int][]int, processedInput [][]int) int {
	var result = 0

	for _, line := range processedInput {
		isOkay := checkConstraints(line, dependencies)
		if isOkay {
			continue
		}

		var final []int
		var queue []int
		dependencyCount := make(map[int]int)

		for _, value := range line {
			count := 0
			for _, dependency := range dependencies[value] {
				if contains(line, dependency) {
					count++
				}
			}
			dependencyCount[value] = count
			if count == 0 {
				queue = append(queue, value)
			}
		}

		for len(queue) > 0 {
			value := queue[0]
			queue = queue[1:]

			final = append(final, value)

			for _, y := range dependents[value] {
				if count, exists := dependencyCount[y]; exists {
					dependencyCount[y] = count - 1
					if dependencyCount[y] == 0 {
						queue = append(queue, y)
					}
				}
			}
		}

		if len(final) > 0 {
			result += final[len(final)/2]
		}
	}

	return result
}
