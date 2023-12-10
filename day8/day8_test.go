package main

import (
	"testing"
)

func TestWithExampleInput(t *testing.T) {
	solution := calculateSteps("testinput.txt")

	if solution != 6 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", solution, 6)
	}
}

func TestWithFullInput(t *testing.T) {
	solution := calculateSteps("input.txt")

	if solution != 19951 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", solution, 19951)
	}
}

func TestPart2WithExampleInput(t *testing.T) {
	solution := calculateSteps("testinput2.txt")

	if solution != 6 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", solution, 6)
	}
}
func TestPart2WithFullInput(t *testing.T) {
	solution := calculateSteps("input.txt")

	if solution != 16342438708751 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", solution, 16342438708751)
	}
}
