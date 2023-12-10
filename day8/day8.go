package main

import (
	"fmt"
	"io"
	"math/big"
	"os"
	"regexp"
	"strings"
)

func main() {
	fmt.Printf("Part 1 - Steps to reach ZZZ: %d\n", calculateSteps("input.txt"))
	fmt.Printf("Part 2 - Steps to reach ZZZ: %d\n", calculateStepsPart2("input.txt"))
}

func calculateSteps(filename string) int {
	input, _ := readInput(filename)
	instructions := input[0]
	totalSteps := 0
	// fmt.Println(instructions)
	maps, _ := buildMaps(input[2:])
	// fmt.Println(maps)
	currentLocation := "AAA"
	for {
		if currentLocation == "ZZZ" {
			break
		}
		for _, instruction := range instructions {
			// fmt.Println("Current location:", currentLocation)
			// fmt.Println("Performing instruction", i, ":", string(instruction))
			left := maps[currentLocation][0]
			right := maps[currentLocation][1]
			// fmt.Printf("Left: %s, Right: %s\n", left, right)
			if instruction == 'L' {
				currentLocation = left
			} else {
				currentLocation = right
			}
			totalSteps++
		}
	}

	return totalSteps
}

func calculateStepsPart2(filename string) int {
	input, _ := readInput(filename)
	instructions := input[0]
	totalSteps := 0
	// fmt.Println(instructions)
	maps, currentLocations := buildMaps(input[2:])
	// fmt.Println("Maps :", maps)
	// fmt.Println("Locations :", currentLocations)
	for {
		if checkAllCurrentLocations(currentLocations) {
			break
		}
		for _, instruction := range instructions {
			for i, locinstruction := range currentLocations {
				// perform the movement for all locations
				if currentLocations[i].Steps == 0 {
					currentLocations[i].CurrentLocation = moveLocation(maps, locinstruction, instruction)
				}
				if strings.HasSuffix(currentLocations[i].CurrentLocation, "Z") {
					// fmt.Printf("Found end in %d steps\n", totalSteps)
					if currentLocations[i].Steps == 0 {
						currentLocations[i].Steps = totalSteps + 1
					}
				}
				// fmt.Println("Current location:", currentLocations[i])
				// moveTo := moveLocation(maps, locinstruction, instruction)
				// if strings.HasSuffix(moveTo, "Z") {
				// 	currentLocations[i] = Path{
				// 		Start: locinstruction.Start,
				// 		End:   moveTo,
				// 		Steps: totalSteps,
				// 	}
				// }
			}
			totalSteps++
		}
	}
	// fmt.Printf("Current Locations: %v\n", currentLocations)
	// print all the locations to find the lowest common multiple
	var result = intToBigInt(currentLocations[0].Steps)
	for _, loc := range currentLocations[1:] {
		result = lcm(result, intToBigInt(loc.Steps))
	}
	// fmt.Printf("LCM is: %d\n", result)

	return int(result.Int64())
}

func gcd(a, b *big.Int) *big.Int {
	for b.Sign() != 0 {
		a, b = b, new(big.Int).Mod(a, b)
	}
	return a
}

func lcm(a, b *big.Int) *big.Int {
	gcdAB := gcd(a, b)
	// fmt.Printf("GCD of %d and %d is: %d\n", a, b, gcdAB)
	return new(big.Int).Div(new(big.Int).Mul(a, b), gcdAB)
}

func intToBigInt(n int) *big.Int {
	return new(big.Int).SetInt64(int64(n))
}

func moveLocation(maps map[string][]string, location Path, instruction rune) string {
	// fmt.Println("Current location:", location)
	// fmt.Println("Performing instruction:", string(instruction))
	left := maps[location.CurrentLocation][0]
	right := maps[location.CurrentLocation][1]
	// fmt.Printf("Left: %s, Right: %s\n", left, right)
	if instruction == 'L' {
		return left
	} else {
		return right
	}
}

func checkAllCurrentLocations(locations []Path) bool {
	allValid := true
	for _, path := range locations {
		if path.Steps == 0 {
			allValid = false
			break
		}
	}
	// fmt.Printf("Locations: %v, Valid: %v\n", locations, allValid)
	return allValid
}

type Path struct {
	Start           string
	CurrentLocation string
	End             string
	Steps           int
}

func buildMaps(input []string) (map[string][]string, []Path) {
	// fmt.Println(input)
	maps := map[string][]string{}
	startingLocations := []Path{}
	re := regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`)
	for _, str := range input[0 : len(input)-1] {
		// fmt.Println(str)
		directions := re.FindStringSubmatch(str)
		location := directions[1]
		if strings.HasSuffix(location, "A") {
			startingLocations = append(startingLocations, Path{location, location, "", 0})
		}
		left := directions[2]
		right := directions[3]
		maps[location] = []string{left, right}
	}
	return maps, startingLocations
}

func readInput(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(content), "\n"), nil
}
