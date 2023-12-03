package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	content, _ := io.ReadAll(file)
	stringArray := strings.Split(string(content), "\n")

	sum, gearRatio := calculateSum(stringArray)
	fmt.Println("Sum is: ", sum)
	fmt.Println("Gear Ratio is: ", gearRatio)
}

type Point struct {
	X int
	Y int
}

var coordinateChecks = []Point{
	{0, -1},  // Up
	{1, -1},  // Up Right
	{1, 0},   // Right
	{1, 1},   // Down Right
	{0, 1},   // Down
	{-1, 1},  // Down Left
	{-1, 0},  // Left
	{-1, -1}, // Up Left
}

var listOfAsteriks = map[Point]bool{}

var gearRatioMap = map[Point][]int{}

func calculateSum(input []string) (int, int) {
	sum := 0
	gearRatio := 0
	currentNumber := 0
	symbolFound := false

	for y, line := range input {
		for x, char := range line {
			if unicode.IsDigit(char) {
				value := int(char - '0')
				currentNumber = currentNumber*10 + value
				for _, point := range coordinateChecks {
					neighbourX := x + point.X
					neighbourY := y + point.Y
					if neighbourY < 0 || neighbourY >= len(input) || neighbourX < 0 || neighbourX >= len(input[0]) {
						continue
					}
					neighbouringChar := rune(input[neighbourY][neighbourX])
					if !unicode.IsDigit(neighbouringChar) && neighbouringChar != '.' {
						symbolFound = true
					}
					if neighbouringChar == '*' {
						listOfAsteriks[Point{neighbourX, neighbourY}] = true
					}
				}
			} else {
				if currentNumber != 0 && symbolFound {
					sum += currentNumber
				}
				if currentNumber != 0 && len(listOfAsteriks) > 0 {
					for asterisk := range listOfAsteriks {
						x := asterisk.X
						y := asterisk.Y
						gearRatioMap[Point{x, y}] = append(gearRatioMap[Point{x, y}], currentNumber)
					}
				}
				currentNumber = 0
				symbolFound = false
				listOfAsteriks = map[Point]bool{}
			}
		}
		if currentNumber != 0 && symbolFound {
			sum += currentNumber
		}
		if currentNumber != 0 && len(listOfAsteriks) > 0 {
			for asterisk := range listOfAsteriks {
				x := asterisk.X
				y := asterisk.Y
				gearRatioMap[Point{x, y}] = append(gearRatioMap[Point{x, y}], currentNumber)
			}
		}
		currentNumber = 0
		symbolFound = false
		listOfAsteriks = map[Point]bool{}
	}

	for _, value := range gearRatioMap {
		if len(value) == 2 {
			gearRatio += value[0] * value[1]
		}
	}

	return sum, gearRatio
}
