package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines, err := readLines(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
	}

	result := calculatePart1(lines)
	fmt.Printf("part1: %d\n", result)

	result2 := calculatePart2(lines)
	fmt.Printf("part2: %d\n", result2)
}

func searchX(lines [][]rune, row int, col int) int {
	result := 0
	if row-3 >= 0 {
		test := string([]rune{lines[row][col], lines[row-1][col], lines[row-2][col], lines[row-3][col]})
		if test == "XMAS" {
			result++
		}

	}
	if row+3 < len(lines) {
		test := string([]rune{lines[row][col], lines[row+1][col], lines[row+2][col], lines[row+3][col]})
		if test == "XMAS" {
			result++
		}

	}

	if col-3 >= 0 {
		test := string([]rune{lines[row][col], lines[row][col-1], lines[row][col-2], lines[row][col-3]})
		if test == "XMAS" {
			result++
		}
	}

	if col+3 < len(lines[0]) {
		test := string([]rune{lines[row][col], lines[row][col+1], lines[row][col+2], lines[row][col+3]})
		if test == "XMAS" {
			result++
		}
	}

	// left up
	if col-3 >= 0 && row-3 >= 0 {
		test := string([]rune{lines[row][col], lines[row-1][col-1], lines[row-2][col-2], lines[row-3][col-3]})
		if test == "XMAS" {
			result++
		}

	}

	// right up
	if col+3 < len(lines[0]) && row-3 >= 0 {
		test := string([]rune{lines[row][col], lines[row-1][col+1], lines[row-2][col+2], lines[row-3][col+3]})
		if test == "XMAS" {
			result++
		}

	}

	//right down
	if col+3 < len(lines[0]) && row+3 < len(lines) {
		test := string([]rune{lines[row][col], lines[row+1][col+1], lines[row+2][col+2], lines[row+3][col+3]})
		if test == "XMAS" {
			result++
		}

	}

	// left down
	if col-3 >= 0 && row+3 < len(lines) {
		test := string([]rune{lines[row][col], lines[row+1][col-1], lines[row+2][col-2], lines[row+3][col-3]})
		if test == "XMAS" {
			result++
		}

	}

	return result

}

func calculatePart1(lines [][]rune) int {
	result := 0
	for row, line := range lines {
		for col, char := range line {
			if char == 'X' {
				result += searchX(lines, row, col)
			}
		}
	}
	return result
}

func calculatePart2(lines [][]rune) int {
	result := 0
	for row, line := range lines {
		for col, char := range line {
			if char == 'A' {
				result += searchMAS(lines, row, col)
			}
		}
	}
	return result
}

func isMS(r rune) bool {
	return r == 'M' || r == 'S'
}

func searchMAS(lines [][]rune, row int, col int) int {
	result := 0

	if row-1 >= 0 && col-1 >= 0 && row+1 <= len(lines)-1 && col+1 <= len(lines[0])-1 {
		up_left := lines[row-1][col-1]
		down_right := lines[row+1][col+1]

		down_left := lines[row+1][col-1]
		up_right := lines[row-1][col+1]

		if up_left != down_right && isMS(up_left) && isMS(down_right) && down_left != up_right && isMS(down_left) && isMS(up_right) {
			result++
		}
	}

	return result
}

func readLines(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		chars := []rune(line)
		lines = append(lines, chars)

	}
	return lines, nil
}
