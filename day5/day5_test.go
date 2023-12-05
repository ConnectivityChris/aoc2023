package main

import (
	"os"
	"testing"
)

func TestWithExampleInput(t *testing.T) {
	file, _ := os.Open("testinput.txt")

	defer file.Close()

	points := findLocation(file)

	if points != 35 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", points, 35)
	}
}

func TestWithFullInput(t *testing.T) {
	file, _ := os.Open("input.txt")

	defer file.Close()

	points := findLocation(file)

	if points != 218513636 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", points, 218513636)
	}
}
