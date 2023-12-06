package main

import (
	"adventofcode-2023/utils"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fileLines, err := utils.ReadFile("day-04/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	//utils.PrintLines(fileLines)

	var cards []Card
	for _, fileLine := range fileLines {
		cards = append(cards, parseCard(fileLine))
	}

	//total := 0
	for i, card := range cards {
		cards[i] = checkCard(card)
		//total += card.Total

		fmt.Printf("%+v\n", cards[i])
	}
	//fmt.Printf("Total: %d\n", total)
	fmt.Printf("WinMore: %d\n", winMore(cards, len(cards)))
}

func winMore(cards []Card, n int) (total int) {

	for i := 0; i < n; i++ {
		total++

		if cards[i].Total > 0 {
			total += winMore(cards[i+1:], cards[i].Total)
		}
	}

	return total
}

func checkCard(card Card) Card {
	for _, winningNumber := range card.WinningNumbers {
		if slices.Contains(card.Numbers, winningNumber) {
			//if card.Total == 0 {
			//	card.Total = 1
			//} else {
			//	card.Total = card.Total * 2
			//}
			card.Total++
		}
	}
	return card
}

type Card struct {
	WinningNumbers []int
	Numbers        []int
	Total          int
}

func parseCard(line string) (card Card) {
	//split by :
	allCardNumbersString := strings.Split(line, ":")[1]

	//split AllNumbersString
	allWinningNumbersString := strings.Split(allCardNumbersString, "|")[0]
	allNumbersString := strings.Split(allCardNumbersString, "|")[1]

	eachWinningNumberString := strings.Split(allWinningNumbersString, " ")
	for _, wNumberString := range eachWinningNumberString {
		if len(strings.TrimSpace(wNumberString)) == 0 {
			continue
		}
		wNumber, err := strconv.Atoi(wNumberString)
		if err != nil {
			log.Fatal(err)
		}
		card.WinningNumbers = append(card.WinningNumbers, wNumber)
	}

	eachNumberString := strings.Split(allNumbersString, " ")
	for _, nString := range eachNumberString {
		if len(strings.TrimSpace(nString)) == 0 {
			continue
		}
		n, err := strconv.Atoi(nString)
		if err != nil {
			log.Fatal(err)
		}
		card.Numbers = append(card.Numbers, n)
	}

	return card
}
