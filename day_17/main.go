package main

import (
	_ "embed"
	"fmt"
	"github.com/Norbiros/AoC2024/utils"
	"log"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var program []int
	registry := make(map[string]int)

	for _, line := range strings.Split(input, "\n") {
		if strings.Contains(line, "Register") {
			text := strings.Replace(line, "Register ", "", 1)
			elements := strings.Split(text, ": ")
			registry[string(elements[0][0])] = utils.ToInt(elements[1])
		} else if strings.Contains(line, "Program") {
			text := strings.Replace(line, "Program: ", "", 1)
			for _, char := range strings.Split(text, ",") {
				program = append(program, utils.ToInt(char))
			}
		}
	}

	fmt.Println("Solving \"Day 17: Chronospatial Computer\"...")
	fmt.Println("Part 1 result", partOne(registry, program))
	fmt.Println("Part 2 result", partTwo(input))
}

func partOne(registry map[string]int, program []int) string {
	var output []string

	instructionPointer := 0
	for instructionPointer < len(program) {
		instruction := program[instructionPointer]
		operand := program[instructionPointer+1]

		if instruction == 0 {
			value := getComboValue(operand, registry)
			divisor := 1 << value.(int)
			r := int(math.Floor(float64(registry["A"]) / float64(divisor)))
			registry["A"] = r
		} else if instruction == 1 {
			registry["B"] = registry["B"] ^ operand
		} else if instruction == 2 {
			value := getComboValue(operand, registry)
			registry["B"] = value.(int) % 8
		} else if instruction == 3 {
			if registry["A"] != 0 {
				instructionPointer = operand
				continue
			}
		} else if instruction == 4 {
			registry["B"] = registry["B"] ^ registry["C"]
		} else if instruction == 5 {
			value := getComboValue(operand, registry)
			output = append(output, strconv.Itoa(value.(int)%8))
		} else if instruction == 6 {
			value := getComboValue(operand, registry)
			divisor := 1 << value.(int)
			r := int(math.Floor(float64(registry["A"]) / float64(divisor)))
			registry["B"] = r
		} else if instruction == 7 {
			value := getComboValue(operand, registry)
			divisor := 1 << value.(int)
			r := int(math.Floor(float64(registry["A"]) / float64(divisor)))
			registry["C"] = r
		}

		instructionPointer += 2
	}

	return strings.Join(output, ",")
}

func getComboValue(operand int, registry map[string]int) interface{} {
	switch operand {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 3
	case 4:
		return registry["A"]
	case 5:
		return registry["B"]
	case 6:
		return registry["C"]
	default:
		log.Panicf("Invalid combo operand: %d", operand)
		return nil
	}
}

// I didn't implement proper solution for this, basicly i was modifying index by huge value like 10^10
// Finding range where suffix matches, then modifing starting number, reducing index modifier and precising suffix match
// And I narrowed it so much I was albe to find it
func partTwo(input string) int {
	index := 190593310993999
	var program []int
	registry := make(map[string]int)

	var programStr string
	for _, line := range strings.Split(input, "\n") {
		if strings.Contains(line, "Register") {
			text := strings.Replace(line, "Register ", "", 1)
			elements := strings.Split(text, ": ")
			registry[string(elements[0][0])] = utils.ToInt(elements[1])
		} else if strings.Contains(line, "Program") {
			text := strings.Replace(line, "Program: ", "", 1)
			for _, char := range strings.Split(text, ",") {
				program = append(program, utils.ToInt(char))
			}
			programStr = text
		}
	}

	for {
		registry["A"] = index
		result := partOne(registry, program)

		if result == programStr {
			break
		}

		if strings.HasSuffix(result, "2,7,5,1,7,4,4,0,3,5,5,3,0") {
			fmt.Println(index, result)
		}

		index += 1
	}

	return index
}
