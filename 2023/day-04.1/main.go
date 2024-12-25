package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

type cardRecord struct {
	Number         int
	WinningNumbers map[int]struct{}
	Numbers        []int
}

func newCardRecord(number int, winningNumbers []int, numbers []int) *cardRecord {
	winningNumbersSet := make(map[int]struct{})
	for _, n := range winningNumbers {
		winningNumbersSet[n] = struct{}{}
	}
	return &cardRecord{
		Number:         number,
		WinningNumbers: winningNumbersSet,
		Numbers:        numbers,
	}
}

// CalculatePoints +1 for the first match, then doubled for each match after the first
func (r *cardRecord) CalculatePoints() int {
	var points = 0
	for _, number := range r.Numbers {
		_, isWinner := r.WinningNumbers[number]
		if isWinner {
			points++
		}
	}
	if points > 0 {
		// works because, 2^0 == 1
		return int(math.Pow(2, float64(points-1)))
	}
	return 0
}

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

func main() {

	cwd, _ := os.Getwd()
	inputPath := path.Join(cwd, "input.txt")
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		panic("the input file does not exist")
	}

	cardsReader := newCardsParser()

	var total = 0

	file, _ := os.Open(inputPath)
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		next := reader.Text()
		card, err := cardsReader.ReadCardFromString(next)
		if err != nil {
			continue
		}
		points := card.CalculatePoints()
		total += points
	}

	fmt.Printf("%d\n", total)

	os.Exit(0)
}
