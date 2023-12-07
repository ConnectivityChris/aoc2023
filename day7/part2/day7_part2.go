package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

type Hand struct {
	Cards string
	Bid   int
}

// AQK56 = 5 (1, 1, 1, 1, 1)	- High Card
// AAK56 = 7 (2, 2, 1, 1, 1)	- One Pair
// AAKK6 = 9 (2, 2, 2, 2, 1)	- Two Pair
// AAA56 = 11 (3, 3, 3, 1, 1) - Three of a kind
// AAAKK = 13 (3, 3, 3, 2, 2) - Full House
// AAAA6 = 17 (4, 4, 4, 4, 1) - Four of a kind
// AAAAA = 25 (5, 5, 5, 5, 5) - Five of a kind
var handValueMap = map[int]string{
	5:  "0",
	7:  "1",
	9:  "2",
	11: "3",
	13: "4",
	17: "5",
	25: "6",
}

func main() {
	fmt.Println("Part 2 - Total Winnings:", calculatePart2("../input.txt"))
}

func calculatePart2(filename string) int {
	input, _ := readInput(filename)
	totalWinnings := 0
	gameOrder := []Hand{}

	for _, game := range input {
		hand := Hand{}
		fmt.Sscanf(game, "%s %d", &hand.Cards, &hand.Bid)
		gameOrder = append(gameOrder, hand)
	}

	slices.SortFunc(gameOrder, func(a, b Hand) int {
		return compareHands(a.Cards, b.Cards)
	})

	for i, game := range gameOrder {
		totalWinnings += game.Bid * (i + 1)
	}
	return totalWinnings
}

func compareHands(a, b string) int {
	cardList := "23456789TQKA"
	// Define a replacer string to turns the card letters into ordered representations
	// J = 0 - Lowest value
	// T = A
	// Q = B
	// K = C
	// A = D - Highest value
	replacer := "TAQBKCADJ0"

	convertCardToNumber := func(cards string) string {
		handValue := 0
		for _, card := range strings.Split(cardList, "") {
			newCards, newHandValue := strings.ReplaceAll(cards, "J", card), 0
			for _, s := range newCards {
				newHandValue += strings.Count(newCards, string(s))
			}
			handValue = slices.Max([]int{handValue, newHandValue})
		}
		return handValueMap[handValue]
	}

	return strings.Compare(
		convertCardToNumber(a)+strings.NewReplacer(strings.Split(replacer, "")...).Replace(a),
		convertCardToNumber(b)+strings.NewReplacer(strings.Split(replacer, "")...).Replace(b),
	)
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
