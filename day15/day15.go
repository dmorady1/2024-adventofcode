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

	return grid, walls, boxes, robot_start_Position, moves
}
func parseLines2(splittedByNewLines []string) ([][]rune, map[Vector]bool, map[Vector]Vector, Vector, []rune, map[Vector]bool) {
	walls := make(map[Vector]bool)
	boxes := make(map[Vector]Vector)
	dots := make(map[Vector]bool)

	gridInput := splittedByNewLines[0]
	movesInput := splittedByNewLines[1]

	gridInput = strings.ReplaceAll(gridInput, "#", "##")
	gridInput = strings.ReplaceAll(gridInput, "O", "[]")
	gridInput = strings.ReplaceAll(gridInput, ".", "..")
	gridInput = strings.ReplaceAll(gridInput, "@", "@.")

	// fmt.Println(gridInput)

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
			if char == '[' {
				boxes[Vector{row, col}] = Vector{row, col + 1}
			}
			if char == ']' {
				boxes[Vector{row, col}] = Vector{row, col - 1}
			}

			if char == '#' {
				walls[Vector{row, col}] = true
			}
			if char == '@' {
				robot_start_Position = Vector{row, col}
			}
			if char == '.' {
				dots[Vector{row, col}] = true
			}

		}
	}

	movesInput = strings.ReplaceAll(movesInput, "\n", "")

	moves := []rune(movesInput)

	return grid, walls, boxes, robot_start_Position, moves, dots
}

func add(a Vector, b Vector) Vector {
	return Vector{a.row + b.row, a.col + b.col}
}

func calculatePart1(grid [][]rune, walls map[Vector]bool, boxes map[Vector]bool, robot Vector, moves []rune) int {
	result := 0

	directions := make(map[rune]Vector)

	directions['^'] = Vector{-1, 0}
	directions['v'] = Vector{1, 0}
	directions['>'] = Vector{0, 1}
	directions['<'] = Vector{0, -1}

	for _, move := range moves {
		grid = resetGrid(grid)
		for box := range boxes {
			grid[box.row][box.col] = 'O'
		}

		for wall := range walls {
			grid[wall.row][wall.col] = '#'
		}

		grid[robot.row][robot.col] = '@'
		// printGrid(grid)
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
						// fmt.Println(toMoveBox)
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

func calculatePart2(grid [][]rune, walls map[Vector]bool, boxes map[Vector]Vector, robot Vector, moves []rune, dots map[Vector]bool) int {
	result := 0

	directions := make(map[rune]Vector)

	directions['^'] = Vector{-1, 0}
	directions['v'] = Vector{1, 0}
	directions['>'] = Vector{0, 1}
	directions['<'] = Vector{0, -1}

	for _, move := range moves {
		grid = resetGrid(grid)
		for box, nextToBox := range boxes {
			if (nextToBox.col - box.col) > 0 {
				grid[nextToBox.row][nextToBox.col] = ']'
				grid[box.row][box.col] = '['
			} else {
				// fmt.Println(nextToBox, box)
				grid[nextToBox.row][nextToBox.col] = '['
				grid[box.row][box.col] = ']'
			}
		}

		for wall := range walls {
			grid[wall.row][wall.col] = '#'
		}

		grid[robot.row][robot.col] = '@'
		// printGrid(grid)
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

		if move == '^' || move == 'v' {
			// fmt.Println("hierherhehrherherhh")
			adjust := moveBoxesVertically(grid, direction, robot, potentialNextPosition, walls, boxes)
			if adjust {
				robot = potentialNextPosition
			}
			continue

		}
		if nextToPotential, ok := boxes[potentialNextPosition]; ok {

			toMoveBoxes = append(toMoveBoxes, potentialNextPosition)
			toMoveBoxes = append(toMoveBoxes, nextToPotential)

			next := nextToPotential
			for {
				next = add(next, direction)

				if nextTo, ok := boxes[next]; ok {
					toMoveBoxes = append(toMoveBoxes, next)
					toMoveBoxes = append(toMoveBoxes, nextTo)
					next = nextTo
				}

				if grid[next.row][next.col] == '.' {
					// fmt.Println("-------------")

					var nextToPositions []Vector
					for _, toMoveBox := range toMoveBoxes {

						nextTo := boxes[toMoveBox]
						nextToPositions = append(nextToPositions, nextTo)
						delete(boxes, toMoveBox)
					}
					for i := 0; i < len(toMoveBoxes); i++ {
						toMoveBox := toMoveBoxes[i]
						nextTo := nextToPositions[i]
						newBoxPos := add(toMoveBox, direction)
						nextToPos := add(nextTo, direction)
						boxes[newBoxPos] = nextToPos
					}
					robot = potentialNextPosition

					// fmt.Println("-------------")
					break
				}
				if grid[next.row][next.col] == '#' {
					break
				}
			}
		}

	}

	leftEdges := make(map[Vector]bool)
	grid = resetGrid(grid)
	for box, nextToBox := range boxes {
		if (nextToBox.col - box.col) > 0 {
			grid[nextToBox.row][nextToBox.col] = ']'
			grid[box.row][box.col] = '['
			leftEdges[box] = true
		} else {
			grid[nextToBox.row][nextToBox.col] = '['
			leftEdges[nextToBox] = true
			grid[box.row][box.col] = ']'
		}
	}

	for wall := range walls {
		grid[wall.row][wall.col] = '#'
	}

	grid[robot.row][robot.col] = '@'
	// printGrid(grid)

	for leftEdge, _ := range leftEdges {
		// fmt.Println(leftEdge)
		result += 100*leftEdge.row + leftEdge.col
	}

	return result
}

func moveBoxesVertically(grid [][]rune, direction Vector, currentPosition Vector, potentialNextPosition Vector, walls map[Vector]bool, boxes map[Vector]Vector) bool {
	var toMoveBoxes []Vector

	queue := []Vector{currentPosition}

	for {
		if len(queue) == 0 {
			break
		}
		currentPotentialPosition := add(queue[0], direction)
		queue = queue[1:]

		if _, ok := walls[currentPotentialPosition]; ok {
			return false
		}

		if grid[currentPotentialPosition.row][currentPotentialPosition.col] == '.' {
			continue
		}
		if nextTo, ok := boxes[currentPotentialPosition]; ok {
			toMoveBoxes = append(toMoveBoxes, currentPotentialPosition)
			toMoveBoxes = append(toMoveBoxes, nextTo)
			queue = append(queue, currentPotentialPosition)
			queue = append(queue, nextTo)
		}

		currentPotentialPosition = add(currentPotentialPosition, direction)
	}

	for _, toMoveBox := range toMoveBoxes {
		// fmt.Println(toMoveBox)
		// nextTo := boxes[toMoveBox]
		// nextToPositions = append(nextToPositions, nextTo)
		delete(boxes, toMoveBox)
	}
	// fmt.Println(boxes)
	// fmt.Println("moving....")
	for i := 0; i < len(toMoveBoxes); i++ {
		toMoveBox := toMoveBoxes[i]
		// nextTo := nextToPositions[i]

		var nextTo Vector
		if grid[toMoveBox.row][toMoveBox.col] == '[' {
			nextTo = Vector{toMoveBox.row, toMoveBox.col + 1}
		} else {
			nextTo = Vector{toMoveBox.row, toMoveBox.col - 1}
		}
		// fmt.Printf("tomoveBox: %v  nextTo %v\n", toMoveBox, nextTo)
		newBoxPos := add(toMoveBox, direction)
		nextToPos := add(nextTo, direction)
		// fmt.Printf("newBoxPos: %v  nextToPos %v\n", newBoxPos, nextToPos)
		boxes[newBoxPos] = nextToPos
	}
	// fmt.Println("moving.end")
	// fmt.Println(boxes)
	return true

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

	grid, walls, boxesNew, robot_start_Position, moves, dots := parseLines2(lines)
	result2 := calculatePart2(grid, walls, boxesNew, robot_start_Position, moves, dots)
	fmt.Printf("part2: %d\n", result2)

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
