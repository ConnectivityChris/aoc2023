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
	fmt.Printf("Part 1: %d\n", part2)
}

func calculateSolution(filename string) (int, int) {
	input, _ := readInput(filename)
	total := 0
	backwardsTotal := 0
	// extrapolatedNumbers := []int{}
	for _, str := range input {
		diffs := [][]string{}
		fmt.Printf("str: %s\n", str)
		values := strings.Fields(str)
		lastDigit, _ := strconv.Atoi(values[len(values)-1])
		firstDigit, _ := strconv.Atoi(values[0])
		for {
			if len(diffs) > 0 && allZeros(diffs[len(diffs)-1]) {
				break
			}
			currentLineDiffs := buildDiff(values)
			values = currentLineDiffs
			diffs = append(diffs, currentLineDiffs)
		}
		// reverse up the diffs and extrapolate the last number
		currentExtrapolated := 0
		backwardsExtrapolated := 0
		for i := len(diffs) - 1; i >= 0; i-- {
			lastNumber, _ := strconv.Atoi(diffs[i][len(diffs[i])-1])
			firstNumber, _ := strconv.Atoi(diffs[i][0])
			fmt.Printf("firstNumber: %d\n", firstNumber)
			diff := lastNumber + currentExtrapolated
			backwardsDiff := firstNumber - backwardsExtrapolated
			fmt.Printf("Backwards diff: %d\n", backwardsDiff)
			// extrapolate the next number
			currentExtrapolated = diff
			backwardsExtrapolated = backwardsDiff
		}

		total += lastDigit + currentExtrapolated
		backwardsTotal += firstDigit - backwardsExtrapolated
	}
	return total, backwardsTotal
}

func buildDiff(line []string) []string {
	currentLineDiffs := []string{}
	for i, char := range line[:len(line)-1] {
		// find the difference between the current and next number
		currentNumber, _ := strconv.Atoi(char)
		nextNumber, _ := strconv.Atoi(line[i+1])
		diff := fmt.Sprint((nextNumber - currentNumber))
		currentLineDiffs = append(currentLineDiffs, diff)
	}
	return currentLineDiffs
}

func allZeros(arr []string) bool {
	if len(arr) == 0 {
		return false
	}

	for _, value := range arr {
		if value != "0" {
			return false
		}
	}
	return true
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
