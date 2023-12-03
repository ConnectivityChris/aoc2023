package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var gearRatioMap = map[string][]int{}

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	content, _ := io.ReadAll(file)

	sum, gearRatio := calculateSum(strings.Split(string(content), "\n"))
	fmt.Println("Sum is: ", sum)
	fmt.Println("Gear Ratio is: ", gearRatio)
}

func calculateSum(stringArray []string) (int, int) {
	sum := 0

	for l, line := range stringArray {
		startIndex := -1
		endIndex := -1
		for i, char := range line {
			if unicode.IsDigit(char) && startIndex == -1 {
				startIndex = i
				endIndex = i
			} else if unicode.IsDigit(char) && startIndex != -1 {
				endIndex = i
			}

			if !unicode.IsDigit(char) || i == len(line)-1 {
				if startIndex != -1 && endIndex != -1 {
					testString := line[startIndex : endIndex+1]
					symbolFound := false

					previousLine := ""
					nextLine := ""

					if l > 0 {
						previousLine = stringArray[l-1]
					}

					if l != len(stringArray)-1 {
						nextLine = stringArray[l+1]
					}

					if startIndex > 0 {
						startIndex--
					}
					if endIndex < len(nextLine)-1 {
						endIndex++
					}

					if previousLine != "" {
						previousLineSlice := previousLine[startIndex : endIndex+1]
						symbolFound = checkForSymbolInLine(previousLineSlice, testString, startIndex, l-1)

					}

					if nextLine != "" && !symbolFound {
						nextLineSlice := nextLine[startIndex : endIndex+1]
						symbolFound = checkForSymbolInLine(nextLineSlice, testString, startIndex, l+1)

					}

					if !symbolFound {
						if startIndex > 0 {
							leftChar := line[startIndex]
							symbolFound = checkForSymbolInLine(string(leftChar), testString, startIndex, l)

						}
						if !symbolFound {
							if endIndex < len(line) {
								rightChar := line[endIndex]
								symbolFound = checkForSymbolInLine(string(rightChar), testString, endIndex, l)

							}
						}
					}

					if symbolFound {
						partNumber, _ := strconv.Atoi(testString)
						sum += partNumber
					}

					startIndex = -1
					endIndex = -1
				}
			}
		}
	}

	return sum, getGearRatio()
}

func checkForSymbolInLine(stringToCheck string, partNumber string, startIndex int, lineNumber int) bool {
	if strings.Contains(stringToCheck, "*") {
		partId, _ := strconv.Atoi(partNumber)
		gearRatioMap[fmt.Sprintf("Line%d;Index%d", lineNumber, strings.Index(stringToCheck, "*")+startIndex)] = append(gearRatioMap[fmt.Sprintf("Line%d;Index%d", lineNumber, strings.Index(stringToCheck, "*")+startIndex)], partId)
		return true
	}
	return strings.ContainsAny(stringToCheck, "*#-+@%&=$/")
}

func getGearRatio() int {
	gearRatio := 0
	for _, value := range gearRatioMap {
		if len(value) > 1 {
			gearRatio += value[0] * value[1]
		}
	}
	return gearRatio
}
