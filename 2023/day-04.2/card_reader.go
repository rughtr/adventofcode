package main

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type cardsParser struct {
	r *regexp.Regexp
}

func newCardsParser() *cardsParser {
	r := regexp.MustCompile("Card\\s+(?P<card>\\d+):\\s+(?P<winners>[\\d\\s]+)\\s+\\|\\s+(?P<numbers>[\\d\\s]+)")
	return &cardsParser{
		r,
	}
}

func (p *cardsParser) ReadCardFromString(s string) (*cardRecord, error) {
	matches := p.r.FindStringSubmatch(s)
	cardIndex := p.r.SubexpIndex("card")
	winningNumbersIndex := p.r.SubexpIndex("winners")
	numbersIndex := p.r.SubexpIndex("numbers")
	if cardIndex >= 0 && winningNumbersIndex >= 0 && numbersIndex >= 0 {
		card, _ := strconv.Atoi(matches[cardIndex])
		winningNumbers := p.parseNumbers(matches[winningNumbersIndex])
		numbers := p.parseNumbers(matches[numbersIndex])
		return newCardRecord(card, winningNumbers, numbers), nil
	}
	return nil, errors.New("cannot parse card record")
}

func (p *cardsParser) parseNumbers(s string) []int {
	result := make([]int, 0)
	for _, numberString := range strings.Split(s, " ") {
		number, err := strconv.Atoi(numberString)
		if err != nil {
			continue
		}
		result = append(result, number)
	}
	return result
}

func readCardsFromFile(cardsFile string) map[int]cardRecord {
	cards := make(map[int]cardRecord)
	cardsReader := newCardsParser()
	file, _ := os.Open(cardsFile)
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		next := reader.Text()
		card, err := cardsReader.ReadCardFromString(next)
		if err != nil {
			continue
		}
		cards[card.Number] = *card
	}
	return cards
}
