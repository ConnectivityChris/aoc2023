package main

import (
	"io"
	"os"
	"testing"
)

func TestDay2(t *testing.T) {
	file, _ := os.Open("input.txt")

	defer file.Close()

	content, _ := io.ReadAll(file)

	possibleTotal, powerSum := findPossibleGames(content)

	if possibleTotal != 2685 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", possibleTotal, 2685)
	}
	if powerSum != 83707 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", powerSum, 83707)
	}
}

func TestDay2Example(t *testing.T) {
	file, _ := os.Open("testinput.txt")

	defer file.Close()

	content, _ := io.ReadAll(file)

	possibleTotal, powerSum := findPossibleGames(content)

	if possibleTotal != 8 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", possibleTotal, 8)
	}
	if powerSum != 2286 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", powerSum, 2286)
	}
}

func BenchmarkDay2(b *testing.B) {
	file, _ := os.Open("input.txt")

	defer file.Close()

	content, _ := io.ReadAll(file)

	for i := 0; i < b.N; i++ {
		findPossibleGames(content)
	}
}
