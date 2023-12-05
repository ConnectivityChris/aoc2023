package main

import (
	"os"
	"testing"
)

func TestPart1WithExampleInput(t *testing.T) {
	file, _ := os.Open("testinput.txt")

	defer file.Close()

	location := findLocation(file, 1)
	// part2Location := findLocation(file, 2)

	if location != 35 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", location, 35)
	}
	// if part2Location != 46 {
	// 	t.Errorf("Result was incorrect, got: %d, want: %d.", part2Location, 46)
	// }
}

func TestPart2WithExampleInput(t *testing.T) {
	file, _ := os.Open("testinput.txt")

	defer file.Close()

	location := findLocation(file, 2)

	if location != 46 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", location, 46)
	}
}

func TestPart1WithFullInput(t *testing.T) {
	file, _ := os.Open("input.txt")

	defer file.Close()

	location := findLocation(file, 1)

	if location != 218513636 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", location, 218513636)
	}
}

func TestPart2WithFullInput(t *testing.T) {
	file, _ := os.Open("input.txt")

	defer file.Close()

	location := findLocation(file, 2)

	if location != 81956384 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", location, 81956384)
	}
}
