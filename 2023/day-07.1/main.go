package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type handAndBid struct {
	Hand   string
	Amount int
}

type CardKind int

const (
	None        CardKind = 0
	FiveOfAKind CardKind = iota
	FourOfAKind
	FullHouse
	ThreeOfAKind
	TwoPairs
	OnePair
)

func cardValue(label rune) uint8 {
	symbolMap := map[rune]uint8{
		'2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
		'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14,
	}
	value, _ := symbolMap[label]
	return value
}

func (v handAndBid) stats() map[uint8]int {
	s := make(map[uint8]int)
	hand := v.Hand
	for i := 0; i < len(hand); i++ {
		value, _ := s[hand[i]]
		s[hand[i]] = value + 1
	}
	return s
}

func (v handAndBid) hasNOfKind(n int) bool {
	s := v.stats()
	for _, count := range s {
		if count == n {
			return true
		}
	}
	return false
}

func (v handAndBid) isFiveOfAKind() bool {
	return v.hasNOfKind(5)
}

func (v handAndBid) isFourOfAKind() bool {
	return v.hasNOfKind(4)
}

func (v handAndBid) isFullHouse() bool {
	return v.hasNOfKind(3) && v.hasOnePair()
}

func (v handAndBid) isThreeOfAKind() bool {
	return v.hasNOfKind(3) && v.hasOnePair() == false
}

func (v handAndBid) hasTowPairs() bool {
	s := v.stats()
	result := 0
	for _, count := range s {
		if count == 2 {
			result++
		}
	}
	return result == 2
}

func (v handAndBid) hasOnePair() bool {
	return v.hasNOfKind(2)
}

func (v handAndBid) kind() CardKind {
	if v.isFiveOfAKind() {
		return FiveOfAKind
	} else if v.isFourOfAKind() {
		return FourOfAKind
	} else if v.isFullHouse() {
		return FullHouse
	} else if v.isThreeOfAKind() {
		return ThreeOfAKind
	} else if v.hasTowPairs() {
		return TwoPairs
	} else if v.hasOnePair() {
		return OnePair
	}
	return None
}

func (v handAndBid) bid() int {
	return v.Amount
}

func (v handAndBid) value() int {
	var result int
	runes := []rune(v.Hand)
	length := len(runes)
	for i := 0; i < length; i++ {
		value := cardValue(runes[i])
		result = (result << 4) | int(value)
	}
	return result
}

func (v handAndBid) compare(other handAndBid) int {
	compareValue := func(a, b int) int {
		return b - a
	}
	value := v.value()
	otherValue := other.value()
	valueKind := v.kind()
	otherValueKind := other.kind()
	if valueKind != None && valueKind == otherValueKind {
		return compareValue(value, otherValue)
	} else if valueKind != None && otherValueKind != None && valueKind != otherValueKind {
		if valueKind < otherValueKind {
			return -1
		} else {
			return 1
		}
	} else if valueKind == None && otherValueKind != None {
		return 1
	} else if valueKind != None && otherValueKind == None {
		return -1
	}
	return compareValue(value, otherValue)
}

type HandsAndBidsList []handAndBid

func (h HandsAndBidsList) Len() int {
	return len(h)
}

func (h HandsAndBidsList) Less(i, j int) bool {
	a := h[i]
	b := h[j]
	return a.compare(b) < 0
}

func (h HandsAndBidsList) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func readInput(filename string) HandsAndBidsList {

	file, _ := os.Open(filename)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	results := make(HandsAndBidsList, 0)

	r := regexp.MustCompile("(?P<hand>[AKQJT2-9]+)\\s(?P<bid>\\d+)")

	reader := bufio.NewScanner(file)
	for reader.Scan() {
		line := reader.Text()
		matches := r.FindAllStringSubmatch(line, -1)
		handIndex := r.SubexpIndex("hand")
		bidIndex := r.SubexpIndex("bid")
		for _, match := range matches {
			if handIndex >= 0 && bidIndex >= 0 && handIndex < len(match) && bidIndex < len(match) {
				hand := match[handIndex]
				bid, err := strconv.Atoi(strings.TrimSpace(match[bidIndex]))
				if err != nil {
					panic("invalid bid")
				}
				value := handAndBid{Hand: hand, Amount: bid}
				results = append(results, value)
			}
		}
	}

	return results
}

func main() {

	cwd, _ := os.Getwd()
	inputPath := path.Join(cwd, "input.txt")
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		panic("the input file does not exist")
	}

	list := readInput(inputPath)
	sort.Sort(list)
	maxRank := len(list)
	var total = int64(0)
	for i, next := range list {
		winning := int64((maxRank - i) * next.bid())
		total += winning
		kindString := formatKind(next)
		fmt.Printf("%s: %s, %d => (%d): %d\n",
			next.Hand, kindString, next.value(), maxRank-i, winning)
	}

	fmt.Printf("%d\n", total)

	os.Exit(0)
}

func formatKind(next handAndBid) string {
	kindString := ""
	switch next.kind() {
	case FiveOfAKind:
		kindString = "Five of a kind"
	case FourOfAKind:
		kindString = "Four of a kind"
	case FullHouse:
		kindString = "Full house"
	case ThreeOfAKind:
		kindString = "Three of a kind"
	case TwoPairs:
		kindString = "Two pairs"
	case OnePair:
		kindString = "One pair"
	default:
		kindString = "High card"
	}
	return kindString
}
