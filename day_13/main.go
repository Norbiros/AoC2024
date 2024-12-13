package main

import (
	_ "embed"
	"fmt"
	"github.com/Norbiros/AoC2024/utils"
	"gonum.org/v1/gonum/mat"
	"math"
	"strings"
)

//go:embed input.txt
var input string

// I know this solution isn't bulletproof
// It won't work, for example, for buttons [1, 1] = 2 [5, 5] = 10
// But it worked for my dataset :3
func main() {
	fmt.Println("Solving \"Day 13: Claw Contraption\"...")
	fmt.Println("Part 1 result", solve(input, 0))
	fmt.Println("Part 2 result", solve(input, 10000000000000))
}

func parseCoordinates(line, prefix string) []int {
	line = strings.ReplaceAll(line, prefix, "")
	line = strings.ReplaceAll(line, ", Y", " ")
	line = strings.ReplaceAll(line, "=", "")
	values := strings.Fields(line)
	return []int{utils.ToInt(values[0]), utils.ToInt(values[1])}
}

func solve(input string, difference int) int {
	var resultValue int
	for _, line := range strings.Split(input, "\n\n") {
		lines := strings.Split(line, "\n")
		buttonA := parseCoordinates(lines[0], "Button A: X")
		buttonB := parseCoordinates(lines[1], "Button B: X")

		prize := parseCoordinates(lines[2], "Prize: X")
		prize[0] += difference
		prize[1] += difference

		matrix := mat.NewDense(2, 2, []float64{
			float64(buttonA[0]), float64(buttonB[0]),
			float64(buttonA[1]), float64(buttonB[1]),
		})
		result := mat.NewDense(2, 1, []float64{
			float64(prize[0]),
			float64(prize[1]),
		})

		var solution mat.Dense
		err := solution.Solve(matrix, result)
		if err != nil {
			fmt.Println("Error solving equations:", err)
			continue
		}

		a, b := solution.At(0, 0), solution.At(1, 0)

		const epsilon = 1e-4
		if math.Abs(a-math.Round(a)) < epsilon && math.Abs(b-math.Round(b)) < epsilon {
			resultValue += int(math.Round(a)*3 + math.Round(b))
		}

	}
	return resultValue
}
