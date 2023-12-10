package main

import "testing"

func TestWithExampleInput(t *testing.T) {
	part1, part2 := calculateSolution("testinput.txt")

	if part1 != 114 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", part1, 114)
	}

	if part2 != 2 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", part1, 2)
	}
}
