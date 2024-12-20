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

type VectorStart struct {
	vector Vector
	start  Vector
}

type Vector struct {
	row int
	col int
}

func add(a Vector, b Vector) Vector {
	return Vector{row: a.row + b.row, col: a.col + b.col}
}
func search(grid [][]rune, label rune, vector Vector, visitedMap map[Vector]bool, group []Vector) (int, int, []Vector) {
	up := Vector{row: -1, col: 0}
	down := Vector{row: 1, col: 0}
	right := Vector{row: 0, col: 1}
	left := Vector{row: 0, col: -1}
	directions := []Vector{right, down, left, up}

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

	return result, groups
}

func findSides(grid [][]rune, group []Vector) int {
	if len(group) == 0 {
		return 0
	}

	up := Vector{row: -1, col: 0}
	down := Vector{row: 1, col: 0}
	right := Vector{row: 0, col: 1}
	left := Vector{row: 0, col: -1}
	directions := []Vector{right, down, left, up}

	horizontals := make(map[VectorStart]bool)
	verticals := make(map[VectorStart]bool)

	for _, vector := range group {
		label := grid[vector.row][vector.col]
		for _, dir := range directions {
			newPosition := add(vector, dir)
			if isOutside(grid, newPosition) || grid[newPosition.row][newPosition.col] != label {
				if dir == up || dir == down {
					newVector := Vector{newPosition.row, 0}
					newVectorStart := VectorStart{newVector, vector}
					horizontals[newVectorStart] = true
				}
				if dir == right || dir == left {
					newVector := Vector{0, newPosition.col}
					newVectorStart := VectorStart{newVector, vector}
					verticals[newVectorStart] = true
				}

			}

		}

	}

	// horizontals
	countHorizontals := 0
	for {
		if len(horizontals) <= 0 {
			break
		}
		vectorStart := getFirstKey(horizontals)
		delete(horizontals, vectorStart)
		countHorizontals++

		queue := []VectorStart{vectorStart}

		for {
			if len(queue) <= 0 {
				break
			}
			newItem := queue[0]
			queue = queue[1:]
			start := newItem.start
			vector := newItem.vector

			newVectorStartRight := VectorStart{vector, Vector{start.row, start.col + 1}}
			if _, ok := horizontals[newVectorStartRight]; ok {

				delete(horizontals, newVectorStartRight)
				queue = append(queue, newVectorStartRight)
			}

			newVectorStartLeft := VectorStart{vector, Vector{start.row, start.col - 1}}
			if _, ok := horizontals[newVectorStartLeft]; ok {
				delete(horizontals, newVectorStartLeft)
				queue = append(queue, newVectorStartLeft)
			}

		}
	}

	// verticals
	countVerticals := 0
	for {
		if len(verticals) <= 0 {
			break
		}
		vectorStart := getFirstKey(verticals)
		delete(verticals, vectorStart)
		countVerticals++

		queue := []VectorStart{vectorStart}

		for {
			if len(queue) <= 0 {
				break
			}
			newItem := queue[0]
			queue = queue[1:]
			start := newItem.start
			vector := newItem.vector

			newVectorStartDown := VectorStart{vector, Vector{start.row + 1, start.col}}
			if _, ok := verticals[newVectorStartDown]; ok {

				delete(verticals, newVectorStartDown)
				queue = append(queue, newVectorStartDown)
			}

			newVectorStartUp := VectorStart{vector, Vector{start.row - 1, start.col}}
			if _, ok := verticals[newVectorStartUp]; ok {
				delete(verticals, newVectorStartUp)
				queue = append(queue, newVectorStartUp)
			}

		}
	}

	return countHorizontals + countVerticals
}
func getFirstKey(m map[VectorStart]bool) VectorStart {
	for k := range m {
		return k
	}
	return VectorStart{}
}
func calculatePart2(grid [][]rune, groups [][]Vector) int {
	totalPrice := 0

	for _, group := range groups {
		area := len(group)
		sides := findSides(grid, group)
		price := area * sides
		totalPrice += price
	}

	return totalPrice
}

func isOutside(grid [][]rune, vector Vector) bool {
	return vector.row < 0 || vector.row >= len(grid) || vector.col < 0 || vector.col >= len(grid[0])
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
