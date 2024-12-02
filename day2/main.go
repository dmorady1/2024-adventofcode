package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	s "strings"
)

func main() {
	numbers, err := readLines(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
	}

	result, unsafes := calculatePart1(numbers)
	fmt.Printf("part1: %d\n", result)

	result2 := calculatePart2(unsafes)
	fmt.Printf("part2: %d\n", result2+result)
}
func check_increase(num1 int, num2 int) bool {
	return num2-num1 > 0
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func calculatePart1(numbers [][]int) (int, [][]int) {
	result := 0
	var unsafes [][]int
	for _, line := range numbers {
		dont_use := false
		var is_increase bool
		for i := 0; i < len(line)-1; i++ {
			if i == 0 {
				is_increase = check_increase(line[i], line[i+1])
			}
			if is_increase != check_increase(line[i], line[i+1]) {
				dont_use = true
				break
			}
			diff := line[i+1] - line[i]
			if abs(diff) < 1 || abs(diff) > 3 {
				dont_use = true
				break
			}
		}
		if !dont_use {
			result++
		} else {
			unsafes = append(unsafes, line)
		}

	}
	return result, unsafes
}

func calculatePart2(unsafes [][]int) int {
	result := 0
	for _, line := range unsafes {
		if canBeMadeValid(line) {
			result++
		}
	}
	return result
}

func canBeMadeValid(line []int) bool {
	for i := 0; i < len(line); i++ {
		newLine := make([]int, 0, len(line)-1)
		newLine = append(newLine, line[:i]...)
		newLine = append(newLine, line[i+1:]...)

		if isValidSequence(newLine) {
			return true
		}
	}
	return false
}

func isValidSequence(line []int) bool {
	if len(line) < 2 {
		return true
	}

	isIncreasing := line[1] > line[0]

	for i := 0; i < len(line)-1; i++ {
		diff := line[i+1] - line[i]

		if (diff > 0) != isIncreasing {
			return false
		}

		if abs(diff) < 1 || abs(diff) > 3 {
			return false
		}
	}
	return true
}

func readLines(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var numbersPerLines [][]int
	for scanner.Scan() {
		line := scanner.Text()
		numbersPerLine := s.Fields(line)
		var numbersPerLineInt []int
		for _, value := range numbersPerLine {
			number, err := strconv.Atoi(value)
			if err != nil {
				return nil, fmt.Errorf("invalid numbers: %s", numbersPerLine[0])
			}
			numbersPerLineInt = append(numbersPerLineInt, number)
		}
		numbersPerLines = append(numbersPerLines, numbersPerLineInt)
	}
	return numbersPerLines, nil
}
