package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

type Vector struct {
	row int
	col int
}

func parseLines(lines []string) (map[rune][]Vector, int, int) {
	m := make(map[rune][]Vector)
	for row, line := range lines {
		for col, char := range line {
			if unicode.IsUpper(char) || unicode.IsLower(char) || unicode.IsDigit(char) {
				m[char] = append(m[char], Vector{row: row, col: col})
			}

		}

	}

	row_size := len(lines)
	col_size := len(lines[0])
	return m, row_size, col_size
}
func main() {
	lines := readLines(os.Args[1])

	antennaMap, row_size, col_size := parseLines(lines)
	fmt.Println(antennaMap)

	result := calculatePart1(antennaMap, row_size, col_size)
	fmt.Printf("part1: %d\n", result)

	result2 := calculatePart2(antennaMap, row_size, col_size)
	fmt.Printf("part2: %d\n", result2)
}

func diff(a Vector, b Vector) Vector {
	return Vector{row: b.row - a.row, col: b.col - a.col}

}

func add(a Vector, b Vector) Vector {
	return Vector{row: a.row + b.row, col: a.col + b.col}
}

func scale(a Vector, scale int) Vector {
	return Vector{row: a.row * scale, col: a.col * scale}
}

func invert(a Vector) Vector {
	return scale(a, -1)
}

func checkAntinodeInsideMap(antinode Vector, row_size int, col_size int) bool {
	return antinode.row >= 0 && antinode.row < row_size && antinode.col >= 0 && antinode.col < col_size

}

func calculatePart1(antennaMap map[rune][]Vector, row_size int, col_size int) int {
	antinodes := make(map[Vector]bool)
	for _, antennasVectors := range antennaMap {
		for i := 0; i < len(antennasVectors); i++ {
			aVector := antennasVectors[i]
			for j, bVector := range antennasVectors {
				if i == j {
					continue
				}
				diffVector := diff(aVector, bVector)
				firstAntinode := add(aVector, invert(diffVector))
				if checkAntinodeInsideMap(firstAntinode, row_size, col_size) {
					antinodes[firstAntinode] = true
				}
				secondAntinode := add(bVector, diffVector)
				if checkAntinodeInsideMap(secondAntinode, row_size, col_size) {
					antinodes[secondAntinode] = true
				}

			}
		}

	}

	return len(antinodes)
}

func calculatePart2(antennaMap map[rune][]Vector, row_size int, col_size int) int {
	antinodes := make(map[Vector]bool)
	for _, antennasVectors := range antennaMap {
		for i := 0; i < len(antennasVectors); i++ {
			aVector := antennasVectors[i]
			for j, bVector := range antennasVectors {
				if i == j {
					continue
				}
				diffVector := diff(aVector, bVector)
				antinodes[aVector] = true
				antinodes[bVector] = true
				firstAntinode := add(aVector, invert(diffVector))
				for checkAntinodeInsideMap(firstAntinode, row_size, col_size) {
					antinodes[firstAntinode] = true
					firstAntinode = add(firstAntinode, invert(diffVector))
				}
				secondAntinode := add(bVector, diffVector)
				for checkAntinodeInsideMap(secondAntinode, row_size, col_size) {
					antinodes[secondAntinode] = true
					secondAntinode = add(secondAntinode, diffVector)
				}

			}
		}

	}

	return len(antinodes)
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
