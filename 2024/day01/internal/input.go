package internal

import (
	shared "aoc24-shared"
	"bufio"
	"os"
	"strconv"
	"strings"
)

const (
	sep = "   "
)

// LoadListsFromInput reads and parses input data from the input file into two integer slices and returns them.
func LoadListsFromInput() ([]int, []int) {

	f := shared.OpenInput()
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	return readLists(f)
}

func readLists(input *os.File) ([]int, []int) {

	left := make([]int, 0)
	right := make([]int, 0)

	reader := bufio.NewScanner(input)
	for reader.Scan() {
		line := reader.Text()
		segments := strings.Split(line, sep)
		if len(segments) != 2 {
			panic("invalid input")
		}
		left = append(left, parseInt(segments[0]))
		right = append(right, parseInt(segments[1]))
	}

	return left, right
}

func parseInt(s string) int {
	value, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return value
}
