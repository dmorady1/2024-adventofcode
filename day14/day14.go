package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Vector struct {
	x int
	y int
}

func parseLines(lines []string) ([]Vector, []Vector, error) {
	startPositions := make([]Vector, 0, len(lines))
	velocities := make([]Vector, 0, len(lines))

	for i, line := range lines {
		p, v, err := parseLine(line)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing line %d: %v", i+1, err)
		}
		startPositions = append(startPositions, Vector{p[0], p[1]})
		velocities = append(velocities, Vector{v[0], v[1]})
	}

	return startPositions, velocities, nil
}

func parseLine(line string) ([2]int, [2]int, error) {
	var p, v [2]int

	// Split the line into two parts
	parts := strings.Split(line, " ")
	if len(parts) != 2 {
		return p, v, fmt.Errorf("invalid line format")
	}

	// Parse p values
	pPart := strings.TrimPrefix(parts[0], "p=")
	pValues := strings.Split(pPart, ",")
	if len(pValues) != 2 {
		return p, v, fmt.Errorf("invalid p format")
	}
	for i, val := range pValues {
		num, err := strconv.Atoi(val)
		if err != nil {
			return p, v, fmt.Errorf("error parsing p value: %v", err)
		}
		p[i] = num
	}

	// Parse v values
	vPart := strings.TrimPrefix(parts[1], "v=")
	vValues := strings.Split(vPart, ",")
	if len(vValues) != 2 {
		return p, v, fmt.Errorf("invalid v format")
	}
	for i, val := range vValues {
		num, err := strconv.Atoi(val)
		if err != nil {
			return p, v, fmt.Errorf("error parsing v value: %v", err)
		}
		v[i] = num
	}

	return p, v, nil
}

func scalarMult(a Vector, num int) Vector {
	return Vector{a.x * num, a.y * num}
}

func add(a Vector, b Vector) Vector {
	return Vector{a.x + b.x, a.y + b.y}
}

func mod(a, b int) int {
	result := a % b
	if result < 0 {
		for {
			if result > 0 {
				return result
			}
			result += b
		}
	}
	return result
}

func modulo(a Vector, wide int, tall int) Vector {
	return Vector{mod(a.x, wide), mod(a.y, tall)}
}

func calculatePart1(startPositions []Vector, velocities []Vector) int {
	topLeftQuarter := 0
	topRightQuarter := 0

	bottomLeftQuarter := 0
	bottomRightQuarter := 0

	wide := 101
	tall := 103
	times := 100

	halfWide := (wide / 2)
	halfTall := (tall / 2)
	for i := 0; i < len(startPositions); i++ {
		start := startPositions[i]
		v := velocities[i]

		result := modulo(add(start, scalarMult(v, times)), wide, tall)

		if result.x == halfWide || result.y == halfTall {
			continue
		}
		if result.x < halfWide {
			// left
			if result.y < halfTall {
				topLeftQuarter++
			} else {
				bottomLeftQuarter++
			}
		} else {
			// right
			if result.y < halfTall {
				topRightQuarter++
			} else {
				bottomRightQuarter++
			}
		}

	}

	return topLeftQuarter * bottomLeftQuarter * topRightQuarter * bottomRightQuarter
}

func resetGrid(grid [][]rune) [][]rune {
	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = ' '
		}
	}
	return grid
}

func setGrid(grid [][]rune, positions []Vector) [][]rune {
	for _, vector := range positions {
		grid[vector.y][vector.x] = '#'
	}
	return grid

}

func calculatePart2(startPositions []Vector, velocities []Vector) uint64 {
	var topLeftQuarter uint64 = 0
	var topRightQuarter uint64 = 0
	var bottomLeftQuarter uint64 = 0
	var bottomRightQuarter uint64 = 0

	wide := 101
	tall := 103

	halfWide := (wide / 2)
	halfTall := (tall / 2)

	grid := make([][]rune, tall)
	for i := range grid {
		grid[i] = make([]rune, wide)
		for j := range grid[i] {
			grid[i][j] = ' '
		}
	}

	reader := bufio.NewReader(os.Stdin)

	count := 0
	for {
		topLeftQuarter = 0
		topRightQuarter = 0
		bottomLeftQuarter = 0
		bottomRightQuarter = 0

		for i := 0; i < len(startPositions); i++ {
			start := startPositions[i]
			v := velocities[i]

			result := modulo(add(start, v), wide, tall)
			startPositions[i] = result

			if result.x == halfWide || result.y == halfTall {
				continue
			}
			if result.x < halfWide {
				// left
				if result.y < halfTall {
					topLeftQuarter++
				} else {
					bottomLeftQuarter++
				}
			} else {
				// right
				if result.y < halfTall {
					topRightQuarter++
				} else {
					bottomRightQuarter++
				}
			}
		}

		count++
		hasLine := hasContinuousLine(startPositions, wide, 15)

		if hasLine {
			grid = setGrid(resetGrid(grid), startPositions)
			printGrid(grid)

			fmt.Println("WWWWWWWWWWWWWWWWW")
			fmt.Printf("count: %d\n", count)
			fmt.Printf("Result %d \n", topLeftQuarter*bottomLeftQuarter*topRightQuarter*bottomRightQuarter)
			fmt.Println("WWWWWWWWWWWWWWWWW")
			fmt.Println("WWWWWWWWWWWWWWWWW")

			fmt.Println("Press Enter to continue or 'f' to finish...")
			input, _ := reader.ReadString('\n')
			if strings.TrimSpace(input) == "f" {
				break
			}
		}
	}

	return topLeftQuarter * bottomLeftQuarter * topRightQuarter * bottomRightQuarter
}

func hasContinuousLine(positions []Vector, width int, minLength int) bool {
	for y := 0; y < width; y++ {
		row := make([]bool, width)
		for _, pos := range positions {
			if pos.y == y {
				row[pos.x] = true
			}
		}

		count := 0
		for _, occupied := range row {
			if occupied {
				count++
				if count > minLength {
					return true
				}
			} else {
				count = 0
			}
		}
	}
	return false
}

// func calculatePart2(startPositions []Vector, velocities []Vector) uint64 {
// 	var topLeftQuarter uint64 = 0
// 	var topRightQuarter uint64 = 0
// 	var bottomLeftQuarter uint64 = 0
// 	var bottomRightQuarter uint64 = 0
//
// 	wide := 101
// 	tall := 103
//
// 	halfWide := (wide / 2)
// 	halfTall := (tall / 2)
//
// 	grid := make([][]rune, tall)
// 	for i := range grid {
// 		grid[i] = make([]rune, wide)
// 		for j := range grid[i] {
// 			grid[i][j] = '.'
// 		}
// 	}
//
// 	reader := bufio.NewReader(os.Stdin)
// 	stopChan := make(chan struct{})
// 	pauseChan := make(chan bool)
//
// 	isPaused := false
//
// 	// Goroutine to check for Enter and 'f' key presses
// 	go func() {
// 		for {
// 			char, _, err := reader.ReadRune()
// 			if err != nil {
// 				fmt.Println("Error reading input:", err)
// 				close(stopChan)
// 				return
// 			}
// 			if char == '\n' {
// 				if isPaused {
// 					fmt.Println("Resuming simulation...")
// 					pauseChan <- false
// 				} else {
// 					fmt.Println("Simulation paused. Press Enter to resume or 'f' to stop...")
// 					pauseChan <- true
// 				}
// 			} else if char == 'f' && isPaused {
// 				fmt.Println("'f' pressed, stopping simulation...")
// 				close(stopChan)
// 				return
// 			}
// 		}
// 	}()
//
// 	count := 0
// 	for {
// 		select {
// 		case <-stopChan:
// 			fmt.Println("Stopping simulation...")
// 			return topLeftQuarter * bottomLeftQuarter * topRightQuarter * bottomRightQuarter
// 		case pause := <-pauseChan:
// 			isPaused = pause
// 		default:
// 			if isPaused {
// 				time.Sleep(100 * time.Millisecond)
// 				continue
// 			}
// 			topLeftQuarter = 0
// 			topRightQuarter = 0
// 			bottomLeftQuarter = 0
// 			bottomRightQuarter = 0
//
// 			for i := 0; i < len(startPositions); i++ {
// 				start := startPositions[i]
// 				v := velocities[i]
//
// 				result := modulo(add(start, v), wide, tall)
// 				startPositions[i] = result
//
// 				if result.x == halfWide || result.y == halfTall {
// 					continue
// 				}
// 				if result.x < halfWide {
// 					// left
// 					if result.y < halfTall {
// 						topLeftQuarter++
// 					} else {
// 						bottomLeftQuarter++
// 					}
// 				} else {
// 					// right
// 					if result.y < halfTall {
// 						topRightQuarter++
// 					} else {
// 						bottomRightQuarter++
// 					}
// 				}
// 			}
//
// 			grid := setGrid(grid, startPositions)
// 			printGrid(grid)
// 			count++
//
// 			fmt.Println("WWWWWWWWWWWWWWWWW")
// 			fmt.Printf("count: %d\n", count)
// 			fmt.Printf("Result %d \n", topLeftQuarter*bottomLeftQuarter*topRightQuarter*bottomRightQuarter)
// 			fmt.Println("WWWWWWWWWWWWWWWWW")
// 			fmt.Println("WWWWWWWWWWWWWWWWW")
//
// 			resetGrid(grid)
//
// 			time.Sleep(500 * time.Millisecond) // Sleep for 0.5 seconds
// 		}
// 	}
// }

func printGrid(grid [][]rune) {

	var builder strings.Builder
	for _, row := range grid {
		builder.WriteString(string(row))
		builder.WriteString("\n")
	}
	result := builder.String()
	fmt.Print(result)

}

func main() {
	startTime := time.Now()

	lines := readLines(os.Args[1])
	startPositions, velocities, _ := parseLines(lines)

	result := calculatePart1(startPositions, velocities)
	fmt.Printf("part1: %d\n", result)

	result2 := calculatePart2(startPositions, velocities)
	fmt.Printf("part2: %d\n", result2)
	//
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
