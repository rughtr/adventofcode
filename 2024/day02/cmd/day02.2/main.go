package main

import (
	"aoc24-02/internal"
	"fmt"
	"math"
)

func main() {

	reports := internal.LoadReportsFromInput()
	// dumpReports("dump.txt", reports)

	numSaveReports := solve(reports)

	fmt.Printf("Num save reports: %d\n", numSaveReports)
}

func solve(reports []internal.Report) int {

	counter := 0
	for _, report := range reports {
		if _, _, ok := isSafe(report); ok {
			counter++
		}
	}

	return counter
}

func isSafe(r internal.Report) (*internal.Report, string, bool) {
	var reason string
	if isSafeReport(r, &reason) {
		return &r, reason, true
	} else {
		for i := 0; i < len(r.Levels); i++ {
			leading := r.Levels[0:i]
			trailing := r.Levels[i+1 : len(r.Levels)]
			modified := make([]int, 0)
			modified = append(modified, leading...)
			modified = append(modified, trailing...)
			modifiedReport := internal.Report{Levels: modified}
			if isSafeReport(modifiedReport, &reason) {
				reason = fmt.Sprintf("dampened %d at %d", r.Levels[i], i)
				return &modifiedReport, reason, true
			}
		}
		return nil, "unsafe", false
	}
}

func isIncreasing(r internal.Report) bool {
	incr, decr := 0, 0
	last := r.Levels[0]
	for _, v := range r.Levels[1:] {
		if v > last {
			incr++
		} else if v < last {
			decr++
		}
		last = v
	}
	return incr > decr
}

func isSafeReport(r internal.Report, reason *string) bool {

	if len(r.Levels) < 2 {
		panic("not enough levels")
	}

	sameLevel := func(a, b int) bool {
		return a == b
	}

	increasing := isIncreasing(r)

	for i := 1; i < len(r.Levels); i++ {
		v := r.Levels[i]
		prev := r.Levels[i-1]
		diff := int(math.Abs(float64(v - prev)))
		if sameLevel(prev, v) || diff > 3 {
			*reason = fmt.Sprintf("diff %d > 3 between %d and %d", diff, prev, v)
			return false
		}
		if increasing && v < prev {
			*reason = fmt.Sprintf("increasing and %d < %d", v, prev)
			return false
		} else if !increasing && v > prev {
			*reason = fmt.Sprintf("not increasing and %d > %d", v, prev)
			return false
		}
		prev = v
	}

	return true
}
