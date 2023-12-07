package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	cards string
	bid   int
}

var (
	fiveOfAKind  = make([]Hand, 0)
	fourOfAKind  = make([]Hand, 0)
	fullHouse    = make([]Hand, 0)
	threeOfAKind = make([]Hand, 0)
	twoPair      = make([]Hand, 0)
	onePair      = make([]Hand, 0)
	highCard     = make([]Hand, 0)
)

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	content, _ := io.ReadAll(file)
	input := strings.Split(string(content), "\n")
	fmt.Println("Total Winnings:", calculatePart1(input))
}

func calculatePart1(input []string) int {
	totalWinnings := 0

	for _, game := range input {
		fields := strings.Fields(game)
		cards := fields[0]
		bid, _ := strconv.Atoi(fields[1])
		hand := Hand{cards, bid}
		cardCount := make(map[string]int)

		for _, card := range cards {
			cardStr := string(card)
			cardCount[cardStr]++
		}

		numberCount := make(map[int]int)
		for _, count := range cardCount {
			numberCount[count]++
		}

		switch {
		case numberCount[5] == 1:
			fiveOfAKind = append(fiveOfAKind, hand)
		case numberCount[4] == 1:
			fourOfAKind = append(fourOfAKind, hand)
		case numberCount[3] == 1 && numberCount[2] == 1:
			fullHouse = append(fullHouse, hand)
		case numberCount[3] == 1:
			threeOfAKind = append(threeOfAKind, hand)
		case numberCount[2] == 2:
			twoPair = append(twoPair, hand)
		case numberCount[2] == 1:
			onePair = append(onePair, hand)
		default:
			highCard = append(highCard, hand)
		}
	}

	// Sort the categories based on the cards left to right
	sortAllTheMaps()
	// print all of the card totals
	fmt.Println("fiveOfAKind:", fiveOfAKind)
	fmt.Println("fourOfAKind:", fourOfAKind)
	fmt.Println("fullHouse:", fullHouse)
	fmt.Println("threeOfAKind:", threeOfAKind)
	fmt.Println("twoPair:", twoPair)
	fmt.Println("onePair:", onePair)
	fmt.Println("highCard:", highCard)
	mergedHands := mergeTheMaps()

	for i, hand := range mergedHands {
		totalWinnings += hand.bid * (i + 1)
	}

	return totalWinnings
}

func mergeTheMaps() []Hand {
	margedMap := make([]Hand, 0)
	margedMap = append(margedMap, highCard...)
	margedMap = append(margedMap, onePair...)
	margedMap = append(margedMap, twoPair...)
	margedMap = append(margedMap, threeOfAKind...)
	margedMap = append(margedMap, fullHouse...)
	margedMap = append(margedMap, fourOfAKind...)
	margedMap = append(margedMap, fiveOfAKind...)

	fmt.Printf("collapsedMap: %v\n", margedMap)
	return margedMap
}

func sortAllTheMaps() {
	sort.Slice(fiveOfAKind[:], func(i, j int) bool {
		return customSort(fiveOfAKind[i], fiveOfAKind[j])
	})
	sort.Slice(fourOfAKind[:], func(i, j int) bool {
		return customSort(fourOfAKind[i], fourOfAKind[j])
	})
	sort.Slice(fullHouse[:], func(i, j int) bool {
		return customSort(fullHouse[i], fullHouse[j])
	})
	sort.Slice(threeOfAKind[:], func(i, j int) bool {
		return customSort(threeOfAKind[i], threeOfAKind[j])
	})
	sort.Slice(twoPair[:], func(i, j int) bool {
		return customSort(twoPair[i], twoPair[j])
	})
	sort.Slice(onePair[:], func(i, j int) bool {
		return customSort(onePair[i], onePair[j])
	})
	sort.Slice(highCard[:], func(i, j int) bool {
		return customSort(highCard[i], highCard[j])
	})
}

func cardRank(card byte) int {
	rankOrder := "AKQJT98765432"
	return strings.Index(rankOrder, string(card)) + 1
}

func customSort(a, b Hand) bool {
	for i := 0; i < len(a.cards); i++ {
		if cardRank(a.cards[i]) > cardRank(b.cards[i]) {
			return true
		} else if cardRank(a.cards[i]) < cardRank(b.cards[i]) {
			return false
		}
	}
	return false
}
