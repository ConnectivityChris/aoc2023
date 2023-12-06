package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type RangeMap struct {
	source int
	dest   int
	rng    int
}

var (
	seedToSoil         = make([]RangeMap, 0)
	soilToFertilizer   = make([]RangeMap, 0)
	fertilizerToWater  = make([]RangeMap, 0)
	waterToLight       = make([]RangeMap, 0)
	lightToTemp        = make([]RangeMap, 0)
	tempToHumidity     = make([]RangeMap, 0)
	humidityToLocation = make([]RangeMap, 0)
)

func main() {

	file, _ := os.Open("../ryaninput.txt")

	defer file.Close()
	part1LowestLocation := findLocation(file)
	fmt.Println("Part 2 - Lowest Locaton: ", part1LowestLocation)
}

func findLocation(file *os.File) int {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	seedString := scanner.Text() // First line is always the seed list
	// seedList := strings.Split(seedString, "seeds: ")[1]
	seedStringList := strings.Fields(strings.Split(seedString, "seeds: ")[1])
	seedList := make([]int, len(seedStringList))
	for i, seed := range seedStringList {
		seedList[i], _ = strconv.Atoi(seed)
	}

	scanner.Scan()

	// Assume that the maps are in the right order
	scanForMap(&seedToSoil, scanner)
	scanForMap(&soilToFertilizer, scanner)
	scanForMap(&fertilizerToWater, scanner)
	scanForMap(&waterToLight, scanner)
	scanForMap(&lightToTemp, scanner)
	scanForMap(&tempToHumidity, scanner)
	scanForMap(&humidityToLocation, scanner)

	locationList := make([]int, 0)

	endSeeds := make([]int, 0)

	// Assume list is in order
	for i, seed := range seedList {
		if i%2 == 0 {
			seedRange := seedList[i+1]
			endSeeds = append(endSeeds, (seed+seedRange)-1)
		}
	}

	endSeed := slices.Max(endSeeds)

	for i := 1; i < endSeed; i++ {
		locationValue := i
		humidityValue := findInverseMappedId(humidityToLocation, locationValue)
		tempValue := findInverseMappedId(tempToHumidity, humidityValue)
		lightValue := findInverseMappedId(lightToTemp, tempValue)
		waterValue := findInverseMappedId(waterToLight, lightValue)
		fertilizerValue := findInverseMappedId(fertilizerToWater, waterValue)
		soilValue := findInverseMappedId(soilToFertilizer, fertilizerValue)
		seed := findInverseMappedId(seedToSoil, soilValue)
		// fmt.Println("Location Value:", locationValue)
		// fmt.Printf("seed: %v, %v %v %v %v %v %v %v\n", seed, soilValue, fertilizerValue, waterValue, lightValue, tempValue, humidityValue, locationValue)
		if checkSeedIsInRange(seed, seedList) {
			locationList = append(locationList, locationValue-1)
			break
		}
	}
	return slices.Min(locationList)
}

func checkSeedIsInRange(seed int, seedList []int) bool {
	// fmt.Println("Seed:", seed)
	for i, originalSeeds := range seedList {
		if i%2 == 0 {
			seedRange := seedList[i+1]
			if seed >= originalSeeds && seed < originalSeeds+seedRange {
				return true
			}
		}
	}
	return false
}

func findInverseMappedId(mapToSearch []RangeMap, initialId int) int {
	mappedId := initialId
	for _, m := range mapToSearch {
		if initialId > m.dest && initialId < m.dest+m.rng {
			mappedId = m.source + (initialId - m.dest)
		}
	}
	return mappedId
}

func createMapItem(line string) RangeMap {
	destination, _ := strconv.Atoi(strings.Fields(line)[0])
	source, _ := strconv.Atoi(strings.Fields(line)[1])
	mapRange, _ := strconv.Atoi(strings.Fields(line)[2])
	newMap := new(RangeMap)
	newMap.dest = destination
	newMap.source = source
	newMap.rng = mapRange
	return *newMap
}

func scanForMap(mapName *[]RangeMap, scanner *bufio.Scanner) {
	var line string
	scanner.Scan()
	for scanner.Scan() {
		line = scanner.Text()
		if line == "" {
			break
		}

		*mapName = append(*mapName, createMapItem(line))
	}
}
