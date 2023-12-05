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
	file, _ := os.Open("input.txt")

	defer file.Close()

	// content, _ := io.ReadAll(file)

	// stringArray := strings.Split(string(content), "\n")
	part1LowestLocation := findLocation(file, 1)
	fmt.Println("Part 1 - Lowest Locaton: ", part1LowestLocation)

	file2, _ := os.Open("input.txt")

	defer file2.Close()
	part2LowestLocation := findLocation(file2, 2)
	fmt.Println("Part 2 - Lowest Locaton: ", part2LowestLocation)
}

func findLocation(file *os.File, partNumber int) int {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	seedString := scanner.Text() // First line is always the seed list
	// seedList := strings.Split(seedString, "seeds: ")[1]
	seedStringList := strings.Fields(strings.Split(seedString, "seeds: ")[1])
	seedList := []int{}
	for _, seed := range seedStringList {
		seedInt, _ := strconv.Atoi(seed)
		seedList = append(seedList, seedInt)
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
		if partNumber == 2 {
			if i%2 == 0 {
				seedRange := seedList[i+1]
				endSeeds = append(endSeeds, (seed+seedRange)-1)
			}
		} else {
			// seed to soil
			soilValue := findMappedId(seedToSoil, seed)
			// soil to fertilizer
			fertilizerValue := findMappedId(soilToFertilizer, soilValue)
			// fertilizer to water
			waterValue := findMappedId(fertilizerToWater, fertilizerValue)
			// water to light
			lightValue := findMappedId(waterToLight, waterValue)
			// light to temp
			tempValue := findMappedId(lightToTemp, lightValue)
			// temp to humidity
			humidityValue := findMappedId(tempToHumidity, tempValue)
			// humidity to location
			locationValue := findMappedId(humidityToLocation, humidityValue)
			// fmt.Printf("seed: %v, %v %v %v %v %v %v %v\n", seed, soilValue, fertilizerValue, waterValue, lightValue, tempValue, humidityValue, locationValue)
			locationList = append(locationList, locationValue)
		}
	}

	if partNumber == 2 {
		endSeed := slices.Max(endSeeds)
		// fmt.Println(endSeed)
		// reachedSeedId := false
		for i := 0; i < endSeed; i++ {
			locationValue := i
			humidityValue := findInverseMappedId(humidityToLocation, locationValue)
			tempValue := findInverseMappedId(tempToHumidity, humidityValue)
			lightValue := findInverseMappedId(lightToTemp, tempValue)
			waterValue := findInverseMappedId(waterToLight, lightValue)
			fertilizerValue := findInverseMappedId(fertilizerToWater, waterValue)
			soilValue := findInverseMappedId(soilToFertilizer, fertilizerValue)
			seed := findInverseMappedId(seedToSoil, soilValue)
			// fmt.Println("Location Value:", locationValue)
			if checkSeedIsInRange(seed, seedList) {
				// fmt.Printf("seed: %v, %v %v %v %v %v %v %v\n", seed, soilValue, fertilizerValue, waterValue, lightValue, tempValue, humidityValue, locationValue)
				locationList = append(locationList, locationValue-1)
				break
			}
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

func findMappedId(mapToSearch []RangeMap, initialId int) int {
	mappedId := initialId
	for _, m := range mapToSearch {
		if initialId > m.source && initialId < m.source+m.rng {
			mappedId = m.dest + (initialId - m.source)
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
