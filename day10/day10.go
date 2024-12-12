package main

import (
	"bufio"
	"fmt"
	"os"
)

type position struct {
	row, col int
}

func parseLines(lines []string) ([][]int, []position, []position) {
	result := make([][]int, len(lines))
	for i := range result {
		result[i] = make([]int, len(lines[0]))
	}

	var zerosPositions []position
	var ninesPositions []position
	for row, line := range lines {
		for col, char := range line {
			result[row][col] = int(char - '0')
			if int(char-'0') == 0 {
				zerosPositions = append(zerosPositions, position{row, col})
			}
			if int(char-'0') == 9 {
				ninesPositions = append(ninesPositions, position{row, col})
			}

		}
	}
	return result, zerosPositions, ninesPositions
}

func calculatePart1(grid [][]int, zeros []position) int {
	totalScore := 0
	rows, cols := len(grid), len(grid[0])

	for _, zero := range zeros {
		visited := make(map[position]bool)
		queue := []position{{zero.row, zero.col}}
		ninesReached := make(map[position]bool)

		// BFS from this trailhead
		for len(queue) > 0 {
			pos := queue[0]
			queue = queue[1:]

			if grid[pos.row][pos.col] == 9 {
				ninesReached[pos] = true
				continue
			}

			directions := []position{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
			for _, dir := range directions {
				newRow, newCol := pos.row+dir.row, pos.col+dir.col
				newPos := position{newRow, newCol}

				if newRow >= 0 && newRow < rows &&
					newCol >= 0 && newCol < cols &&
					!visited[newPos] &&
					grid[newRow][newCol]-grid[pos.row][pos.col] == 1 {
					visited[newPos] = true
					queue = append(queue, newPos)
				}
			}
		}

		totalScore += len(ninesReached)
	}

	return totalScore
}

func calculatePart2(grid [][]int, zeros []position, nines []position) int {
	rows, cols := len(grid), len(grid[0])
	zeroValues := make(map[position]int)

	for _, nine := range nines {
		pathsTo := make(map[position]int)
		queue := []position{nine}
		pathsTo[nine] = 1

		for len(queue) > 0 {
			pos := queue[0]
			queue = queue[1:]
			currentHeight := grid[pos.row][pos.col]
			waysFromHere := pathsTo[pos]

			directions := []position{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
			for _, dir := range directions {
				newRow, newCol := pos.row+dir.row, pos.col+dir.col
				newPos := position{newRow, newCol}

				if newRow >= 0 && newRow < rows &&
					newCol >= 0 && newCol < cols &&
					grid[newRow][newCol] == currentHeight-1 {

					pathsTo[newPos] += waysFromHere

					if pathsTo[newPos] == waysFromHere {
						queue = append(queue, newPos)
					}
				}
			}
		}

		for _, zero := range zeros {
			if paths, exists := pathsTo[zero]; exists {
				zeroValues[zero] += paths
			}
		}
	}
	totalPaths := 0
	for _, paths := range zeroValues {
		totalPaths += paths
	}

	return totalPaths
}

func main() {
	lines := readLines(os.Args[1])
	grid, zeros, nines := parseLines(lines)
	result := calculatePart1(grid, zeros)
	fmt.Printf("part1: %d\n", result)

	result2 := calculatePart2(grid, zeros, nines)
	fmt.Printf("part2: %d\n", result2)
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
