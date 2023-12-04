package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestDay1(t *testing.T) {
	file, _ := os.Open("input.txt")

	defer file.Close()

	content, _ := io.ReadAll(file)

	stringArray := strings.Split(string(content), "\n")
	part1Result := calculateTotal(stringArray)
	part2Result := calculateTotalWithWords(stringArray)
	if part1Result != 56049 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", part1Result, 56049)
	}
	if part2Result != 54530 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", part2Result, 54530)
	}
}

func TestDay1Part1Example(t *testing.T) {
	file, _ := os.Open("testinput.txt")

	defer file.Close()

	content, _ := io.ReadAll(file)

	stringArray := strings.Split(string(content), "\n")
	part1Result := calculateTotal(stringArray)
	if part1Result != 142 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", part1Result, 142)
	}
}

func TestDay1Part2Example(t *testing.T) {
	file, _ := os.Open("testinput2.txt")

	defer file.Close()

	content, _ := io.ReadAll(file)

	stringArray := strings.Split(string(content), "\n")
	part2Result := calculateTotalWithWords(stringArray)

	if part2Result != 281 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", part2Result, 281)
	}
}
