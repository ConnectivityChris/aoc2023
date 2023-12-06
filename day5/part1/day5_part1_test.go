package main

import (
	"os"
	"testing"
)

func TestPart1WithExampleInput(t *testing.T) {
	file, _ := os.Open("../testinput.txt")

	defer file.Close()

	location := findLocation(file)

	if location != 35 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", location, 35)
	}
}

func TestPart1WithFullInput(t *testing.T) {
	file, _ := os.Open("../input.txt")

	defer file.Close()

	location := findLocation(file)

	if location != 218513636 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", location, 218513636)
	}
}

func BenchmarkDay5Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		file, _ := os.Open("../input.txt")

		defer file.Close()
		findLocation(file)
	}
}
