package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type Vector struct {
	row int
	col int
}

func parseLines(lines []string) ([][]rune, map[Vector]bool, Vector, Vector, map[Vector]int) {
	walls := make(map[Vector]bool)
	dots := make(map[Vector]int)

	rows := len(lines)
	cols := len(lines[0])
	grid := make([][]rune, rows)
	for i := range grid {
		grid[i] = make([]rune, cols)
	}

	var startPosition Vector
	var endPosition Vector
	for row, line := range lines {
		for col, char := range line {
			grid[row][col] = char
			if char == '#' {
				walls[Vector{row, col}] = true
			}
			if char == 'S' {
				startPosition = Vector{row, col}
			}
			if char == 'E' {
				endPosition = Vector{row, col}
			}
			if char == '.' {
				dots[Vector{row, col}] = 0
			}

		}
	}

	return grid, walls, startPosition, endPosition, dots
}

func bfs(grid [][]rune, wallMap map[Vector]bool, start Vector, end Vector) ([]Vector, int) {
	// // Mark walls
	// wallMap := make()
	// for i := 0; i < numBytes && i < len(walls); i++ {
	// 	wallMap[walls[i]] = true
	// }

	// Initialize queue with start position
	queue := []Vector{start}

	// Store both distance and parent for each visited node
	type NodeInfo struct {
		distance int
		parent   Vector
	}
	visited := make(map[Vector]NodeInfo)
	visited[start] = NodeInfo{distance: 0, parent: Vector{-1, -1}} // Use invalid coords for start parent

	// Possible directions: up, down, left, right
	directions := []Vector{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == end {
			// Reconstruct path
			path := []Vector{}
			for curr := current; curr != (Vector{-1, -1}); curr = visited[curr].parent {
				path = append([]Vector{curr}, path...)
			}
			return path, visited[current].distance
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
				visited[next] = NodeInfo{
					distance: visited[current].distance + 1,
					parent:   current,
				}
				queue = append(queue, next)
			}
		}
	}

	return nil, -1 // No path found
}

func sortVectors(vectors []Vector) []Vector {
	sorted := make([]Vector, len(vectors))
	copy(sorted, vectors)

	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].row == sorted[j].row {
			return sorted[i].col < sorted[j].col
		}
		return sorted[i].row < sorted[j].row
	})

	return sorted
}

type CheatPosition struct {
	start     Vector
	end       Vector
	timeSaved int
}

func addVector(v1, v2 Vector) Vector {
	return Vector{v1.row + v2.row, v1.col + v2.col}
}
func (c CheatPosition) key() string {
	// Always create key with smaller position first
	if c.start.row < c.end.row || (c.start.row == c.end.row && c.start.col < c.end.col) {
		return fmt.Sprintf("%d,%d-%d,%d", c.start.row, c.start.col, c.end.row, c.end.col)
	}
	return fmt.Sprintf("%d,%d-%d,%d", c.end.row, c.end.col, c.start.row, c.start.col)
}

func calculatePart1(grid [][]rune, walls map[Vector]bool, startPosition Vector, endPosition Vector) int {
	possibleMoves := []Vector{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1}, // up, down, left, right
	}

	normalPath, normalPathLength := bfs(grid, walls, startPosition, endPosition)
	fmt.Printf("Normal path length: %d\n", normalPathLength)

	// Create distance-to-end map from normal path
	distanceToEnd := make(map[Vector]int)
	for i, pos := range normalPath {
		distanceToEnd[pos] = len(normalPath) - i - 1
	}

	cheatPositions := make(map[string]int)
	cheatDistanceOccurences := make(map[int]int)

	for i := 0; i < len(normalPath); i++ {
		currentPos := normalPath[i]

		for _, direction := range possibleMoves {
			potentialCheatPosition1 := addVector(currentPos, direction)
			_, isWall := walls[potentialCheatPosition1]
			if !isWall {
				continue
			}

			potentialCheatPosition2 := addVector(potentialCheatPosition1, direction)
			_, isWall2 := walls[potentialCheatPosition2]

			cheatEndPosition := potentialCheatPosition2
			wallCount := 1
			if isWall2 {
				cheatEndPosition = addVector(potentialCheatPosition2, direction)
				wallCount = 2
			}

			// Validate end position is on track
			if cheatEndPosition.row <= 0 || cheatEndPosition.row >= len(grid)-1 ||
				cheatEndPosition.col <= 0 || cheatEndPosition.col >= len(grid[0])-1 ||
				(grid[cheatEndPosition.row][cheatEndPosition.col] != '.' &&
					grid[cheatEndPosition.row][cheatEndPosition.col] != 'E') {
				continue
			}

			// If end position is on normal path, we can calculate the new path length
			if endDistance, onPath := distanceToEnd[cheatEndPosition]; onPath {
				shortcutPathLength := i + 1 + wallCount + endDistance // distance to current + step to wall + walls + distance from cheat end to goal

				if shortcutPathLength < normalPathLength {
					timeSaved := normalPathLength - shortcutPathLength
					currentCheat := CheatPosition{currentPos, cheatEndPosition, timeSaved}
					cheatKey := currentCheat.key()

					if _, exists := cheatPositions[cheatKey]; !exists {
						// fmt.Printf("Found cheat: from (%d,%d) through %d walls to (%d,%d) saves %d picoseconds\n",
						// 	currentPos.row, currentPos.col,
						// 	wallCount,
						// 	cheatEndPosition.row, cheatEndPosition.col,
						// 	timeSaved)

						cheatPositions[cheatKey] = timeSaved
						cheatDistanceOccurences[timeSaved]++
					}
				}
			}
		}
	}

	// debug prints
	// var savings []int
	// for saved := range cheatDistanceOccurences {
	// 	savings = append(savings, saved)
	// }
	// sort.Ints(savings)
	//
	// for _, saved := range savings {
	// 	count := cheatDistanceOccurences[saved]
	// 	if count == 1 {
	// 		fmt.Printf("There is one cheat that saves %d picoseconds.\n", saved)
	// 	} else {
	// 		fmt.Printf("There are %d cheats that save %d picoseconds.\n", count, saved)
	// 	}
	// }

	// Count cheats that save at least 100 picoseconds
	result := 0
	for _, timeSaved := range cheatPositions {
		if timeSaved >= 100 {
			result++
		}
	}

	return result
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

type CheatPath struct {
	start     Vector
	end       Vector
	wallPath  []Vector // The actual walls broken
	timeSaved int
}

func (c CheatPath) key() string {
	// Create unique key including the wall path
	var wallsStr strings.Builder
	for _, w := range c.wallPath {
		wallsStr.WriteString(fmt.Sprintf("_%d,%d", w.row, w.col))
	}
	return fmt.Sprintf("%d,%d-%d,%d%s", c.start.row, c.start.col, c.end.row, c.end.col, wallsStr.String())
}

type CheatState struct {
	pos      Vector
	wallPath []Vector
	steps    int
}

func calculatePart2(grid [][]rune, walls map[Vector]bool, startPosition Vector, endPosition Vector) int {
	directions := []Vector{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	normalPath, normalPathLength := bfs(grid, walls, startPosition, endPosition)

	// Pre-calculate distance to end
	distanceToEnd := make(map[Vector]int, len(normalPath))
	for i, pos := range normalPath {
		distanceToEnd[pos] = len(normalPath) - i - 1
	}

	type CheatKey struct {
		start, end Vector
	}
	cheats := make(map[CheatKey]int, 1000) // Preallocate with estimated capacity

	type QueueItem struct {
		pos   Vector
		steps int
		walls int
	}

	maxRow, maxCol := len(grid)-1, len(grid[0])-1

	visited := make(map[string]bool, 1000)
	queue := make([]QueueItem, 0, 1000)

	var visitKeyBuffer strings.Builder
	visitKeyBuffer.Grow(32)

	for i, currentPos := range normalPath {
		for k := range visited {
			delete(visited, k)
		}
		queue = queue[:0]
		queue = append(queue, QueueItem{currentPos, 0, 0})

		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]

			if current.steps >= 20 {
				continue
			}

			for _, dir := range directions {
				next := Vector{current.pos.row + dir.row, current.pos.col + dir.col}

				if next.row < 0 || next.row > maxRow || next.col < 0 || next.col > maxCol {
					continue
				}

				newSteps := current.steps + 1
				newWalls := current.walls
				if walls[next] {
					newWalls++
					if newWalls > 20 {
						continue
					}
				}

				visitKeyBuffer.Reset()
				fmt.Fprintf(&visitKeyBuffer, "%d,%d,%d", next.row, next.col, newWalls)
				visitKey := visitKeyBuffer.String()

				if !visited[visitKey] {
					visited[visitKey] = true

					if endDist, ok := distanceToEnd[next]; ok && newWalls > 0 {
						totalDist := i + newSteps + endDist
						if totalDist < normalPathLength {
							timeSaved := normalPathLength - totalDist
							cheatKey := CheatKey{currentPos, next}

							if existing, exists := cheats[cheatKey]; !exists || timeSaved > existing {
								cheats[cheatKey] = timeSaved
							}
						}
					}

					queue = append(queue, QueueItem{next, newSteps, newWalls})
				}
			}
		}
	}

	// Optimize counting phase
	result := 0
	for _, timeSaved := range cheats {
		if timeSaved >= 100 {
			result++
		}
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	startTime := time.Now()

	lines := readLines(os.Args[1])
	grid, walls, startPosition, endPosition, _ := parseLines(lines)

	result := calculatePart1(grid, walls, startPosition, endPosition)
	fmt.Printf("Part 1: %d\n", result)

	result2 := calculatePart2(grid, walls, startPosition, endPosition)
	fmt.Printf("Part 2: %d\n", result2)

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
