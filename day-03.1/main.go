package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"regexp"
	"strconv"
)

func main() {

	cwd, _ := os.Getwd()
	inputPath := path.Join(cwd, "input.txt")
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		panic("the input file does not exist")
	}

	grid := make([][]rune, 0)

	var sum = 0

	isAdjacent := func(line int, from int, to int) bool {
		for x := line - 1; x < line+2; x++ {
			if x < 0 {
				continue
			}
			for y := from - 1; y < to+1; y++ {
				if y < 0 || (x == line && y >= from && y < to) {
					continue
				}
				if x < len(grid) {
					record := grid[x]
					if y < len(record) {
						r := record[y]
						if r != '.' {
							return true
						}
					}
				}
			}
		}
		return false
	}

	file, _ := os.Open(inputPath)
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		next := reader.Text()
		runes := []rune(next)
		grid = append(grid, runes)
	}

	r := regexp.MustCompile("\\d+")
	for lineNumber, line := range grid {
		s := string(line)
		indices := r.FindAllStringSubmatchIndex(s, -1)
		for _, p := range indices {
			from := p[0]
			to := p[1]
			if isAdjacent(lineNumber, from, to) {
				numberString := s[from:to]
				value, _ := strconv.Atoi(numberString)
				sum += value
			}
		}
	}

	fmt.Printf("%d\n", sum)

	os.Exit(0)
}
