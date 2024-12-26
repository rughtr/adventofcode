package main

import (
	"aoc24-01/internal"
	"fmt"
)

func main() {

	left, right := internal.LoadListsFromInput()

	similarity := solve(left, right)

	fmt.Printf("Similarity score: %d\n", similarity)
}

func solve(left []int, right []int) uint64 {

	set := make(map[int]int)
	for _, v := range right {
		counter, ok := set[v]
		if ok {
			set[v] = counter + 1
		} else {
			set[v] = 1
		}
	}

	var similarity = uint64(0)

	for _, v := range left {
		counter, _ := set[v]
		score := v * counter
		similarity += uint64(score)
	}

	return similarity
}
