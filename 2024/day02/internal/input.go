package internal

import (
	shared "aoc24-shared"
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

// LoadReportsFromInput reads reports from the "input.txt" file and returns a slice of parsed Report structures.
func LoadReportsFromInput() []Report {

	f := shared.OpenInput()
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	return readReports(f)
}

func readReports(input *os.File) []Report {

	reports := make([]Report, 0)

	reader := bufio.NewScanner(input)
	for reader.Scan() {
		line := reader.Text()
		segments := strings.Split(line, " ")
		mapIter := shared.Map(segments, func(s string) int {
			value, _ := strconv.Atoi(s)
			return value
		})
		levels := slices.Collect(mapIter)
		r := NewReport(levels)
		reports = append(reports, r)
	}

	return reports
}
