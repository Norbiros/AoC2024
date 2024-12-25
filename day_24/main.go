package main

import (
	_ "embed"
	"fmt"
	"github.com/Norbiros/AoC2024/utils"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	inputLines := strings.Split(input, "\n")
	fmt.Println("Solving \"Day 24: Crossed Wires\"...")
	fmt.Println("Part 1 result", partOne(inputLines))
	fmt.Println("Part 2 result", partTwo(input))
}

func partOne(input []string) int {
	values := make(map[string]int)
	var operations [][]string

	for _, line := range input {
		if strings.Contains(line, ":") {
			parts := strings.Split(line, ": ")
			values[parts[0]] = utils.ToInt(parts[1])
		} else if strings.Contains(line, "->") {
			re := regexp.MustCompile(`\b\w+\b`)
			result := re.FindAllString(line, -1)
			operations = append(operations, result)
		}
	}

	for {
		allSet := true
		for _, operation := range operations {
			valueOne, existsOne := values[operation[0]]
			valueTwo, existsTwo := values[operation[2]]
			operator := operation[1]

			if !existsOne || !existsTwo {
				allSet = false
				continue
			}

			var result bool
			if operator == "OR" {
				result = valueOne == 1 || valueTwo == 1
			} else if operator == "AND" {
				result = valueOne == 1 && valueTwo == 1
			} else if operator == "XOR" {
				result = valueOne != valueTwo
			} else {
				continue
			}

			if result == true {
				values[operation[3]] = 1
			} else {
				values[operation[3]] = 0
			}
		}

		if allSet {
			break
		}

	}

	index := 0
	binaryString := ""

	for {
		key := "z" + fmt.Sprintf("%02d", index)

		value, exists := values[key]
		if !exists {
			break
		}

		binaryString = strconv.Itoa(value) + binaryString

		index++
	}

	result, err := strconv.ParseInt(binaryString, 2, 64)
	if err != nil {
		fmt.Println("Error converting binary to decimal:", err)
	}

	return int(result)
}

// I solved it by generating .dot
// Then converting it to svg dot -Tsvg graph.dot -o output.svg
// And then manually finding errors - it was pretty straight forward
// for my input: cqm,mps,vcv,vjv,vwp,z13,z19,z25
// Visualisation: outputs.svg Visualisation with colored errors: output-error.svg
func partTwo(inputStr string) string {
	input := strings.Split(inputStr, "\n")

	values := make(map[string]int)
	var operations [][]string

	for _, line := range input {
		if strings.Contains(line, ":") {
			parts := strings.Split(line, ": ")
			values[parts[0]] = utils.ToInt(parts[1])
		} else if strings.Contains(line, "->") {
			re := regexp.MustCompile(`\b\w+\b`)
			result := re.FindAllString(line, -1)
			operations = append(operations, result)
		}
	}

	var sb strings.Builder
	sb.WriteString("digraph Circuit {\n")
	sb.WriteString("    rankdir=LR;\n")
	sb.WriteString("    node [shape=ellipse];\n\n")

	inputs := make(map[string]bool)
	outputs := make(map[string]bool)

	for _, op := range operations {
		inputs[op[0]] = true
		inputs[op[2]] = true
		outputs[op[3]] = true
	}

	for input := range inputs {
		sb.WriteString(fmt.Sprintf("    %s [label=\"%s\", color=blue, style=filled, fillcolor=lightblue];\n", input, input))
	}

	for output := range outputs {
		sb.WriteString(fmt.Sprintf("    %s [label=\"%s\", color=red, style=filled, fillcolor=pink];\n", output, output))
	}

	sb.WriteString("\n")

	for i, op := range operations {
		gateName := fmt.Sprintf("gate%d", i+1)
		operator := op[1]
		var gateLabel, gateColor string

		switch operator {
		case "AND":
			gateLabel = "AND"
			gateColor = "green"
		case "OR":
			gateLabel = "OR"
			gateColor = "orange"
		case "XOR":
			gateLabel = "XOR"
			gateColor = "yellow"
		default:
			continue
		}

		sb.WriteString(fmt.Sprintf("    %s [label=\"%s\", shape=box, color=%s];\n", gateName, gateLabel, gateColor))
		sb.WriteString(fmt.Sprintf("    %s -> %s;\n", op[0], gateName))
		sb.WriteString(fmt.Sprintf("    %s -> %s;\n", op[2], gateName))
		sb.WriteString(fmt.Sprintf("    %s -> %s;\n", gateName, op[3]))
	}

	sb.WriteString("}\n")

	err := os.WriteFile("graph.dot", []byte(sb.String()), 0644)
	if err != nil {
		fmt.Printf("Error writing DOT file: %v\n", err)
	}

	return "graph generated in graph.dot. Now manual work :3"
}
