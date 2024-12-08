package main

import (
	"bufio"
	"fmt"
	"os"
)

type Vector struct {
	row int
	col int
}

func main() {
	lines := readLines(os.Args[1])

	start, obstacles, row_size, col_size := parseLines(lines)
	result, positions := calculatePart1(start, obstacles, row_size, col_size)
	fmt.Printf("part1: %d\n", result)

	result2 := calculatePart2(start, positions, row_size, col_size, obstacles)
	fmt.Printf("part2: %d\n", result2)
}

func parseLines(lines []string) (Vector, map[Vector]bool, int, int) {

	var start Vector
	obstacles := make(map[Vector]bool)
	for row, line := range lines {
		for col, char := range line {
			if char == '#' {
				obstacles[Vector{row: row, col: col}] = true
			}
			if char == '^' {
				start = Vector{row: row, col: col}

			}
		}

	}
	row_size := len(lines)
	col_size := len(lines[0])
	return start, obstacles, row_size, col_size
}

func add(a Vector, b Vector) Vector {
	return Vector{a.row + b.row, a.col + b.col}
}
func calculatePart1(start Vector, obstacles map[Vector]bool, row_size int, col_size int) (int, map[Vector]int) {
	positions := make(map[Vector]int)

	currentPos := start
	var rotations = []Vector{
		{-1, 0}, // up
		{0, 1},  // right
		{1, 0},  // down
		{0, -1}, // left
	}
	currentIndex := 0

	currentAddVector := rotations[currentIndex]

	for currentPos.row >= 0 && currentPos.row < row_size && currentPos.col >= 0 && currentPos.col < col_size {
		if _, exists := obstacles[add(currentPos, currentAddVector)]; exists {
			currentIndex = (currentIndex + 1) % len(rotations)
			currentAddVector = rotations[currentIndex]
			continue
		}
		positions[currentPos] = currentIndex
		currentPos = add(currentPos, currentAddVector)
		if _, exists := positions[currentPos]; exists {
			if positions[currentPos] == currentIndex {
				break
			}
		}
	}

	return len(positions), positions
}

func calculatePart2(start Vector, part1_positions map[Vector]int, row_size int, col_size int, obstacles map[Vector]bool) int {
	validPositions := make(map[Vector]bool)

	for pos := range part1_positions {
		if pos == start || obstacles[pos] {
			continue
		}

		obstacles[pos] = true
		if willCreateLoop(start, 0, obstacles, row_size, col_size) {
			validPositions[pos] = true
		}
		delete(obstacles, pos)
	}

	return len(validPositions)
}

func willCreateLoop(start Vector, startDir int, obstacles map[Vector]bool, row_size, col_size int) bool {
	type State struct {
		pos Vector
		dir int
	}

	visited := make(map[State]struct{})
	current := State{start, startDir}

	rotations := []Vector{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	for len(visited) < row_size*col_size*4 {
		if _, exists := visited[current]; exists {
			return true
		}

		visited[current] = struct{}{}

		// Check if going out of bounds
		nextPos := add(current.pos, rotations[current.dir])
		if nextPos.row < 0 || nextPos.row >= row_size ||
			nextPos.col < 0 || nextPos.col >= col_size {
			return false
		}

		// If obstacle ahead, turn right
		if obstacles[nextPos] {
			current.dir = (current.dir + 1) % 4
			continue
		}

		// Move forward
		current.pos = nextPos

	}
	return false
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
