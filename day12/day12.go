package main

import (
	"bufio"
	"fmt"
	"os"
)

func parseLines(lines []string) [][]rune {
	var result [][]rune
	for _, line := range lines {
		lineRunes := []rune(line)
		result = append(result, lineRunes)

	}
	return result

}

func add(a Vector, b Vector) Vector {
	return Vector{row: a.row + b.row, col: a.col + b.col}
}
func search(grid [][]rune, label rune, vector Vector, visitedMap map[Vector]bool, group []Vector) (int, int, []Vector) {
	up := Vector{row: -1, col: 0}
	down := Vector{row: 1, col: 0}
	right := Vector{row: 0, col: 1}
	left := Vector{row: 0, col: -1}
	directions := []Vector{up, down, right, left}

	visitedMap[vector] = true
	group = append(group, vector)

	area := 1
	perimeter := 0

	for _, dir := range directions {
		newPosition := add(vector, dir)
		if newPosition.row < 0 || newPosition.row >= len(grid) || newPosition.col < 0 || newPosition.col >= len(grid[0]) {
			perimeter++
			continue
		}
		newLabel := grid[newPosition.row][newPosition.col]
		if newLabel != label {
			perimeter++
			continue
		}
		_, ok := visitedMap[newPosition]
		if !ok {
			areaNew, perimeterNew, groupNew := search(grid, newLabel, newPosition, visitedMap, group)
			area += areaNew
			perimeter += perimeterNew
			group = groupNew
		}

	}
	return area, perimeter, group

}

type Vector struct {
	row int
	col int
}

func calculatePart1(grid [][]rune) (int, [][]Vector) {
	visitedMap := make(map[Vector]bool)

	var groups [][]Vector

	result := 0
	for row, line := range grid {
		for col, char := range line {
			vector := Vector{row: row, col: col}
			_, ok := visitedMap[vector]
			if !ok {
				var group []Vector
				area, perimeter, group := search(grid, char, vector, visitedMap, group)
				groups = append(groups, group)
				result += area * perimeter
			}
		}
	}

	for _, group := range groups {
		fmt.Println(group)
		fmt.Println("--")
	}
	return result, groups
}

func calculatePart2(grid [][]rune, groups [][]Vector) int {
	// visitedMap := make(map[Vector]bool)

	result := 0
	// for row, line := range grid {
	// 	for col, char := range line {
	// 		vector := Vector{row: row, col: col}
	// 		_, ok := visitedMap[vector]
	// 		if !ok {
	// 			var group []Vector
	// 			area, perimeter, group := search2(grid, char, vector, visitedMap, group)
	// 			groups = append(groups, group)
	// 			result += area * perimeter
	// 		}
	// 	}
	// }
	//
	// for _, group := range groups {
	// 	fmt.Println(group)
	// 	fmt.Println("--")
	// }
	return result
}

func isOutside(grid [][]rune, vector Vector) bool {
	return vector.row < 0 || vector.row >= len(grid) || vector.col < 0 || vector.col >= len(grid[0])
}

func search2(grid [][]rune, label rune, vector Vector, visitedMap map[Vector]bool, group []Vector) (int, int, []Vector) {
	up := Vector{row: -1, col: 0}
	down := Vector{row: 1, col: 0}
	right := Vector{row: 0, col: 1}
	left := Vector{row: 0, col: -1}
	directions := []Vector{up, down, right, left}

	visitedMap[vector] = true
	group = append(group, vector)

	area := 1
	perimeter := 0

	for _, dir := range directions {
		newPosition := add(vector, dir)
		if isOutside(grid, newPosition) {
			perimeter++
			continue
		}
		newLabel := grid[newPosition.row][newPosition.col]
		if newLabel != label {
			perimeter++
			continue
		}
		_, ok := visitedMap[newPosition]
		if !ok {
			areaNew, perimeterNew, groupNew := search(grid, newLabel, newPosition, visitedMap, group)
			area += areaNew
			perimeter += perimeterNew
			group = groupNew
		}

	}
	return area, perimeter, group

}

func main() {
	lines := readLines(os.Args[1])
	input := parseLines(lines)
	result, groups := calculatePart1(input)
	fmt.Printf("part1: %d\n", result)

	result2 := calculatePart2(input, groups)
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
