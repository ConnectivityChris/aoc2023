package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestSum(t *testing.T) {
	gearRatioMap = map[string][]int{}
	file, _ := os.Open("testinput.txt")

	defer file.Close()

	content, _ := io.ReadAll(file)

	sum, gearRatio := calculateSum(strings.Split(string(content), "\n"))
	if sum != 4361 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", sum, 4361)
	}
	if gearRatio != 467835 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", gearRatio, 467835)
	}
}

func TestWithFullInput(t *testing.T) {
	gearRatioMap = map[string][]int{}
	file, _ := os.Open("input.txt")

	defer file.Close()

	content, _ := io.ReadAll(file)

	sum, gearRatio := calculateSum(strings.Split(string(content), "\n"))
	if sum != 540212 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", sum, 540212)
	}
	if gearRatio != 87605697 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", gearRatio, 87605697)
	}
}
