package main

import (
	"testing"
)

func TestPart2WithExampleInput(t *testing.T) {
	solution := calculatePart2("../testinput.txt")

	if solution != 5905 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", solution, 5905)
	}
}

func TestPart2(t *testing.T) {
	solution := calculatePart2("../input.txt")

	if solution != 250665248 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", solution, 250665248)
	}
}

// 250665248
