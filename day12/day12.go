package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1, part2 := calculateSolution("input.txt")
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}

func calculateSolution(filename string) (int, int) {
	return part1(filename), part2(filename)
}

func part1(filename string) int {
	input, _ := readInput(filename)

	return runPermutations(input)
}

func part2(filename string) int {
	input, _ := readInput(filename)
	expandedInput := make([]string, 0)

	for _, line := range input {
		splitLine := strings.Split(line, " ")
		springs := strings.Repeat(splitLine[0]+"?", 4) + splitLine[0]
		config := strings.Repeat(splitLine[1]+",", 4) + splitLine[1]
		expandedInput = append(expandedInput, springs+" "+config)
	}

	return runPermutations(expandedInput)
}

func runPermutations(input []string) int {
	sum := 0

	for _, line := range input {
		splitLine := strings.Split(line, " ")
		springs := splitLine[0]
		config := splitAndConvertToInts(splitLine[1], ",")
		sum += generatePermutations(springs, config)
	}

	return sum
}

func generatePermutations(line string, config []int) int {
	cache := make([][]int, len(line))
	for i := range cache {
		cache[i] = make([]int, len(config)+1)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}

	return dynamicRecursion(0, 0, line, config, cache)
}

func dynamicRecursion(linePos, configPos int, line string, config []int, cache [][]int) int {
	if linePos >= len(line) {
		// end of line and config, no valid permutations
		if configPos < len(config) {
			return 0
		}
		// end of line, valid permutation
		return 1
	}

	// If there is a cache value, just use that and skip the recursion
	if cache[linePos][configPos] != -1 {
		return cache[linePos][configPos]
	}

	possiblePermutations := 0
	if line[linePos] == '.' {
		// Move along, nothing to see here
		possiblePermutations = dynamicRecursion(linePos+1, configPos, line, config, cache)
	} else {
		if line[linePos] == '?' {
			// Variation possible start recursion
			possiblePermutations += dynamicRecursion(linePos+1, configPos, line, config, cache)
		}
		if configPos < len(config) {
			count := 0
			for k := linePos; k < len(line); k++ {
				if count > config[configPos] || line[k] == '.' || count == config[configPos] && line[k] == '?' {
					break
				}
				count += 1
			}

			if count == config[configPos] {
				if linePos+count < len(line) && line[linePos+count] != '#' {
					possiblePermutations += dynamicRecursion(linePos+count+1, configPos+1, line, config, cache)
				} else {
					possiblePermutations += dynamicRecursion(linePos+count, configPos+1, line, config, cache)
				}
			}
		}
	}

	cache[linePos][configPos] = possiblePermutations
	return possiblePermutations
}

// func checkIfValidConfig(line string, config []int) bool {
// 	var counts []int
// 	count := 0

// 	for _, char := range line {
// 		if char == '#' {
// 			count++
// 		} else if count > 0 {
// 			counts = append(counts, count)
// 			count = 0
// 		}
// 	}

// 	// Add the last count if the line ends with '#'
// 	if count > 0 {
// 		counts = append(counts, count)
// 	}

// 	return reflect.DeepEqual(counts, config)
// }

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

func splitAndConvertToInts(input string, sep string) []int {
	parts := strings.Split(input, sep)
	config := make([]int, len(parts))
	for i, part := range parts {
		val, _ := strconv.Atoi(part)
		config[i] = val
	}
	return config
}
