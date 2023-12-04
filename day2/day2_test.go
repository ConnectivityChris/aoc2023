package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestDay2(t *testing.T) {
	file, _ := os.Open("input.txt")

	defer file.Close()

	content, _ := io.ReadAll(file)

	stringArray := strings.Split(string(content), "\n")
	possibleTotal := 0
	powerSum := 0
	for _, str := range stringArray {
		// Part 1
		possible, gameId := checkGameIsPossible(str)
		if possible {
			possibleTotal += gameId
		}

		// Part 2
		powerSum += calculateGamePower(str)
	}
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

	stringArray := strings.Split(string(content), "\n")
	possibleTotal := 0
	powerSum := 0
	for _, str := range stringArray {
		// Part 1
		possible, gameId := checkGameIsPossible(str)
		if possible {
			possibleTotal += gameId
		}

		// Part 2
		powerSum += calculateGamePower(str)
	}
	if possibleTotal != 8 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", possibleTotal, 8)
	}
	if powerSum != 2286 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", powerSum, 2286)
	}
}
