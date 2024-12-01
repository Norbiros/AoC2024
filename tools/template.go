//go:build ignore

package main

import (
	_ "embed"
	"fmt"
	"log"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Solving \"${TASK_TITLE}\"...")
	fmt.Println("Part 1 result", partOne(input))
	fmt.Println("Part 2 result", partTwo(input))
}

func partOne(input string) int {
	log.Panic("Solution for Part 1 not implemented!")
	return -1
}

func partTwo(input string) int {
	log.Panic("Solution for Part 2 not implemented!")
	return -1
}
