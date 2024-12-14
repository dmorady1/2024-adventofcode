package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type position struct {
	row, col int
}

func parseLines(lines []string) []int {
	strNumbers := strings.Fields(lines[0])

	numbers := make([]int, 0, len(strNumbers))

	for _, str := range strNumbers {
		if num, err := strconv.Atoi(str); err == nil {
			numbers = append(numbers, num)
		}
	}

	return numbers

}

func blink(stones []int, count int) int {
	result := 0
	for _, number := range stones {
		var test []int
		test = append(test, number)

		var new_test []int
		for i := 0; i < count; i++ {
			for numIndex := 0; numIndex < len(test); numIndex++ {
				if test[numIndex] == 0 {
					new_test = append(new_test, 1)
					continue
				}
				numString := strconv.Itoa(test[numIndex])
				if len(numString)%2 == 0 {
					half := len(numString) / 2
					firstHalf, secondHalf := numString[:half], numString[half:]
					firstHalfInt, _ := strconv.Atoi(firstHalf)
					secondHalfInt, _ := strconv.Atoi(secondHalf)
					new_test = append(new_test, firstHalfInt)
					new_test = append(new_test, secondHalfInt)
					continue
				}
				new_test = append(new_test, test[numIndex]*2024)
			}
			result += len(new_test)
			new_test = nil
		}
	}
	return result

}

func calculatePart1(stones []int, count int) int {
	result := 0
	for index, number := range stones {
		test := []int{number}
		new_test := make([]int, 0, len(test)*2)

		fmt.Println(index, len(stones))
		for i := 0; i < count; i++ {
			for numIndex := 0; numIndex < len(test); numIndex++ {
				if test[numIndex] == 0 {
					new_test = append(new_test, 1)
					continue
				}
				numString := strconv.Itoa(test[numIndex])
				if len(numString)%2 == 0 {
					half := len(numString) / 2
					firstHalf, secondHalf := numString[:half], numString[half:]
					firstHalfInt, _ := strconv.Atoi(firstHalf)
					secondHalfInt, _ := strconv.Atoi(secondHalf)
					new_test = append(new_test, firstHalfInt)
					new_test = append(new_test, secondHalfInt)
					continue
				}
				new_test = append(new_test, test[numIndex]*2024)
			}
			test, new_test = new_test, test[:0]
		}
		result += len(test)
	}
	return result
}

type NumberDistance struct {
	number   int
	distance int
}

func generateMap(stones []int) map[NumberDistance]int {
	result := make(map[NumberDistance]int)
	for number := 0; number < 10; number++ {
		fmt.Println("pre calculate", number)
		_, intermediateResults := calculateArray([]int{number}, 40)

		for i := 1; i <= 40; i++ {
			result[NumberDistance{number, i}] = intermediateResults[i-1]
		}
	}
	return result
}

func calculateArray(stones []int, count int) ([][]int, []int) {
	var result [][]int
	var intermediateResults []int
	for _, number := range stones {
		fmt.Println(number, len(stones))
		test := []int{number}
		new_test := make([]int, 0, len(test)*2)
		intermediateResults = nil

		for i := 0; i < count; i++ {
			for numIndex := 0; numIndex < len(test); numIndex++ {
				if test[numIndex] == 0 {
					new_test = append(new_test, 1)
					continue
				}
				numString := strconv.Itoa(test[numIndex])
				if len(numString)%2 == 0 {
					half := len(numString) / 2
					firstHalf, secondHalf := numString[:half], numString[half:]
					firstHalfInt, _ := strconv.Atoi(firstHalf)
					secondHalfInt, _ := strconv.Atoi(secondHalf)
					new_test = append(new_test, firstHalfInt)
					new_test = append(new_test, secondHalfInt)
					continue
				}
				new_test = append(new_test, test[numIndex]*2024)
			}
			test, new_test = new_test, test[:0]
			intermediateResults = append(intermediateResults, len(test))
			if i == count-1 {
				result = append(result, test)
			}
		}
	}
	return result, intermediateResults

}

func flattenArrays(arrayAt35 [][]int) []int {
	var result []int
	for _, test := range arrayAt35 {
		for _, number := range test {
			result = append(result, number)
		}
	}
	return result
}

func findGroups(numbers []int) map[int]int {
	result := make(map[int]int)
	for _, number := range numbers {
		result[number] += 1
	}
	return result
}

func calculatePart2(stones []int, count int) int {
	preMap := generateMap(stones)

	arrayAt35, _ := calculateArray(stones, 35)
	flattenedArray := flattenArrays(arrayAt35)
	checkAt35 := calculatePart1(stones, 35)
	if len(flattenedArray) != checkAt35 {
		fmt.Println(len(flattenedArray))
		fmt.Println(checkAt35)
		fmt.Println("not equal")
		os.Exit(0)

	}
	groups := findGroups(flattenedArray)
	result := 0

	for number, times := range groups {
		result += preMap[NumberDistance{number, 40}] * times
	}

	var continueArray []int

	arrayAt35 = nil
	for _, number := range flattenedArray {
		_, ok := preMap[NumberDistance{number, 40}]
		if !ok {
			continueArray = append(continueArray, number)
		}
	}

	fmt.Println("at 35 interemediate result", result)
	fmt.Println("calculate Rest")
	fmt.Println(len(continueArray))
	result2 := calculatePart2New(continueArray, 40, preMap)

	return result + result2

}

func checkFilter(stones []int, preMap map[NumberDistance]int, distance int) ([]int, int) {
	count := 0
	var newStones []int
	for _, number := range stones {
		value, ok := preMap[NumberDistance{number, distance}]
		if ok {
			count += value
			continue
		}

		newStones = append(newStones, number)
	}
	return newStones, count

}

func calculatePart2New(stones []int, count int, preMap map[NumberDistance]int) int {
	result := 0
	for index, number := range stones {
		test := []int{number}
		new_test := make([]int, 0, len(test)*2)

		fmt.Println(index, len(stones))
		for i := 0; i < count; i++ {
			for numIndex := 0; numIndex < len(test); numIndex++ {
				if test[numIndex] == 0 {
					new_test = append(new_test, 1)
					continue
				}
				numString := strconv.Itoa(test[numIndex])
				if len(numString)%2 == 0 {
					half := len(numString) / 2
					firstHalf, secondHalf := numString[:half], numString[half:]
					firstHalfInt, _ := strconv.Atoi(firstHalf)
					secondHalfInt, _ := strconv.Atoi(secondHalf)
					new_test = append(new_test, firstHalfInt)
					new_test = append(new_test, secondHalfInt)
					continue
				}
				new_test = append(new_test, test[numIndex]*2024)
			}
			test, new_test = new_test, test[:0]
			var x int
			test, x = checkFilter(test, preMap, count-(i+1))
			result += x
		}
		result += len(test)
	}
	return result
}

func main() {
	lines := readLines(os.Args[1])
	stones := parseLines(lines)
	result := calculatePart1(stones, 25)
	fmt.Printf("part1: %d\n", result)

	result2 := calculatePart2(stones, 75)
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
