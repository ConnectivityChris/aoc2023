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
	input := strings.Split(string(content), "\n")
	solution := calculatePart1(input)

	if solution != 288 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", solution, 288)
	}
}

func TestPart2WithExampleInput(t *testing.T) {
	file, _ := os.Open("testinput.txt")

	defer file.Close()

	content, _ := io.ReadAll(file)
	input := strings.Split(string(content), "\n")
	solution := calculatePart2(input)

	if solution != 71503 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", solution, 71503)
	}
}
