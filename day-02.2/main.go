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

	powerOfGame := func(gameSamples []map[string]int) int {
		sampleStats := make(map[string]int)
		for _, sample := range gameSamples {
			for color, count := range sample {
				value, found := sampleStats[color]
				if found == false || (found && value < count) {
					sampleStats[color] = count
				}
			}
		}
		var product = 1
		for _, value := range sampleStats {
			product *= value
		}
		return product
	}

	file, _ := os.Open(inputPath)
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		next := reader.Text()
		samples := readSamples(next)
		power := powerOfGame(samples)
		sum += power
	}

	fmt.Printf("%d\n", sum)

	os.Exit(0)
}
