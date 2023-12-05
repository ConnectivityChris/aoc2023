package main

import (
	"os"
	"testing"
)

func TestPart1WithExampleInput(t *testing.T) {
	file, _ := os.Open("testinput.txt")

	defer file.Close()

	part1Location := findLocation(file, 1)
	// part2Location := findLocation(file, 2)

	if part1Location != 35 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", part1Location, 35)
	}
	// if part2Location != 46 {
	// 	t.Errorf("Result was incorrect, got: %d, want: %d.", part2Location, 46)
	// }
}

func TestPart2WithExampleInput(t *testing.T) {
	file, _ := os.Open("testinput.txt")

	defer file.Close()

	part1Location := findLocation(file, 2)

	if part1Location != 46 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", part1Location, 46)
	}
}

func TestPart1WithFullInput(t *testing.T) {
	file, _ := os.Open("input.txt")

	defer file.Close()

	part1Location := findLocation(file, 1)

	if part1Location != 218513636 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", part1Location, 218513636)
	}
}
