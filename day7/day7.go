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
	handCategories = map[string][]Hand{
		"fiveOfAKind":  {},
		"fourOfAKind":  {},
		"fullHouse":    {},
		"threeOfAKind": {},
		"twoPair":      {},
		"onePair":      {},
		"highCard":     {},
	}
)

func main() {
	fmt.Println("Part 1 -Total Winnings:", calculatePart1("input.txt"))
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

func calculatePart1(filename string) int {
	input, _ := readInput(filename)
	totalWinnings := 0

	for _, game := range input {
		cards, bid := parseGame(game)
		hand := Hand{cards, bid}
		numberCount := countCards(cards)

		switch {
		case numberCount[5] == 1:
			handCategories["fiveOfAKind"] = append(handCategories["fiveOfAKind"], hand)
		case numberCount[4] == 1:
			handCategories["fourOfAKind"] = append(handCategories["fourOfAKind"], hand)
		case numberCount[3] == 1 && numberCount[2] == 1:
			handCategories["fullHouse"] = append(handCategories["fullHouse"], hand)
		case numberCount[3] == 1:
			handCategories["threeOfAKind"] = append(handCategories["threeOfAKind"], hand)
		case numberCount[2] == 2:
			handCategories["twoPair"] = append(handCategories["twoPair"], hand)
		case numberCount[2] == 1:
			handCategories["onePair"] = append(handCategories["onePair"], hand)
		default:
			handCategories["highCard"] = append(handCategories["highCard"], hand)
		}
	}

	sortAllTheCategories()
	mergedHands := mergeCategories()

	for i, hand := range mergedHands {
		totalWinnings += hand.bid * (i + 1)
	}

	return totalWinnings
}

func parseGame(game string) (string, int) {
	fields := strings.Fields(game)
	cards := fields[0]
	bid, _ := strconv.Atoi(fields[1])
	return cards, bid
}

func countCards(cards string) map[int]int {
	cardCount := make(map[string]int)
	numberCount := make(map[int]int)
	for _, card := range cards {
		cardCount[string(card)]++
	}

	for _, count := range cardCount {
		numberCount[count]++
	}
	return numberCount
}

func mergeCategories() []Hand {
	var mergedHands []Hand
	order := []string{"highCard", "onePair", "twoPair", "threeOfAKind", "fullHouse", "fourOfAKind", "fiveOfAKind"}

	for _, category := range order {
		mergedHands = append(mergedHands, handCategories[category]...)
	}
	return mergedHands
}

func sortAllTheCategories() {
	for _, category := range handCategories {
		sort.Slice(category, func(i, j int) bool {
			return customSort(category[i], category[j])
		})
	}
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
