package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	cwd, _ := os.Getwd()
	inputPath := path.Join(cwd, "input.txt")
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		panic("the input file does not exist")
	}

	var sum = 0

	matchGroup := func(r *regexp.Regexp, s string, groupName string) (bool, string) {
		matches := r.FindStringSubmatch(s)
		idIndex := r.SubexpIndex(groupName)
		numMatches := len(matches)
		if numMatches > 0 && idIndex >= 0 && idIndex < numMatches {
			return true, matches[idIndex]
		}
		return false, ""
	}

	idRegex := regexp.MustCompile("Game\\s(?P<id>\\d+)")

	readGameId := func(record string) int {
		gameAndSubsets := strings.Split(record, ":")
		if len(gameAndSubsets) >= 0 {
			found, idString := matchGroup(idRegex, strings.TrimSpace(gameAndSubsets[0]), "id")
			if found && len(idString) > 0 {
				n, err := strconv.Atoi(idString)
				if err == nil {
					return n
				}
			}
		}
		return 0
	}

	sampleRegex := regexp.MustCompile("(?P<count>\\d+)\\s+\\b(?P<color>blue|green|red)\\b")

	readSamples := func(record string) []map[string]int {
		result := make([]map[string]int, 0)
		gameAndSubsets := strings.Split(record, ":") // get the subsets strings
		subsets := strings.Split(strings.TrimSpace(gameAndSubsets[1]), ";")
		for _, subset := range subsets {
			set := make(map[string]int)
			sample := strings.Split(subset, ",") // separate the cubes
			for _, next := range sample {
				hasCount, countString := matchGroup(sampleRegex, next, "count")
				hasColor, color := matchGroup(sampleRegex, next, "color")
				if hasCount && hasColor {
					count, _ := strconv.Atoi(countString)
					n, found := set[color]
					if found {
						count += n
					}
					set[color] = count
				}
			}
			result = append(result, set)
		}
		return result
	}

	// 12 red cubes, 13 green cubes, and 14 blue cubes
	possibleGames := map[string]int{
		"red": 12, "green": 13, "blue": 14,
	}

	isPossibleGame := func(game map[string]int) bool {
		for color, count := range possibleGames {
			n, _ := game[color]
			if n > count {
				return false
			}
		}
		return true
	}

	file, _ := os.Open(inputPath)
	reader := bufio.NewScanner(file)
nextGame:
	for reader.Scan() {
		next := reader.Text()
		gameId := readGameId(next)
		samples := readSamples(next)
		for _, sample := range samples {
			if isPossibleGame(sample) == false {
				continue nextGame
			}
		}
		sum += gameId
	}

	fmt.Printf("%d\n", sum)

	os.Exit(0)
}
