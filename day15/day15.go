package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Vector struct {
	row int
	col int
}

func parseLines(splittedByNewLines []string) ([][]rune, map[Vector]bool, map[Vector]bool, Vector, []rune) {
	walls := make(map[Vector]bool)
	boxes := make(map[Vector]bool)

	gridInput := splittedByNewLines[0]
	movesInput := splittedByNewLines[1]

	lines := strings.Split(gridInput, "\n")
	rows := len(lines)
	cols := len(lines[0])
	grid := make([][]rune, rows)
	for i := range grid {
		grid[i] = make([]rune, cols)
	}

	var robot_start_Position Vector
	for row, line := range lines {
		for col, char := range line {
			grid[row][col] = char
			if char == 'O' {
				boxes[Vector{row, col}] = true
			}
			if char == '#' {
				walls[Vector{row, col}] = true
			}
			if char == '@' {
				robot_start_Position = Vector{row, col}
			}

		}
	}

	movesInput = strings.ReplaceAll(movesInput, "\n", "")

	moves := []rune(movesInput)

	fmt.Println(grid)
	fmt.Println(walls)
	fmt.Println(boxes)

	fmt.Println(robot_start_Position)
	fmt.Println(moves)
	return grid, walls, boxes, robot_start_Position, moves
}

// func scalarMult(a Vector, num int) Vector {
// 	return Vector{a.x * num, a.y * num}
// }

func add(a Vector, b Vector) Vector {
	return Vector{a.row + b.row, a.col + b.col}
}

// func mod(a, b int) int {
// 	result := a % b
// 	if result < 0 {
// 		for {
// 			if result > 0 {
// 				return result
// 			}
// 			result += b
// 		}
// 	}
// 	return result
// }

//	func modulo(a Vector, wide int, tall int) Vector {
//		return Vector{mod(a.x, wide), mod(a.y, tall)}
//	}

func calculatePart1(grid [][]rune, walls map[Vector]bool, boxes map[Vector]bool, robot Vector, moves []rune) int {
	result := 0

	directions := make(map[rune]Vector)

	directions['^'] = Vector{-1, 0}
	directions['v'] = Vector{1, 0}
	directions['>'] = Vector{0, 1}
	directions['<'] = Vector{0, -1}

	for _, move := range moves {
		displayGrid(grid, boxes, walls, robot)
		direction := directions[move]

		potentialNextPosition := add(robot, direction)

		if _, ok := walls[potentialNextPosition]; ok {
			continue
		}
		if grid[potentialNextPosition.row][potentialNextPosition.col] == '.' {
			robot = potentialNextPosition
			continue
		}

		var toMoveBoxes []Vector
		if _, ok := boxes[potentialNextPosition]; ok {

			toMoveBoxes = append(toMoveBoxes, potentialNextPosition)

			next := potentialNextPosition
			for {
				next = add(next, direction)

				if _, ok := boxes[next]; ok {
					toMoveBoxes = append(toMoveBoxes, next)
				}

				if grid[next.row][next.col] == '.' {
					for _, toMoveBox := range toMoveBoxes {
						fmt.Println(toMoveBox)
						delete(boxes, toMoveBox)
					}
					for _, toMoveBox := range toMoveBoxes {
						newBoxPos := add(toMoveBox, direction)
						boxes[newBoxPos] = true
					}
					robot = potentialNextPosition

					break
				}
				if grid[next.row][next.col] == '#' {
					break
				}
			}
		}

	}

	for box := range boxes {
		result += 100*box.row + box.col

	}

	return result
}
func displayGrid(grid [][]rune, boxes map[Vector]bool, walls map[Vector]bool, robot Vector) {
	grid = resetGrid(grid)
	for box, _ := range boxes {
		grid[box.row][box.col] = 'O'
	}

	for wall, _ := range walls {
		grid[wall.row][wall.col] = '#'
	}

	grid[robot.row][robot.col] = '@'
	printGrid(grid)

}

func resetGrid(grid [][]rune) [][]rune {
	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}
	return grid
}

// func setGrid(grid [][]rune, positions []Vector) [][]rune {
// 	for _, vector := range positions {
// 		grid[vector.y][vector.x] = '#'
// 	}
// 	return grid
//
// }
//
// func calculatePart2(startPositions []Vector, velocities []Vector) uint64 {}

func printGrid(grid [][]rune) {

	var builder strings.Builder
	for _, row := range grid {
		builder.WriteString(string(row))
		builder.WriteString("\n")
	}
	result := builder.String()
	fmt.Print(result)

}

func main() {
	startTime := time.Now()

	lines := readLines(os.Args[1])
	grid, walls, boxes, robot_start_Position, moves := parseLines(lines)

	result := calculatePart1(grid, walls, boxes, robot_start_Position, moves)
	fmt.Printf("part1: %d\n", result)

	// result2 := calculatePart2(startPositions, velocities)
	// fmt.Printf("part2: %d\n", result2)

	executionTime := time.Since(startTime)
	fmt.Printf("Execution time: %v\n", executionTime)
}

func readLines(filename string) []string {
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	return strings.Split(string(file), "\n\n")

}
