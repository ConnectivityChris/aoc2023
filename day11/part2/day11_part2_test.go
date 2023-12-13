package main

import "testing"

func TestWithExampleInput(t *testing.T) {
	solution := calculateSolution("../testinput.txt", 99)

	if solution != 8410 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", solution, 8410)
	}
}

func TestWithExampleInput2(t *testing.T) {
	solution := calculateSolution("../testinput.txt", 9)

	if solution != 1030 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", solution, 1030)
	}
}

func TestWithFullInput(t *testing.T) {
	solution := calculateSolution("../input.txt", 999999)

	if solution != 904633799472 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", solution, 904633799472)
	}
}
