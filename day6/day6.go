package main

import (
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	content, _ := io.ReadAll(file)
	input := strings.Split(string(content), "\n")
	part1Solution := calculatePart1(input)
	log.Printf("Part 1 - Margin of error: %d\n", part1Solution)
	part2Solution := calculatePart2(input)
	log.Printf("Part 2 - Margin of error: %d\n", part2Solution)
}

func calculatePart1(input []string) int {
	times := strings.Fields(strings.Split(input[0], ":")[1])
	distancesToBeat := strings.Fields(strings.Split(input[1], ":")[1])

	return getMarginOfError(times, distancesToBeat)
}

func calculatePart2(input []string) int {
	times := strings.ReplaceAll(strings.Split(input[0], ":")[1], " ", "")
	distanceToBeat := strings.ReplaceAll(strings.Split(input[1], ":")[1], " ", "")
	return getMarginOfError([]string{times}, []string{distanceToBeat})
}

func solveQuadratic(time int, distance int) int {
	d := math.Sqrt(float64(time*time - 4*distance))

	from := math.Floor((float64(time) - d) / 2)
	to := math.Ceil((float64(time) + d) / 2)

	return int(to) - int(from) - 1
}

func getMarginOfError(times []string, distancesToBeat []string) int {
	marginOfError := 1

	for i, time := range times {
		maximumTime, _ := strconv.Atoi(time)
		distanceToBeat, _ := strconv.Atoi(distancesToBeat[i])
		totalWon := solveQuadratic(maximumTime, distanceToBeat)
		marginOfError *= totalWon
	}

	return marginOfError
}
