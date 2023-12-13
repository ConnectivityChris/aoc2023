package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	part1, part2 := calculateSolution("input.txt")
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}

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

var totalSteps = 0

func calculateSolution(filename string) (int, int) {

	// | is a vertical pipe connecting north and south.
	// - is a horizontal pipe connecting east and west.
	// L is a 90-degree bend connecting north and east.
	// J is a 90-degree bend connecting north and west.
	// 7 is a 90-degree bend connecting south and west.
	// F is a 90-degree bend connecting south and east.
	// . is ground; there is no pipe in this tile.
	// S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.

	// Find the starting point
	input, _ := readInput(filename)
	// fmt.Printf("Input: %v\n", input)
	currentIndex := [2]int{}
	previousIndex := [2]int{}
	// leftPathIndex := [2]int{}
	// rightPathIndex := [2]int{}
	for i, item := range input {
		if strings.Contains(item, "S") {
			index := strings.Index(item, "S")
			if index != -1 {
				result := [2]int{i, index}
				currentIndex = result
				// currentIndex = [2]int{1, 2}
				previousIndex = result
				// startIndex = result
				// leftPathIndex = result
				// rightPathIndex = result
			}
		}
	}
	fmt.Printf("Starting index: %v\n", currentIndex)
	fmt.Printf("Previous index: %v\n", previousIndex)

	foundStart := false
	for {
		if foundStart {
			break
		}
		nextDirections := [][2]int{}
		possibleDirections := checkDirectionOfTravel(input[currentIndex[0]][currentIndex[1]])
		for direction, validDirections := range possibleDirections {
			// fmt.Printf("Direction: %s, Pipe: %s\n", direction, validDirections)

			switch direction {
			case "UP":
				// Check upwards for a valid symbol
				neighbourY := currentIndex[0] - 1
				neighbourX := currentIndex[1]
				if neighbourY < 0 || neighbourY >= len(input) || neighbourX < 0 || neighbourX >= len(input[0]) {
					continue
				}
				neighbourPipe := string(input[neighbourY][neighbourX])
				// fmt.Printf("Symbol Up is: %s\n", neighbourPipe)
				// Check if neighbourPipe is in the lift of valid choices (pipe)
				if neighbourPipe == "S" && totalSteps > 1 {
					fmt.Printf("Found the start location\n")
					foundStart = true
					break
				}
				for _, validDirection := range validDirections {
					if neighbourPipe == validDirection && [2]int{neighbourY, neighbourX} != previousIndex {
						nextDirections = append(nextDirections, [2]int{neighbourY, neighbourX})
						break
					}
				}
			case "RIGHT":
				neighbourY := currentIndex[0]
				neighbourX := currentIndex[1] + 1
				if neighbourY < 0 || neighbourY >= len(input) || neighbourX < 0 || neighbourX >= len(input[0]) {
					continue
				}
				neighbourPipe := string(input[neighbourY][neighbourX])
				// fmt.Printf("Symbol Right is: %s\n", neighbourPipe)
				// Check if neighbourPipe is in the lift of valid choices (pipe)
				for _, validDirection := range validDirections {
					if neighbourPipe == validDirection && [2]int{neighbourY, neighbourX} != previousIndex {
						nextDirections = append(nextDirections, [2]int{neighbourY, neighbourX})
						break
					}
				}
			case "DOWN":
				neighbourY := currentIndex[0] + 1
				neighbourX := currentIndex[1]
				if neighbourY < 0 || neighbourY >= len(input) || neighbourX < 0 || neighbourX >= len(input[0]) {
					continue
				}
				neighbourPipe := string(input[neighbourY][neighbourX])
				// fmt.Printf("Symbol Down is: %s\n", neighbourPipe)
				// Check if neighbourPipe is in the lift of valid choices (pipe)
				for _, validDirection := range validDirections {
					if neighbourPipe == validDirection && [2]int{neighbourY, neighbourX} != previousIndex {
						nextDirections = append(nextDirections, [2]int{neighbourY, neighbourX})
						break
					}
				}
			case "LEFT":
				neighbourY := currentIndex[0]
				neighbourX := currentIndex[1] - 1
				if neighbourY < 0 || neighbourY >= len(input) || neighbourX < 0 || neighbourX >= len(input[0]) {
					continue
				}
				neighbourPipe := string(input[neighbourY][neighbourX])
				// fmt.Printf("Symbol Left is: %s\n", neighbourPipe)
				// Check if neighbourPipe is in the lift of valid choices (pipe)
				if neighbourPipe == "S" && totalSteps > 1 {
					fmt.Printf("Found the start location\n")
					break
				}
				for _, validDirection := range validDirections {
					if neighbourPipe == validDirection && [2]int{neighbourY, neighbourX} != previousIndex {
						nextDirections = append(nextDirections, [2]int{neighbourY, neighbourX})
						break
					}
				}
			}
		}
		fmt.Printf("Next Directions: %v\n", nextDirections)
		if len(nextDirections) == 2 || len(nextDirections) == 0 {
			if totalSteps > 0 {
				fmt.Printf("Completed Loop \n")
				fmt.Printf("Total Steps: %d\n", totalSteps+1)
				totalSteps++
				break
			} else {
				// First step, just move for now
				previousIndex = currentIndex
				currentIndex = nextDirections[0]
				totalSteps++
			}
		} else {
			previousIndex = currentIndex
			currentIndex = nextDirections[0]
			totalSteps++
		}
	}

	return totalSteps / 2, 0
}

// |
// UP if |, F, 7
// Cannot Travel Right
// DOWN if |, L, J
// Cannot Travel Left
// -
// Cannot Travel Up
// RIGHT if -, J, 7
// Cannot Travel Down
// LEFT if -. F, L
// L
// UP if |, F, 7
// RIGHT if -, J, 7
// Cannot Travel Down
// Cannot Travel Left
// J
// UP if |, F, 7
// Cannot Travel Right
// Cannot Travel Down
// LEFT if -. F, L
// 7
// Cannot Travel Up
// Cannot Travel Right
// DOWN if |, L, J
// LEFT if -. F, L
// F
// Cannot Travel Up
// RIGHT if -, J, 7
// DOWN if |, L, J
// Cannot Travel Left

func checkDirectionOfTravel(currentPipe byte) map[string][]string {
	validDirections := make(map[string][]string)
	switch currentPipe {
	case 'S':
		validDirections["UP"] = []string{"|", "F", "7"}
		validDirections["RIGHT"] = []string{"-", "J", "7"}
		validDirections["DOWN"] = []string{"|", "L", "J"}
		validDirections["LEFT"] = []string{"-", "F", "L"}
	case '|':
		validDirections["UP"] = []string{"|", "F", "7"}
		validDirections["DOWN"] = []string{"|", "L", "J"}
	case '-':
		validDirections["RIGHT"] = []string{"-", "J", "7"}
		validDirections["LEFT"] = []string{"-", "F", "L"}
	case 'L':
		validDirections["UP"] = []string{"|", "F", "7"}
		validDirections["RIGHT"] = []string{"-", "J", "7"}
	case 'J':
		validDirections["UP"] = []string{"|", "F", "7"}
		validDirections["LEFT"] = []string{"-", "F", "L"}
	case '7':
		validDirections["DOWN"] = []string{"|", "L", "J"}
		validDirections["LEFT"] = []string{"-", "F", "L"}
	case 'F':
		validDirections["RIGHT"] = []string{"-", "J", "7"}
		validDirections["DOWN"] = []string{"|", "L", "J"}
	}
	return validDirections
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
