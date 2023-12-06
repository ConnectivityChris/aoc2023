package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var cubeLimit map[string]int = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	content, _ := io.ReadAll(file)

	possibleTotal, powerSum := findPossibleGames(content)

	fmt.Println("Total of Possible Games:", possibleTotal)
	fmt.Println("Total Power:", powerSum)
}

func findPossibleGames(input []byte) (int, int) {
	stringArray := strings.Split(string(input), "\n")
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
	return possibleTotal, powerSum
}

func checkGameIsPossible(game string) (bool, int) {
	possible := true
	initialSplit := strings.Split(game, ": ")
	gameId, _ := strconv.Atoi(strings.Split(initialSplit[0], "Game ")[1])
	gameInfo := initialSplit[1]
	randSelectedCubes := strings.Split(gameInfo, ";")
	for _, str := range randSelectedCubes {
		if !possible {
			break
		}
		listOfCubes := strings.Split(str, ", ")
		for _, cube := range listOfCubes {
			number, color := extractNumberAndColor(cube)
			gameCheck := number <= cubeLimit[color]
			if !gameCheck {
				possible = false
				break
			}
		}
	}
	return possible, gameId
}

func calculateGamePower(game string) int {
	var cubeMinimums map[string]int = map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	pulls := strings.Split(strings.Split(game, ": ")[1], "; ")
	for _, pull := range pulls {
		listOfCubes := strings.Split(pull, ", ")
		for _, cube := range listOfCubes {
			number, color := extractNumberAndColor(cube)
			if number > cubeMinimums[color] {
				cubeMinimums[color] = number
			}
		}
	}
	power := cubeMinimums["red"] * cubeMinimums["green"] * cubeMinimums["blue"]
	return power
}

func extractNumberAndColor(pull string) (int, string) {
	cubeSplit := strings.Split(strings.Trim(pull, " "), " ")
	number, _ := strconv.Atoi(cubeSplit[0])
	color := cubeSplit[1]
	return number, color
}
