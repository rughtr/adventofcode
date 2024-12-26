package main

import (
	"aoc24-01/internal"
	"fmt"
	"sort"
)

func main() {

	left, right := internal.LoadListsFromInput()

	distance := solve(left, right)

	fmt.Printf("Distance: %d\n", distance)
}

func solve(left []int, right []int) uint64 {

	sort.Ints(left)
	sort.Ints(right)

	var distance = uint64(0)

	for i := 0; i < len(left) && i < len(right); i++ {
		a := left[i]
		b := right[i]
		d := a - b
		if a < b {
			d = b - a
		}
		distance += uint64(d)
	}

	return distance
}
