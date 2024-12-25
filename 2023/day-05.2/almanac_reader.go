package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// buildAlmanacFromFile Reads the puzzle file and returns the data as an almanac instance
func buildAlmanacFromFile(filename string) almanac {
	content := readInputFileAsString(filename)
	return buildAlmanac(content)
}

func readInputFileAsString(filename string) string {
	file, _ := os.Open(filename)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	reader := bufio.NewScanner(file)
	builder := strings.Builder{}
	for reader.Scan() {
		line := reader.Text()
		builder.WriteString(line)
		builder.WriteString("\n")
	}
	return builder.String()
}

func buildAlmanac(content string) almanac {
	var seeds []int64
	results := make(map[string][]mapping)
	mappingsPattern := regexp.MustCompile("(?P<key>(?:-?\\w+)+\\b(?: map)?\\b):(?P<numbers>(?:\\s\\d+)+)")
	matches := mappingsPattern.FindAllStringSubmatch(content, -1)
	keyIndex := mappingsPattern.SubexpIndex("key")
	numbersIndex := mappingsPattern.SubexpIndex("numbers")
	for _, match := range matches {
		numMatches := len(match)
		if keyIndex < numMatches && numbersIndex < numMatches {
			key := match[keyIndex]
			dirtyNumbersString := match[numbersIndex]
			dirtyNumbersString = strings.Replace(dirtyNumbersString, "\n", " ", -1)
			numbers := splitNumbers(strings.TrimSpace(dirtyNumbersString))
			switch key {
			case "seeds":
				seeds = numbers
				continue

			default:
				mappings := make([]mapping, 0)
				length := len(numbers)
				for i := 0; i < length; i += 3 {
					destinationRangeStart := numbers[i]
					sourceRangeStart := numbers[i+1]
					rangeLength := numbers[i+2]
					mappings = append(mappings, mapping{
						SourceRangeStart:      sourceRangeStart,
						SourceRangeEnd:        sourceRangeStart + rangeLength,
						DestinationRangeStart: destinationRangeStart,
						RangeLength:           rangeLength,
					})
				}
				results[key] = mappings
			}

		}
	}
	return almanac{
		seeds: seeds,
		maps:  results,
	}
}

func splitNumbers(s string) []int64 {
	numberStrings := strings.Split(s, " ")
	numbers := make([]int64, 0)
	for _, numberString := range numberStrings {
		n, err := strconv.ParseInt(numberString, 10, 64)
		if err == nil {
			numbers = append(numbers, n)
		}
	}
	return numbers
}
