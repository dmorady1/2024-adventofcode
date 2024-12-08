package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	adjacencyList, updates := parseFile(os.Args[1])
	result := calculatePart1(adjacencyList, updates)
	fmt.Printf("part1: %d\n", result)

	result2 := calculatePart2(adjacencyList, updates)
	fmt.Printf("part2: %d\n", result2)
}

func parseFile(filename string) (map[int][]int, [][]int) {
	data, error := os.ReadFile(filename)
	if error != nil {
		panic(error)
	}

	splits := strings.Split(string(data), "\n\n")
	pages := splits[0]

	adjacencyList := make(map[int][]int)
	for _, edge := range strings.Split(pages, "\n") {
		numbers := strings.Split(edge, "|")
		start, err := strconv.Atoi(numbers[0])
		if err != nil {
			panic(err)

		}

		end, err := strconv.Atoi(numbers[1])
		if err != nil {
			panic(err)

		}
		adjacencyList[start] = append(adjacencyList[start], end)

	}

	updates := splits[1]

	var newUpdates [][]int
	for _, update := range strings.Split(strings.TrimSpace(updates), "\n") {
		var lineInt []int

		line := strings.Split(update, ",")
		for _, value := range line {

			num, err := strconv.Atoi(value)
			if err != nil {
				panic(err)
			}
			lineInt = append(lineInt, num)
		}
		newUpdates = append(newUpdates, lineInt)

	}

	return adjacencyList, newUpdates

}

func isValidOrder(order []int, adjacencyList map[int][]int) bool {
	positions := make(map[int]int)
	for i, num := range order {
		positions[num] = i
	}

	for from, toList := range adjacencyList {
		if fromPos, exists := positions[from]; exists {
			for _, to := range toList {
				if toPos, toExists := positions[to]; toExists {
					if fromPos >= toPos {
						return false
					}
				}
			}
		}
	}
	return true
}

func calculatePart1(adjacencyList map[int][]int, updates [][]int) int {
	result := 0

	for _, update := range updates {
		if isValidOrder(update, adjacencyList) {
			middleIndex := len(update) / 2
			result += update[middleIndex]
		}
	}

	return result
}

func topologicalSort(adjacencyList map[int][]int) []int {
	// Get all vertices first
	vertices := make(map[int]bool)
	for from, toList := range adjacencyList {
		vertices[from] = true
		for _, to := range toList {
			vertices[to] = true
		}
	}

	// Create in-degree map
	inDegree := make(map[int]int)
	for vertex := range vertices {
		inDegree[vertex] = 0
	}

	// Calculate in-degrees
	for _, toList := range adjacencyList {
		for _, to := range toList {
			inDegree[to]++
		}
	}

	// Find all vertices with in-degree 0
	var queue []int
	for vertex := range vertices {
		if inDegree[vertex] == 0 {
			queue = append(queue, vertex)
		}
	}

	var result []int
	for len(queue) > 0 {
		// Get next vertex with in-degree 0
		current := queue[0]
		queue = queue[1:]
		result = append(result, current)

		// For each neighbor, decrease in-degree and check if it becomes 0
		if neighbors, exists := adjacencyList[current]; exists {
			for _, neighbor := range neighbors {
				inDegree[neighbor]--
				if inDegree[neighbor] == 0 {
					queue = append(queue, neighbor)
				}
			}
		}
	}

	return result
}

func calculatePart2(adjacencyList map[int][]int, updates [][]int) int {
	result := 0

	for _, update := range updates {
		updateMap := make(map[int]int)
		for _, value := range update {
			updateMap[value] = value
		}
		if !isValidOrder(update, adjacencyList) {
			adjustedAdjacencyList := make(map[int][]int)
			for _, value := range update {
				adjustedAdjacencyList[value] = adjacencyList[value]
			}
			sorted := topologicalSort(adjustedAdjacencyList)

			var sortedAdjusted []int

			for _, num := range sorted {
				if _, ok := updateMap[num]; ok {
					sortedAdjusted = append(sortedAdjusted, num)
				}
			}

			middleIndex := len(sortedAdjusted) / 2
			result += sortedAdjusted[middleIndex]
		}
	}

	return result
}
