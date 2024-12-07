package main

import (
	_ "embed"
	"fmt"
	"github.com/Norbiros/AoC2024/utils"
	"gonum.org/v1/gonum/stat/combin"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	inputLines := strings.Split(input, "\n")

	fmt.Println("Solving \"Day 7: Bridge Repair\"...")
	fmt.Println("Part 1 result", partOne(inputLines))
	fmt.Println("Part 2 result", partTwo(inputLines))
}

func partOne(input []string) int {
	return solveTask(input, []string{"+", "*"})
}

func partTwo(input []string) int {
	return solveTask(input, []string{"+", "*", "||"})
}

func solveTask(input []string, symbols []string) int {
	var result int
	numSymbols := len(symbols)

	for _, line := range input {
		values := strings.Split(line, " ")
		resultValue := utils.ToInt(strings.Replace(values[0], ":", "", -1))
		var numberValues []int
		for _, value := range values[1:] {
			numberValues = append(numberValues, utils.ToInt(value))
		}

		var permutationsToGenerate []int
		for i := 0; i < len(numberValues)-1; i++ {
			permutationsToGenerate = append(permutationsToGenerate, numSymbols)
		}

		permutations := combin.Cartesian(permutationsToGenerate)

		for _, permutation := range permutations {
			currentValue := numberValues[0]
			for i, value := range numberValues[1:] {
				switch symbols[permutation[i]] {
				case "+":
					currentValue += value
				case "*":
					currentValue *= value
				case "||":
					currentValue = utils.ToInt(strconv.Itoa(currentValue) + strconv.Itoa(value))
				}
			}

			if currentValue == resultValue {
				result += resultValue
				break
			}
		}
	}

	return result
}
