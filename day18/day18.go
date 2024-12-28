package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Vector struct {
	row int
	col int
}

func parseLines(lines []string) []Vector {
	var positions []Vector
	for _, line := range lines {
		lineSlice := strings.Split(line, ",")
		row, _ := strconv.Atoi(lineSlice[1])
		col, _ := strconv.Atoi(lineSlice[0])
		positions = append(positions, Vector{row, col})
	}
	return positions
}

func bfs(grid [][]rune, walls []Vector, start Vector, end Vector, numBytes int) int {
	// Mark walls
	wallMap := make(map[Vector]bool)
	for i := 0; i < numBytes && i < len(walls); i++ {
		wallMap[walls[i]] = true
	}

	// Initialize queue with start position
	queue := []Vector{start}
	visited := make(map[Vector]int)
	visited[start] = 0

	// Possible directions: up, down, left, right
	directions := []Vector{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == end {
			return visited[current]
		}

		for _, dir := range directions {
			next := Vector{current.row + dir.row, current.col + dir.col}

			// Check bounds and walls
			if next.row < 0 || next.row >= len(grid) ||
				next.col < 0 || next.col >= len(grid[0]) ||
				wallMap[next] {
				continue
			}

			// If not visited
			if _, exists := visited[next]; !exists {
				visited[next] = visited[current] + 1
				queue = append(queue, next)
			}
		}
	}

	return -1 // No path found
}

func printGrid(grid [][]rune) {
	var builder strings.Builder
	for _, row := range grid {
		builder.WriteString(string(row))
		builder.WriteString("\n")
	}
	result := builder.String()
	fmt.Print(result)
}

func calculatePart2(grid [][]rune, walls []Vector, start Vector, end Vector, numBytes int) Vector {
	for {
		if bfs(grid, walls, start, end, numBytes) == -1 {
			return walls[numBytes-1]
		}
		numBytes++
	}

}

func main() {
	startTime := time.Now()

	lines := readLines(os.Args[1])
	walls := parseLines(lines)

	wide := 71
	tall := 71

	grid := make([][]rune, tall)
	for i := range grid {
		grid[i] = make([]rune, wide)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	startPosition := Vector{0, 0}
	endPosition := Vector{70, 70}

	result := bfs(grid, walls, startPosition, endPosition, 1024)
	fmt.Printf("Part 1: %d\n", result)

	result2 := calculatePart2(grid, walls, startPosition, endPosition, 1024)
	fmt.Printf("Part 2: %d,%d\n", result2.col, result2.row)

	executionTime := time.Since(startTime)
	fmt.Printf("Execution time: %v\n", executionTime)
}
func readLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
