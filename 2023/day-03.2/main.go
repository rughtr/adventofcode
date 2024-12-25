package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"regexp"
	"strconv"
)

type gearId struct {
	X int
	Y int
}

func main() {

	cwd, _ := os.Getwd()
	inputPath := path.Join(cwd, "input.txt")
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		panic("the input file does not exist")
	}

	grid := make([][]rune, 0)

	var sum = 0

	isAdjacentWith := func(line int, from int, to int, chr rune) (bool, gearId) {
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
						if r == chr {
							return true, gearId{x, y}
						}
					}
				}
			}
		}
		return false, gearId{-1, -1}
	}

	file, _ := os.Open(inputPath)
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		next := reader.Text()
		runes := []rune(next)
		grid = append(grid, runes)
	}

	const GearChar = '*'

	possibleGears := make(map[gearId][]int)

	r := regexp.MustCompile("\\d+")
	for lineNumber, line := range grid {
		s := string(line)
		indices := r.FindAllStringSubmatchIndex(s, -1)
		for _, p := range indices {
			from := p[0]
			to := p[1]
			connectWithGear, gearId := isAdjacentWith(lineNumber, from, to, GearChar)
			if connectWithGear {
				numberString := s[from:to]
				value, _ := strconv.Atoi(numberString)
				pieces, found := possibleGears[gearId]
				if found {
					pieces = append(pieces, value)
				} else {
					pieces = []int{value}
				}
				possibleGears[gearId] = pieces
			}
		}
	}

	for _, pieces := range possibleGears {
		if len(pieces) == 2 {
			gearRatio := 1
			for _, value := range pieces {
				gearRatio *= value
			}
			sum += gearRatio
		}
	}

	fmt.Printf("%d\n", sum)

	os.Exit(0)
}
