package main

import (
	_ "embed"
	"fmt"
	"github.com/Norbiros/AoC2024/utils"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	inputLines := strings.Split(input, "\n")

	fmt.Println("Solving \"Day 22: Monkey Market\"...")
	fmt.Println("Part 1 result", partOne(inputLines))
	fmt.Println("Part 2 result", partTwo(inputLines))
}

func partOne(input []string) int {
	result := 0
	for _, line := range input {
		number := utils.ToInt(line)
		for i := 0; i < 2000; i++ {
			number = evolveSecretNumber(number)
		}
		result += number
	}
	return result
}

func partTwo(input []string) int {
	results := make(map[string]int)

	for _, line := range input {
		number := utils.ToInt(line)
		numberPrice := utils.ToInt(string(line[len(line)-1]))
		localResults := make(map[string]int)

		differences := make([]string, 4)
		for i := 0; i < 2000; i++ {
			newNumber := evolveSecretNumber(number)
			differences = differences[1:]

			newNumberPrice := utils.ToInt(string(strconv.Itoa(newNumber)[len(strconv.Itoa(newNumber))-1]))

			difference := strconv.Itoa(newNumberPrice - numberPrice)
			differences = append(differences, difference)

			if i >= 4 {
				if _, ok := localResults[strings.Join(differences, ",")]; !ok {
					localResults[strings.Join(differences, ",")] = newNumberPrice
				}
			}

			number = newNumber
			numberPrice = utils.ToInt(string(strconv.Itoa(newNumber)[len(strconv.Itoa(newNumber))-1]))
		}

		for k, v := range localResults {
			results[k] += v
		}
	}

	maxValue := 0
	for _, value := range results {
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue
}

func prune(secretNumber int) int {
	return secretNumber % 16777216
}

func mix(secretNumber, value int) int {
	return secretNumber ^ value
}

func evolveSecretNumber(secretNumber int) int {
	secretNumber = mix(secretNumber, secretNumber*64)
	secretNumber = prune(secretNumber)

	dividedValue := secretNumber / 32
	secretNumber = mix(secretNumber, dividedValue)
	secretNumber = prune(secretNumber)

	secretNumber = mix(secretNumber, secretNumber*2048)
	secretNumber = prune(secretNumber)

	return secretNumber
}
