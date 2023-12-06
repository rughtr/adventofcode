package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"regexp"
	"strconv"
)

type gameRecord struct {
	Time     int64
	Distance int64 // i.e. the record to beat
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

func main() {

	cwd, _ := os.Getwd()
	inputPath := path.Join(cwd, "input.txt")
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		panic("the input file does not exist")
	}

	fetchRecords := func(content string) map[string]int64 {
		pattern := regexp.MustCompile(`(?P<key>\w+):\s*(?P<numbers>\d+(?:\s+\d*)*)`)
		anyWhitespacePattern := regexp.MustCompile(`\s+`)
		matches := pattern.FindAllStringSubmatch(content, -1)
		sets := make(map[string]int64)
		for _, match := range matches {
			key := match[1]
			numberString := anyWhitespacePattern.ReplaceAllString(match[2], "")
			number, err := strconv.ParseInt(numberString, 10, 64)
			if err == nil {
				sets[key] = number
			}
		}
		return sets
	}

	readRacingGame := func(filename string) gameRecord {
		content := readInputFileAsString(filename)
		records := fetchRecords(content)
		time := records["Time"]
		distance := records["Distance"]
		return gameRecord{
			Time:     time,
			Distance: distance,
		}
	}

	countWays := func(g gameRecord) int64 {
		result := int64(0)
		for chargeMs := int64(1); chargeMs < g.Time; chargeMs++ {
			remainingRacingTime := g.Time - chargeMs
			speed := chargeMs * 1
			distance := remainingRacingTime * speed
			if distance > g.Distance {
				result++
			}
		}
		return result
	}

	game := readRacingGame(inputPath)
	numberOfWaysToBeatRecord := countWays(game)

	fmt.Printf("%d\n", numberOfWaysToBeatRecord)

	os.Exit(0)

}
