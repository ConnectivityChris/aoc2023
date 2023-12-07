package main

import (
	"testing"
)

func TestWithExampleInput(t *testing.T) {
	solution := calculatePart1("../testinput.txt")

	if solution != 6440 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", solution, 6440)
	}
}
