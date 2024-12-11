package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseLines(lines []string) []int {
	fileID := 0

	line := lines[0]

	var result []int

	for i := 0; i < len(line); i++ {
		value, _ := strconv.Atoi(string(line[i]))
		if i%2 == 0 {
			for n := 0; n < value; n++ {
				result = append(result, fileID)

			}
			fileID++
		} else {
			for n := 0; n < value; n++ {
				result = append(result, -1)

			}
		}
	}
	return result
}

func main() {
	lines := readLines(os.Args[1])
	line := parseLines(lines)

	line2 := make([]int, len(line))
	copy(line2, line)

	fmt.Println(line)
	result := calculatePart1(line)
	fmt.Printf("part1: %d\n", result)

	result2 := calculatePart2(line2)
	fmt.Printf("part2: %d\n", result2)
}

func calculateResult(line []int) int {
	result := 0
	for i, value := range line {
		if value == -1 {
			break
		}
		result += i * value
	}
	return result
}

func calculatePart1(line []int) int {
	left := 0
	right := len(line) - 1

	for left < right {
		// Find next empty space from left
		for left < right && line[left] != -1 {
			left++
		}

		// Find next file from right
		for left < right && line[right] == -1 {
			right--
		}

		// Swap if we found both empty space and file
		if left < right {
			line[left], line[right] = line[right], line[left]
			right--
		}
	}

	return calculateResult(line)
}

type Space struct {
	startIndex int
	endIndex   int
	length     int
}

func getPositionsOfSpaceAndNoneSpaces(line []int) ([]Space, []Space) {
	var gaps []Space
	var files []Space
	var startIndex int
	var endIndex int
	var length int
	var prevNumber int

	for i, num := range line {
		if i == 0 {
			startIndex = 0
			prevNumber = num
			length = 0
		}
		if prevNumber == num {
			endIndex = i
			length++
		}
		if prevNumber != num {
			endIndex = i - 1
			space := Space{startIndex: startIndex, endIndex: endIndex, length: length}
			// fmt.Println(prevNumber)
			if prevNumber == -1 {
				gaps = append(gaps, space)

			} else {
				files = append(files, space)
			}
			startIndex = i
			length = 1
			endIndex = i

		}

		prevNumber = num
	}

	space := Space{startIndex: startIndex, endIndex: endIndex, length: length}

	if prevNumber == -1 {
		gaps = append(gaps, space)

	} else {
		files = append(files, space)
	}
	return files, gaps
}

func swap(line []int, file Space, gap Space) []int {
	for i := 0; i < file.length; i++ {
		line[gap.startIndex+i], line[file.startIndex+i] = line[file.startIndex+i], line[gap.startIndex+i]
	}
	return line
}

func calculateResult2(line []int) int {
	result := 0
	for i, value := range line {
		if value == -1 {
			continue
		}
		result += i * value
	}
	return result
}

func calculatePart2(line []int) int {
	files, gaps := getPositionsOfSpaceAndNoneSpaces(line)

	for i := len(files) - 1; i >= 0; i-- {
		file := files[i]
		for j := 0; j < len(gaps); j++ {
			gap := gaps[j]
			if gap.length >= file.length && gap.endIndex < file.startIndex {
				line = swap(line, file, gap)

				if gap.length == file.length {
					gaps = append(gaps[:j], gaps[j+1:]...)
					break
				} else {
					gap.startIndex = gap.startIndex + file.length
					gap.length = gap.length - file.length
				}
				gaps[j] = gap
				break

			}
		}
	}

	var builder strings.Builder
	for _, num := range line {
		if num == -1 {

			builder.WriteString(".")
		} else {

			builder.WriteString(strconv.Itoa(num))
		}
	}

	return calculateResult2(line)

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
