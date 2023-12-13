package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

type Point struct {
	X int
	Y int
}

var coordinateChecks = []Point{
	{0, -1}, // Up
	{1, 0},  // Right
	{0, 1},  // Down
	{-1, 0}, // Left
}

func main() {
	part1 := calculateSolution("../input.txt")
	fmt.Printf("Part 1: %d\n", part1)
}

func calculateSolution(filename string) int {
	input, _ := readInput(filename)
	expandedGalaxy := expandGalaxy(input)
	sum := 0
	// for _, str := range expandedGalaxy {
	// 	fmt.Printf("%s\n", str)
	// }
	// Find all '#' positions
	var hashPositions []Point
	for y, row := range expandedGalaxy {
		for x, char := range row {
			if char == '#' {
				hashPositions = append(hashPositions, Point{x, y})
			}
		}
	}

	// Channel to collect BFS results
	results := make(chan int)
	// WaitGroup to wait for all goroutines to complete
	var wg sync.WaitGroup

	// fmt.Printf("Hash positions: %v\n", hashPositions)
	// Run BFS in parallel for each pair
	for i := 0; i < len(hashPositions)-1; i++ {
		for j := i + 1; j < len(hashPositions); j++ {
			wg.Add(1) // Increment the WaitGroup counter
			go func(start, end Point) {
				defer wg.Done() // Decrement the counter when the goroutine completes
				distance := breadthFirstSearch(expandedGalaxy, start, end)
				results <- distance // Send the result to the channel
			}(hashPositions[i], hashPositions[j])
		}
	}
	// Close the results channel once all goroutines are done
	go func() {
		wg.Wait()
		close(results)
	}()
	for distance := range results {
		sum += distance
	}
	return sum
}

func breadthFirstSearch(expandedGalaxy []string, start Point, end Point) int {
	// Queue for BFS
	queue := []Point{start}

	// Map to track visited points and their distances from the start
	visited := make(map[Point]int)
	visited[start] = 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		// Check if we have reached the destination
		if current == end {
			return visited[current]
		}
		// Explore neighboring cells
		for _, direction := range coordinateChecks {
			next := Point{X: current.X + direction.X, Y: current.Y + direction.Y}
			// Check if the next point is within bounds and not visited
			if next.X >= 0 && next.Y >= 0 && next.Y < len(expandedGalaxy) && next.X < len(expandedGalaxy[next.Y]) &&
				(expandedGalaxy[next.Y][next.X] == '.' || expandedGalaxy[next.Y][next.X] == '#') && visited[next] == 0 {

				queue = append(queue, next)
				visited[next] = visited[current] + 1
			}
		}
	}

	return -1
}

func expandGalaxy(input []string) []string {
	for i := 0; i < len(input); i++ {
		if strings.Count(input[i], ".") == len(input[i]) {
			input = append(input[:i+1], input[i:]...)
			i++ // Skip the newly added row
		}
	}

	for col := 0; col < len(input[0]); col++ {
		allDots := true
		for row := 0; row < len(input); row++ {
			if input[row][col] != '.' {
				allDots = false
				break
			}
		}
		if allDots {
			for row := range input {
				input[row] = input[row][:col+1] + string(input[row][col]) + input[row][col+1:]
			}
			col++ // Skip the newly added column
		}
	}
	return input
}

func readInput(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(content), "\n"), nil
}
