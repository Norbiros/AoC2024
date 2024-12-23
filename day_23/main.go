package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	inputLines := strings.Split(input, "\n")

	graph := make(map[string][]string)
	for _, line := range inputLines {
		if line == "" {
			continue
		}
		elements := strings.Split(line, "-")
		graph[elements[0]] = append(graph[elements[0]], elements[1])
		graph[elements[1]] = append(graph[elements[1]], elements[0])
	}

	fmt.Println("Solving \"Day 23: LAN Party\"...")
	fmt.Println("Part 1 result:", partOne(graph))
	fmt.Println("Part 2 result:", partTwo(graph))
}

func partOne(graph map[string][]string) int {
	count := 0

	nodes := make([]string, 0, len(graph))
	for node := range graph {
		nodes = append(nodes, node)
	}

	for i := 0; i < len(nodes); i++ {
		a := nodes[i]
		for j := i + 1; j < len(nodes); j++ {
			b := nodes[j]
			for k := j + 1; k < len(nodes); k++ {
				c := nodes[k]
				if isConnected(graph, a, b) && isConnected(graph, a, c) && isConnected(graph, b, c) {
					if strings.HasPrefix(a, "t") || strings.HasPrefix(b, "t") || strings.HasPrefix(c, "t") {
						count++
					}
				}
			}
		}
	}

	return count
}

func partTwo(graph map[string][]string) string {
	var maxClique []string

	nodes := make([]string, 0, len(graph))
	for node := range graph {
		nodes = append(nodes, node)
	}

	bronKerbosch([]string{}, nodes, []string{}, graph, &maxClique)

	sort.Strings(maxClique)
	password := strings.Join(maxClique, ",")

	return password
}

func bronKerbosch(r, p, x []string, graph map[string][]string, maxClique *[]string) {
	if len(p) == 0 && len(x) == 0 {
		if len(r) > len(*maxClique) {
			*maxClique = append([]string{}, r...)
		}
		return
	}

	for i := 0; i < len(p); i++ {
		v := p[i]
		bronKerbosch(
			append(r, v),
			intersect(p, graph[v]),
			intersect(x, graph[v]),
			graph,
			maxClique,
		)
		p = append(p[:i], p[i+1:]...)
		x = append(x, v)
		i--
	}
}

func isConnected(graph map[string][]string, node1, node2 string) bool {
	for _, neighbor := range graph[node1] {
		if neighbor == node2 {
			return true
		}
	}
	return false
}

func intersect(a, b []string) []string {
	set := make(map[string]bool)
	for _, v := range b {
		set[v] = true
	}

	var result []string
	for _, v := range a {
		if set[v] {
			result = append(result, v)
		}
	}

	return result
}
