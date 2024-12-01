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
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var left []int
	var right []int
	for scanner.Scan() {
		line := scanner.Text()
		numbersPerLine := s.Fields(line)
		leftNumber, _ := strconv.Atoi(numbersPerLine[0])
		rightNumber, _ := strconv.Atoi(numbersPerLine[1])
		left = append(left, leftNumber)
		right = append(right, rightNumber)
	}

	slices.Sort(left)
	slices.Sort(right)
	result := 0
	for i := 0; i < len(left); i++ {
		result += int(math.Abs(float64(left[i] - right[i])))

	}
	fmt.Printf("part1: %d\n", result)

	m := make(map[int]int)

	for _, value := range right {
		m[value] += 1
	}

	result2 := 0
	for _, value := range left {
		result2 += value * m[value]
	}

	fmt.Printf("part2: %d\n", result2)
}
