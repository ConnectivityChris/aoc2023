package main

import (
	"bufio"
	"bytes"
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
	file, _ := os.Open("test")

	defer file.Close()

	part1LowestLocation := findLocation(file)
	fmt.Println("Part 1 - Lowest Locaton: ", part1LowestLocation)
}

func findLocation(file *os.File) int {
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

	// Assume list is in order
	for _, seed := range seedList {
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
		locationList = append(locationList, locationValue)
	}

	return slices.Min(locationList)
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

func createMapItem(line []byte) RangeMap {
	fields := bytes.Fields(line)

	destination, _ := strconv.Atoi(string(fields[0]))
	source, _ := strconv.Atoi(string(fields[1]))
	mapRange, _ := strconv.Atoi(string(fields[2]))

	newMap := RangeMap{
		dest:   destination,
		source: source,
		rng:    mapRange,
	}

	return newMap
}

func scanForMap(mapName *[]RangeMap, scanner *bufio.Scanner) {
	var line []byte
	scanner.Scan()
	for scanner.Scan() {
		line = scanner.Bytes()
		if len(line) == 0 {
			break
		}

		*mapName = append(*mapName, createMapItem(line))
	}
}
