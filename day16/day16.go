package main

import (
	"container/heap"
	"fmt"
	"os"
	"strings"
	"time"
)

type Vector struct {
	row int
	col int
}

func parseLines(splittedByNewLines []string) ([][]rune, map[Vector]bool, Vector, Vector, map[Vector]int) {
	walls := make(map[Vector]bool)
	dots := make(map[Vector]int)

	gridInput := splittedByNewLines[0]

	lines := strings.Split(gridInput, "\n")
	rows := len(lines)
	cols := len(lines[0])
	grid := make([][]rune, rows)
	for i := range grid {
		grid[i] = make([]rune, cols)
	}

	var startPosition Vector
	var endPosition Vector
	for row, line := range lines {
		for col, char := range line {
			grid[row][col] = char
			if char == '#' {
				walls[Vector{row, col}] = true
			}
			if char == 'S' {
				startPosition = Vector{row, col}
			}
			if char == 'E' {
				endPosition = Vector{row, col}
			}
			if char == '.' {
				dots[Vector{row, col}] = 0
			}

		}
	}

	return grid, walls, startPosition, endPosition, dots
}

func add(a Vector, b Vector) Vector {
	return Vector{a.row + b.row, a.col + b.col}
}

// An Item is something we manage in a priority queue.

type Item struct {
	value    VectorDir // The value of the item; arbitrary.
	priority int       // The priority of the item in the queue.
	index    int       // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value VectorDir, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

type VectorDist struct {
	vector Vector
	dist   int
}

type VectorDir struct {
	vector Vector
	dir    int
}

func calculateAdjancencyList(grid [][]rune, walls map[Vector]bool, startPosition Vector, endPosition Vector) map[Vector][]VectorDist {
	vectorAdjacencyList := make(map[Vector][]VectorDist)

	directions := []Vector{
		{-1, 0}, // Up
		{1, 0},  // Down
		{0, -1}, // Left
		{0, 1},  // Right
	}

	for row := range grid {
		for col := range grid[row] {
			currentPos := Vector{row, col}
			if walls[currentPos] {
				continue
			}

			var adjacent []VectorDist
			for _, dir := range directions {
				nextPos := add(currentPos, dir)
				if !walls[nextPos] {
					adjacent = append(adjacent, VectorDist{nextPos, 1})
				}
			}

			vectorAdjacencyList[currentPos] = adjacent
		}
	}

	return vectorAdjacencyList
}

func calculatePart1(grid [][]rune, walls map[Vector]bool, startPosition Vector, endPosition Vector, dots map[Vector]int) int {
	directions := make(map[rune]Vector)
	directions['^'] = Vector{-1, 0}
	directions['v'] = Vector{1, 0}
	directions['>'] = Vector{0, 1}
	directions['<'] = Vector{0, -1}

	vectorToDirectionRune := make(map[Vector]rune)
	vectorToDirectionRune[Vector{-1, 0}] = '^'
	vectorToDirectionRune[Vector{1, 0}] = 'v'
	vectorToDirectionRune[Vector{0, 1}] = '>'
	vectorToDirectionRune[Vector{0, -1}] = '<'

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	adjacencyList := calculateAdjancencyList(grid, walls, startPosition, endPosition)

	distance, path := dijkstra(pq, adjacencyList, startPosition, endPosition)

	// Visualize the path
	for i := 0; i < len(path)-1; i++ {
		curr := path[i]
		next := path[i+1]
		direction := Vector{next.row - curr.row, next.col - curr.col}
		char := vectorToDirectionRune[direction]
		grid[next.row][next.col] = char
	}

	printGrid(grid)

	return distance
}

func dijkstra(pq PriorityQueue, adjacencyList map[Vector][]VectorDist, startPosition Vector, endPosition Vector) (int, []Vector) {
	dist := make(map[VectorDir]int)
	prev := make(map[VectorDir]VectorDir)
	visited := make(map[VectorDir]bool)

	// Initialize distances
	for vertex := range adjacencyList {
		for dir := 0; dir < 4; dir++ {
			dist[VectorDir{vertex, dir}] = 99999999
		}
	}

	// initial direction right
	startDir := 2
	dist[VectorDir{startPosition, startDir}] = 0

	heap.Push(&pq, &Item{value: VectorDir{startPosition, startDir}, priority: 0})

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		current := item.value

		if current.vector == endPosition {
			break
		}

		if visited[current] {
			continue
		}
		visited[current] = true

		for _, neighbor := range adjacencyList[current.vector] {
			newDir := getDirection(current.vector, neighbor.vector)
			turningCost := 0
			if current.dir != newDir {
				turningCost = 1000
			}

			newDist := dist[current] + neighbor.dist + turningCost
			newState := VectorDir{neighbor.vector, newDir}

			if newDist < dist[newState] {
				dist[newState] = newDist
				prev[newState] = current
				heap.Push(&pq, &Item{
					value:    newState,
					priority: newDist,
				})
			}
		}
	}

	// Find minimum distance to end position
	minDist := 99999999
	var endState VectorDir
	for dir := 0; dir < 4; dir++ {
		endStateCandidate := VectorDir{endPosition, dir}
		if dist[endStateCandidate] < minDist {
			minDist = dist[endStateCandidate]
			endState = endStateCandidate
		}
	}

	// Reconstruct path
	path := []Vector{endPosition}
	curr := endState
	for curr.vector != startPosition {
		path = append([]Vector{curr.vector}, path...)
		curr = prev[curr]
	}

	return minDist, path
}

func getDirection(from, to Vector) int {
	diff := Vector{to.row - from.row, to.col - from.col}
	switch diff {
	case Vector{-1, 0}:
		return 0 // Up
	case Vector{1, 0}:
		return 1 // Down
	case Vector{0, -1}:
		return 2 // Left
	case Vector{0, 1}:
		return 3 // Right
	}
	return -1 // Invalid direction
}

func calculatePart2(grid [][]rune, walls map[Vector]bool, startPosition Vector, endPosition Vector, dots map[Vector]int) int {
	directions := make(map[rune]Vector)
	directions['^'] = Vector{-1, 0}
	directions['v'] = Vector{1, 0}
	directions['>'] = Vector{0, 1}
	directions['<'] = Vector{0, -1}

	vectorToDirectionRune := make(map[Vector]rune)
	vectorToDirectionRune[Vector{-1, 0}] = '^'
	vectorToDirectionRune[Vector{1, 0}] = 'v'
	vectorToDirectionRune[Vector{0, 1}] = '>'
	vectorToDirectionRune[Vector{0, -1}] = '<'

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	adjacencyList := calculateAdjancencyList(grid, walls, startPosition, endPosition)

	result := dijkstra2(pq, adjacencyList, startPosition, endPosition)

	return result
}

func dijkstra2(pq PriorityQueue, adjacencyList map[Vector][]VectorDist, startPosition Vector, endPosition Vector) int {
	dist := make(map[VectorDir]int)
	prev := make(map[VectorDir][]VectorDir)
	visited := make(map[VectorDir]bool)

	// Initialize distances
	for vertex := range adjacencyList {
		for dir := 0; dir < 4; dir++ {
			dist[VectorDir{vertex, dir}] = 99999999
		}
	}

	// initial direction right
	startDir := 2
	dist[VectorDir{startPosition, startDir}] = 0

	heap.Push(&pq, &Item{value: VectorDir{startPosition, startDir}, priority: 0})

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		current := item.value

		if visited[current] {
			continue
		}
		visited[current] = true

		for _, neighbor := range adjacencyList[current.vector] {
			newDir := getDirection(current.vector, neighbor.vector)
			turningCost := 0
			if current.dir != newDir {
				turningCost = 1000
			}

			newDist := dist[current] + neighbor.dist + turningCost
			newState := VectorDir{neighbor.vector, newDir}

			if newDist <= dist[newState] {
				if newDist < dist[newState] {
					dist[newState] = newDist
					prev[newState] = []VectorDir{current}
				} else {
					prev[newState] = append(prev[newState], current)
				}
				heap.Push(&pq, &Item{
					value:    newState,
					priority: newDist,
				})
			}
		}
	}

	// Find minimum distance to end position
	minDist := 99999999
	var endStates []VectorDir
	for dir := 0; dir < 4; dir++ {
		endStateCandidate := VectorDir{endPosition, dir}
		if dist[endStateCandidate] < minDist {
			minDist = dist[endStateCandidate]
			endStates = []VectorDir{endStateCandidate}
		} else if dist[endStateCandidate] == minDist {
			endStates = append(endStates, endStateCandidate)
		}
	}

	// Reconstruct paths
	bestTiles := make(map[Vector]bool)
	bestTiles[endPosition] = true

	var dfs func(VectorDir)
	dfs = func(state VectorDir) {
		if state.vector == startPosition {
			return
		}
		for _, prevState := range prev[state] {
			bestTiles[prevState.vector] = true
			dfs(prevState)
		}
	}

	for _, endState := range endStates {
		dfs(endState)
	}

	return len(bestTiles)
}
func resetGrid(grid [][]rune) [][]rune {
	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}
	return grid
}

// func setGrid(grid [][]rune, positions []Vector) [][]rune {
// 	for _, vector := range positions {
// 		grid[vector.y][vector.x] = '#'
// 	}
// 	return grid
//
// }
//

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
	grid, walls, startPosition, endPosition, dots := parseLines(lines)

	result := calculatePart1(grid, walls, startPosition, endPosition, dots)
	fmt.Printf("part1: %d\n", result)

	result2 := calculatePart2(grid, walls, startPosition, endPosition, dots)
	fmt.Printf("part2: %d\n", result2)

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
