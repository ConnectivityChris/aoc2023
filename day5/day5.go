package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// type RangeMap struct {
// 	sourceRange []int
// 	destRange   []int
// }

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	content, _ := io.ReadAll(file)

	stringArray := strings.Split(string(content), "\n")
	lowestLocation := findLocation(stringArray)
	fmt.Println("Lowest Locaton: ", lowestLocation)
}

func findLocation(input []string) int {
	lowestLoction := 0
	seedList := strings.Fields(strings.Split(input[0], "seeds: ")[1])
	// fmt.Println(seedList)
	re := regexp.MustCompile("map:")

	mappingIndicies := make([]int, 0)

	for i, str := range input {
		if re.MatchString(str) {
			mappingIndicies = append(mappingIndicies, i)
		}
	}

	overallMappings := make(map[string]map[int]int)
	for i, index := range mappingIndicies {
		startIndex := index + 1
		endIndex := len(input) - 1
		if i != len(mappingIndicies)-1 {
			endIndex = mappingIndicies[i+1] - 2
		}

		mappings := make(map[int]int)
		for j := startIndex; j <= endIndex; j++ {
			// fmt.Println(strings.Fields(input[j]))
			source, _ := strconv.Atoi(strings.Fields(input[j])[1])
			destination, _ := strconv.Atoi(strings.Fields(input[j])[0])
			mapRange, _ := strconv.Atoi(strings.Fields(input[j])[2])
			for k := 0; k < mapRange; k++ {
				mappings[source+k] = destination + k
			}
		}
		overallMappings[strings.Fields(input[index])[0]] = mappings

	}
	fmt.Println(overallMappings)
	for _, seed := range seedList {
		// lookup soil mapping
		seedId, _ := strconv.Atoi(seed)
		soilId := checkIDIsMapped(seedId, overallMappings["seed-to-soil"])
		// fmt.Printf("Seed ID %d needs Soil ID %d\n", seedId, soilId)
		fertilizerId := checkIDIsMapped(soilId, overallMappings["soil-to-fertilizer"])
		// fmt.Printf("Soil ID %d needs Fertilizer ID %d\n", soilId, fertilizerId)
		waterId := checkIDIsMapped(fertilizerId, overallMappings["fertilizer-to-water"])
		// fmt.Printf("Fertilizer ID %d needs Water ID %d\n", fertilizerId, waterId)
		lightId := checkIDIsMapped(waterId, overallMappings["water-to-light"])
		// fmt.Printf("Water ID %d needs Light ID %d\n", waterId, lightId)
		temperatureId := checkIDIsMapped(lightId, overallMappings["light-to-temperature"])
		// fmt.Printf("Light ID %d needs Temperature ID %d\n", lightId, temperatureId)
		humidityId := checkIDIsMapped(temperatureId, overallMappings["temperature-to-humidity"])
		// fmt.Printf("Temperature ID %d needs Humidity ID %d\n", temperatureId, humidityId)
		locationId := checkIDIsMapped(humidityId, overallMappings["humidity-to-location"])
		// fmt.Printf("Humidity ID %d needs Location ID %d\n", humidityId, locationId)

		if locationId < lowestLoction || lowestLoction == 0 {
			lowestLoction = locationId
		}
	}

	return lowestLoction
}

func checkIDIsMapped(id int, mapToCheck map[int]int) int {
	mappedId := mapToCheck[id]
	if mappedId != 0 {
		return mappedId
	}
	return id
}
