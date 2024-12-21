package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func parseLines(lines []string) [][]int {
	var groups [][]int
	var group []int
	pattern := regexp.MustCompile(`\d+`)

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			if len(group) > 0 {
				groups = append(groups, group)
				group = nil
			}
			continue
		}

		matches := pattern.FindAllString(line, -1)
		for _, match := range matches {
			num, err := strconv.Atoi(match)
			if err != nil {
				continue
			}
			group = append(group, num)
		}
	}

	// Append the last group if not empty
	if len(group) > 0 {
		groups = append(groups, group)
	}

	return groups
}

func add(nums []int, num int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] += num
	}
	return nums
}

func mul(nums []int, num int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] *= num
	}
	return nums
}

func div(nums []int, num int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] /= num
	}
	return nums
}

func subSlices(a []int, b []int) []int {
	var result []int
	for i := 0; i < len(a); i++ {
		result = append(result, a[i]-b[i])
	}
	return result
}

func calculatePart1(groups [][]int) int {
	result := 0

	for _, group := range groups {
		xEquation := []int{group[0], group[2], group[4]}
		yEquation := []int{group[1], group[3], group[5]}

		mulFactorA := yEquation[1]
		mulFactorB := xEquation[1]
		xEquation = mul(xEquation, mulFactorA)
		yEquation = mul(yEquation, mulFactorB)

		aEquation := subSlices(xEquation, yEquation)
		if aEquation[0] == 0 {
			continue
		}
		A := aEquation[2] / aEquation[0]

		xEquation[0] *= A
		xEquation[2] -= xEquation[0]

		if xEquation[1] == 0 {
			continue
		}
		B := xEquation[2] / xEquation[1]

		yEquation[0] *= A
		yEquation[1] *= B
		if yEquation[0]+yEquation[1] != yEquation[2] {
			continue
		}
		if A > 100 || B > 100 {
			continue
		}
		result += A*3 + B*1
	}

	return result
}

func calculatePart2(groups [][]int) int {
	result := 0

	for _, group := range groups {
		xEquation := []int{group[0], group[2], group[4] + 10000000000000}
		yEquation := []int{group[1], group[3], group[5] + 10000000000000}

		mulFactorA := yEquation[1]
		mulFactorB := xEquation[1]
		xEquation = mul(xEquation, mulFactorA)
		yEquation = mul(yEquation, mulFactorB)

		aEquation := subSlices(xEquation, yEquation)
		if aEquation[0] == 0 {
			continue
		}
		A := aEquation[2] / aEquation[0]

		xEquation[0] *= A
		xEquation[2] -= xEquation[0]

		if xEquation[1] == 0 {
			continue
		}
		B := xEquation[2] / xEquation[1]

		yEquation[0] *= A
		yEquation[1] *= B

		if yEquation[0]+yEquation[1] != yEquation[2] {
			continue
		}
		result += A*3 + B*1
	}

	return result
}

func main() {
	startTime := time.Now()

	lines := readLines(os.Args[1])
	groups := parseLines(lines)

	result := calculatePart1(groups)
	fmt.Printf("part1: %d\n", result)

	result2 := calculatePart2(groups)
	fmt.Printf("part2: %d\n", result2)

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
