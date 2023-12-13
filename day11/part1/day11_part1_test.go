package main

import "testing"

func TestWithExampleInput(t *testing.T) {
	part1 := calculateSolution("../testinput.txt")

	if part1 != 374 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", part1, 374)
	}

	// if part2 != 2 {
	// 	t.Errorf("Result was incorrect, got: %d, want: %d.", part1, 2)
	// }
}
