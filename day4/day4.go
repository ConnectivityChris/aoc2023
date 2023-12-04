package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	content, _ := io.ReadAll(file)
	stringArray := strings.Split(string(content), "\n")
	points := calculatePoints(stringArray)
	fmt.Println("Total Points: ", points)

	totalCards := findAllCards(stringArray)
	fmt.Println("Total Cards: ", totalCards)
}

func calculatePoints(input []string) int {
	points := 0

	for _, line := range input {
		matches := findMatches(line)
		if matches > 0 {
			points += 1 << (matches - 1)
		}
	}

	return points
}

func findAllCards(input []string) int {
	totalCards := 0
	count := make([]int, len(input))
	for i, line := range input {
		matches := findMatches(line)
		for j := i + 1; j < i+matches+1; j++ {
			count[j] += count[i] + 1
		}
		totalCards += count[i] + 1
	}
	return totalCards
}

func findMatches(line string) int {
	re := regexp.MustCompile(`\d+`)
	totalWinningNumbers := 0
	initialSplit := strings.Split(line, ": ")
	// cardNumber := strings.Split(initialSplit[0], " ")[1]
	gameSplit := strings.Split(initialSplit[1], " | ")
	revealedNumbers := re.FindAllString(gameSplit[0], -1)
	winningNumbers := re.FindAllString(gameSplit[1], -1)

	for _, revealedNumber := range revealedNumbers {
		if stringInSlice(revealedNumber, winningNumbers) {
			totalWinningNumbers++
		}
	}
	return totalWinningNumbers
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
