package main

import (
	_ "embed"
	"fmt"
	"github.com/Norbiros/AoC2024/utils"
	"strconv"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Solving \"Day 9: Disk Fragmenter\"...")
	fmt.Println("Part 1 result", partOne(input))
	fmt.Println("Part 2 result", partTwo(input))
}

// Yes, I could combine it with part 2, but I'm lazy :3
func partOne(input string) int {
	var newString []string

	var fileIndex = 0
	for i, char := range input {
		charInt := utils.ToInt(string(char))

		for j := 0; j < charInt; j++ {
			if i%2 == 0 {
				newString = append(newString, strconv.Itoa(fileIndex))
			} else {
				newString = append(newString, ".")
			}
		}

		if i%2 == 0 {
			fileIndex += 1
		}
	}

	for i := len(newString) - 1; i >= 0; i-- {
		if newString[i] == "." {
			newString = newString[:len(newString)-1]
			continue
		}

		thereAreEmptySpaces := false
		for index := 0; index < len(newString); index++ {
			if newString[index] == "." {
				newString[index] = newString[i]
				thereAreEmptySpaces = true
				break
			}
		}

		if !thereAreEmptySpaces {
			break
		}

		newString = newString[:len(newString)-1]
	}

	var result int
	for i, char := range newString {
		result += utils.ToInt(char) * i
	}

	return result
}

func partTwo(input string) int {
	var newString []string
	var blocks [][]int
	var emptyBlocks [][]int

	var fileIndex = 0
	var positionInFile = 0
	for i, char := range input {
		if i%2 == 0 {
			blocks = append(blocks, []int{positionInFile, utils.ToInt(string(char)), fileIndex})
			for i := 0; i < utils.ToInt(string(char)); i++ {
				newString = append(newString, strconv.Itoa(fileIndex))
				positionInFile += 1
			}
			fileIndex += 1
		} else {
			emptyBlocks = append(emptyBlocks, []int{positionInFile, utils.ToInt(string(char))})
			for i := 0; i < utils.ToInt(string(char)); i++ {
				newString = append(newString, ".")
				positionInFile += 1
			}
		}
	}

	for blockIndex := len(blocks) - 1; blockIndex >= 0; blockIndex-- {
		var block = blocks[blockIndex]
		for emptyBlockIndex, emptyBlock := range emptyBlocks {
			if emptyBlock[0] < block[0] && block[1] <= emptyBlock[1] {
				for i := 0; i < block[1]; i++ {
					newString[block[0]+i] = "."
					newString[emptyBlock[0]+i] = strconv.Itoa(block[2])
				}

				emptyBlocks[emptyBlockIndex] = []int{emptyBlock[0] + block[1], emptyBlock[1] - block[1]}
				break
			}
		}
	}

	var result int
	for i, char := range newString {
		if char != "." {
			result += utils.ToInt(char) * i
		}
	}

	return result
}
