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
	result := calculatePart1(start, obstacles, row_size, col_size)
	fmt.Printf("part1: %d\n", result)

	// result2 := calculatePart2(string(data))
	// fmt.Printf("part2: %d\n", result2)
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
func calculatePart1(start Vector, obstacles map[Vector]bool, row_size int, col_size int) int {
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
		fmt.Println(currentPos)
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

	return len(positions)
}

func calculatePart2(text string) int {
	result := 0
	return result
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
