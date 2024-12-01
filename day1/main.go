package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	s "strings"
)

func main() {
	left, right, err := readInput(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
	}

	slices.Sort(left)
	slices.Sort(right)
	result := calculatePart1(left, right)
	fmt.Printf("part1: %d\n", result)

	result2 := calculatePart2(left, right)
	fmt.Printf("part2: %d\n", result2)
}
func calculatePart1(left, right []int) int {
	result := 0
	for i := 0; i < len(left); i++ {
		result += int(math.Abs(float64(left[i] - right[i])))
	}
	return result
}

func calculatePart2(left, right []int) int {
	m := make(map[int]int)

	for _, value := range right {
		m[value] += 1
	}
	result := 0
	for _, value := range left {
		result += value * m[value]
	}
	return result
}

func readInput(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var left []int
	var right []int
	for scanner.Scan() {
		line := scanner.Text()
		numbersPerLine := s.Fields(line)
		leftNumber, err := strconv.Atoi(numbersPerLine[0])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid numbers: %s", numbersPerLine[0])
		}
		rightNumber, err := strconv.Atoi(numbersPerLine[1])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid numbers: %s", numbersPerLine[1])
		}
		left = append(left, leftNumber)
		right = append(right, rightNumber)
	}
	return left, right, nil
}
