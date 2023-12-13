package main

import "testing"

func TestWithExampleInput(t *testing.T) {
	part1, part2 := calculateSolution("testinput.txt")

	if part1 != 21 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", part1, 21)
	}

	if part2 != 525152 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", part2, 525152)
	}
}

func TestWithFullInput(t *testing.T) {
	part1, part2 := calculateSolution("input.txt")

	if part1 != 7490 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", part1, 7490)
	}

	if part2 != 65607131946466 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", part2, 65607131946466)
	}
}
