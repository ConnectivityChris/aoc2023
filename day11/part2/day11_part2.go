package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Point struct {
	X, Y int
}

func main() {
	part1 := calculateSolution("../input.txt", 999999)

	fmt.Printf("Total distance: %d\n", part1)
}

func calculateSolution(filename string, expandBy int) int {
	grid, _ := readInput(filename)

	// duplicates := 999999
	hashPositions := findHashPositions(grid)
	// fmt.Printf("Hash positions %v\n", hashPositions)
	effectiveHashPositions := make([]Point, len(hashPositions))
	for i, p := range hashPositions {
		effectiveHashPositions[i] = calculateEffectiveCoordinates(p, grid, expandBy)
	}
	// fmt.Printf("Effective Hash positions %v\n", effectiveHashPositions)

	sum := 0
	for i := 0; i < len(effectiveHashPositions)-1; i++ {
		for j := i + 1; j < len(effectiveHashPositions); j++ {
			distance := optimizedDistance(effectiveHashPositions[i], effectiveHashPositions[j])
			sum += distance
			// fmt.Printf("Distance from %v to %v: %d\n", effectiveHashPositions[i], effectiveHashPositions[j], distance)
		}
	}
	return sum
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

func findHashPositions(grid []string) []Point {
	var positions []Point
	for y, row := range grid {
		for x, char := range row {
			if char == '#' {
				positions = append(positions, Point{x, y})
			}
		}
	}
	return positions
}

func calculateEffectiveCoordinates(point Point, grid []string, duplicates int) Point {
	effectiveX, effectiveY := point.X, point.Y
	for y := 0; y < point.Y; y++ {
		if isDuplicatableRow(grid[y]) {
			effectiveY += duplicates
		}
	}
	for x := 0; x < point.X; x++ {
		if isDuplicatableColumn(grid, x) {
			effectiveX += duplicates
		}
	}
	return Point{effectiveX, effectiveY}
}

func isDuplicatableRow(row string) bool {
	return strings.Count(row, ".") == len(row)
}

func isDuplicatableColumn(grid []string, colIndex int) bool {
	for _, row := range grid {
		if row[colIndex] != '.' {
			return false
		}
	}
	return true
}

func optimizedDistance(start, end Point) int {
	return abs(start.X-end.X) + abs(start.Y-end.Y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
