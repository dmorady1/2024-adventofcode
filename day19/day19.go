package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func parseLines(splittedByNewLines []string) ([]string, []string) {
	towelPatterns := splittedByNewLines[0]

	towels := strings.Split(towelPatterns, ", ")

	designsString := splittedByNewLines[1]

	designsLines := strings.Split(designsString, "\n")

	var designs []string

	for _, design := range designsLines {
		fmt.Println(design)
		if design != "" || design == "\n" {
			designs = append(designs, design)
		}
	}

	return towels, designs
}

type memoKey struct {
	design string
	test   string
}

func recursiveCheckWithMemo(towels []string, design string, test string, memo map[memoKey]bool) bool {
	key := memoKey{design, test}

	if result, exists := memo[key]; exists {
		return result
	}

	if test == design {
		memo[key] = true
		return true
	}

	designCheck := false
	for _, towel := range towels {
		newTest := test + towel

		if strings.HasPrefix(design, newTest) {
			designCheck = recursiveCheckWithMemo(towels, design, newTest, memo)
		}
		if designCheck {
			memo[key] = true
			return true
		}
	}

	memo[key] = false
	return false
}

func recursiveCheck2WithMemo(towels []string, design string, test string, memo map[memoKey]int) int {
	key := memoKey{design, test}

	if result, exists := memo[key]; exists {
		return result
	}

	if test == design {
		memo[key] = 1
		return 1
	}

	count := 0
	for _, towel := range towels {
		newTest := test + towel

		if strings.HasPrefix(design, newTest) {
			count += recursiveCheck2WithMemo(towels, design, newTest, memo)
		}
	}

	memo[key] = count
	return count
}

func calculatePart1(towels []string, designs []string) int {
	count := 0
	for _, design := range designs {
		memo := make(map[memoKey]bool)
		if recursiveCheckWithMemo(towels, design, "", memo) {
			count++
		}
	}
	return count
}

func calculatePart2(towels []string, designs []string) int {
	count := 0
	for _, design := range designs {
		memoCheck := make(map[memoKey]bool)
		if recursiveCheckWithMemo(towels, design, "", memoCheck) {
			memoCount := make(map[memoKey]int)
			count += recursiveCheck2WithMemo(towels, design, "", memoCount)
		}
	}
	return count
}

func main() {
	startTime := time.Now()

	lines := readLines(os.Args[1])
	towels, designs := parseLines(lines)
	result := calculatePart1(towels, designs)
	fmt.Printf("part1: %v\n", result)

	result2 := calculatePart2(towels, designs)
	fmt.Printf("part2: %v\n", result2)

	executionTime := time.Since(startTime)
	fmt.Printf("Execution time: %v\n", executionTime)
}

func readLines(filename string) []string {
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	return strings.Split(string(file), "\n\n")
}
