package main

import (
	"fmt"
	"os"
	"path"
)

func main() {

	cwd, _ := os.Getwd()
	inputPath := path.Join(cwd, "input.txt")
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		panic("the input file does not exist")
	}

	instances := MakeStack[int]()

	cards := readCardsFromFile(inputPath)
	for _, card := range cards {
		instances.Push(card.Number)
	}

	traverseCardsAndCountInstances := func() int {

		statistics := newScratchCardStats()

		for instances.Any() {

			cardNumber := instances.Pop()
			statistics.CountScratchCard(cardNumber)

			card, _ := cards[cardNumber]
			numMatches := card.GetNumMatches()

			for i := 1; i <= numMatches; i++ {
				nextCardNumber := card.Number + i
				_, found := cards[nextCardNumber]
				if found {
					instances.Push(nextCardNumber)
				}
			}
		}

		return statistics.ScratchCardsTotal()
	}

	total := traverseCardsAndCountInstances()
	fmt.Printf("%d\n", total)

	os.Exit(0)
}
