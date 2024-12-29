package main

import (
	"aoc24-02/internal"
	"fmt"
	"math"
)

func main() {

	reports := internal.LoadReportsFromInput()

	numSaveReports := solve(reports)

	fmt.Printf("Num save reports: %d\n", numSaveReports)
}

func solve(reports []internal.Report) int {

	counter := 0
	for _, report := range reports {
		if isSafe(report) {
			counter++
		}
	}

	return counter
}

func isSafe(r internal.Report) bool {

	if len(r.Levels) < 2 {
		panic("not enough r")
	}

	sameLevel := func(a, b int) bool {
		return a == b
	}

	prev := r.Levels[0]
	increasing := r.Levels[1] > prev
	for _, v := range r.Levels[1:] {
		diff := int(math.Abs(float64(v - prev)))
		if sameLevel(prev, v) || diff > 3 {
			return false
		}
		if increasing && v < prev {
			return false
		} else if !increasing && v > prev {
			return false
		}
		prev = v
	}

	return true
}
