package main

import (
	"os"
	"testing"
)

func TestPart2WithExampleInput(t *testing.T) {
	file, _ := os.Open("../testinput.txt")

	defer file.Close()

	location := findLocation(file)

	if location != 46 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", location, 46)
	}
}

func TestPart2WithFullInput(t *testing.T) {
	file, _ := os.Open("../input.txt")

	defer file.Close()

	location := findLocation(file)

	if location != 81956384 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", location, 81956384)
	}
}
