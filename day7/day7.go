package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseLines(lines []string) ([]int, [][]int) {
	var testValues []int
	var values [][]int
	for _, line := range lines {
		parts := strings.Split(line, ":")
		var numbers []int
		testValue, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		testValues = append(testValues, testValue)
		for _, number := range strings.Split(strings.TrimSpace(parts[1]), " ") {
			number, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}
			numbers = append(numbers, number)
		}
		values = append(values, numbers)

	}

	return testValues, values
}
func main() {
	lines := readLines(os.Args[1])

	testValues, values := parseLines(lines)

	result := calculatePart1(testValues, values)
	fmt.Printf("part1: %d\n", result)

	// result2 := calculatePart2(string(data))
	// fmt.Printf("part2: %d\n", result2)
}

func search(testValue int, numbers []int) int {
	if len(numbers) == 1 && numbers[0] == testValue {
		fmt.Println("success")
		fmt.Println(testValue, numbers)
		return testValue
	}
	if len(numbers) == 1 && numbers[0] != testValue {
		return 0
	}

	numbersAdded := numbers[0] + numbers[1]
	numbersMul := numbers[0] * numbers[1]

	newNumbersAdded := append([]int{numbersAdded}, numbers[2:]...)
	newNumbersMul := append([]int{numbersMul}, numbers[2:]...)
	valueAdd := search(testValue, newNumbersAdded)
	if valueAdd == testValue {
		return testValue
	}

	valueMul := search(testValue, newNumbersMul)

	if valueMul == testValue {
		return testValue
	}
	return 0
}

func calculatePart1(testValues []int, values [][]int) int {
	result := 0

	for i := 0; i < len(testValues); i++ {
		testValue := testValues[i]
		numbers := values[i]

		if search(testValue, numbers) == testValue {
			result += testValue
		}

	}

	return result
}

func calculatePart2(text string) int {
	return 0
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
