package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

type gameRecord struct {
	Time     int
	Distance int // i.e. the record to beat
}

func readInputFileAsString(filename string) string {
	file, _ := os.Open(filename)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	bytes, err := io.ReadAll(file)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func splitNumbers(s string) []int {
	numberStrings := strings.Split(s, ",")
	numbers := make([]int, 0)
	for _, numberString := range numberStrings {
		n, err := strconv.Atoi(numberString)
		if err == nil {
			numbers = append(numbers, n)
		}
	}
	return numbers
}

func main() {

	cwd, _ := os.Getwd()
	inputPath := path.Join(cwd, "input.txt")
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		panic("the input file does not exist")
	}

	fetchRecords := func(content string) map[string][]int {
		pattern := regexp.MustCompile(`(?P<key>\w+):\s*(?P<numbers>\d+(?:\s+\d*)*)`)
		anyWhitespacePattern := regexp.MustCompile(`\s+`)
		matches := pattern.FindAllStringSubmatch(content, -1)
		sets := make(map[string][]int)
		for _, match := range matches {
			key := match[1]
			numbersString := anyWhitespacePattern.ReplaceAllString(match[2], ",")
			numbers := splitNumbers(numbersString)
			sets[key] = numbers
		}
		return sets
	}

	readRacingGames := func(filename string) []gameRecord {
		games := make([]gameRecord, 0)
		content := readInputFileAsString(filename)
		records := fetchRecords(content)
		times := records["Time"]
		distances := records["Distance"]
		if len(times) == len(distances) {
			for i := 0; i < len(times); i++ {
				g := gameRecord{
					Time:     times[i],
					Distance: distances[i],
				}
				games = append(games, g)
			}
		}
		return games
	}

	countWays := func(g gameRecord) int {
		result := 0
		for chargeMs := 1; chargeMs < g.Time; chargeMs++ {
			remainingRacingTime := g.Time - chargeMs
			speed := chargeMs * 1
			distance := remainingRacingTime * speed
			if distance > g.Distance {
				result++
			}
		}
		return result
	}

	result := 1
	games := readRacingGames(inputPath)
	for _, game := range games {
		numberOfWaysToBeatRecord := countWays(game)
		result *= numberOfWaysToBeatRecord
	}

	fmt.Printf("%d\n", result)

	os.Exit(0)

}
