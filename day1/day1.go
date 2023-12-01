package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	stringArray := strings.Split(string(content), "\n")
	fmt.Println("Part 1 total: ", calculateTotal(stringArray))
	// part2Array := calculateTotalWithWords(stringArray)
	fmt.Println("Part 2 total: ", calculateTotalWithWords(stringArray))
}

func calculateTotal(stringArray []string) int {
	re := regexp.MustCompile(`\d`)
	total := 0
	for _, str := range stringArray {
		submatchall := re.FindAllString(str, -1)
		firstNumber := submatchall[0]
		// fmt.Println("First number: ", firstNumber)
		lastNumber := submatchall[len(submatchall)-1]
		// fmt.Println("Last number: ", lastNumber)
		numAsString := firstNumber + lastNumber
		// fmt.Println("Num as string: ", numAsString)
		num, _ := strconv.Atoi(numAsString)
		// fmt.Println("Num: ", num)
		total += num
	}
	return total
}

var numberMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"ten":   10,
}

func calculateTotalWithWords(stringArray []string) int {
	total := 0
	for _, str := range stringArray {
		var charSlice []int
		charMap := make(map[int]int)
		for stringDigit, intDigit := range numberMap {
			start := strings.Index(str, stringDigit)
			last := strings.LastIndex(str, stringDigit)
			if start != -1 {
				charSlice = append(charSlice, start)
				// fmt.Println("Found ", stringDigit, " at ", start, " in ", str)
				charMap[start] = intDigit
			}
			if last != -1 {
				charSlice = append(charSlice, last)
				// fmt.Println("Found ", stringDigit, " at ", last, " in ", str)
				charMap[last] = intDigit
			}
		}

		for index, char := range str {
			if unicode.IsDigit(char) {
				charSlice = append(charSlice, index)
				charMap[index] = int(char) - 48
			}
		}

		// fmt.Println("Char map: ", charMap)
		sort.Ints(charSlice)
		// fmt.Println("Char slice: ", charSlice)
		if len(charSlice) > 0 {
			// fmt.Println("The first digit is: ", charMap[charSlice[0]])
			// fmt.Println("The last digit is: ", charMap[charSlice[len(charSlice)-1]])
			first := strconv.Itoa(charMap[charSlice[0]])
			last := strconv.Itoa(charMap[charSlice[len(charSlice)-1]])
			combinedString := first + last
			// fmt.Println("Combined string: ", combinedString)
			num, _ := strconv.Atoi(combinedString)
			total += num
		}
	}
	return total
}
