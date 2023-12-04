package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestWithExampleInput(t *testing.T) {
	file, _ := os.Open("testinput.txt")

	defer file.Close()

	content, _ := io.ReadAll(file)

	points := calculatePoints(strings.Split(string(content), "\n"))
	totalScratchCards := findAllCards(strings.Split(string(content), "\n"))

	if points != 13 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", points, 13)
	}
	if totalScratchCards != 30 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", totalScratchCards, 30)
	}
}

func TestWithFullInput(t *testing.T) {
	file, _ := os.Open("input.txt")

	defer file.Close()

	content, _ := io.ReadAll(file)

	points := calculatePoints(strings.Split(string(content), "\n"))
	totalScratchCards := findAllCards(strings.Split(string(content), "\n"))

	if points != 18619 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", points, 18619)
	}
	if totalScratchCards != 8063216 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", totalScratchCards, 8063216)
	}
}
