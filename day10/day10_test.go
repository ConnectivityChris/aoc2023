package main

import "testing"

func TestWithExampleInput(t *testing.T) {
	// part1, _ := calculateSolution("testinput.txt")
	part1_2, _ := calculateSolution("testinput2.txt")

	// if part1 != 4 {
	// 	t.Errorf("Result was incorrect, got: %d, want: %d.", part1, 4)
	// }

	// if part2 != 2 {
	// 	t.Errorf("Result was incorrect, got: %d, want: %d.", part1, 2)
	// }

	if part1_2 != 8 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", part1_2, 8)
	}

	// if part2_2 != 2 {
	// 	t.Errorf("Result was incorrect, got: %d, want: %d.", part2_2, 2)
	// }
}
