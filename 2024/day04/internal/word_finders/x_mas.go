package word_finders

import (
	"aoc24-04/internal/types"
	shared "aoc24-shared"
	"slices"
)

type xMasFinder struct {
	BaseFinder
	wordLength int
}

func (wf *xMasFinder) findWord(x int, y int, d types.Direction) bool {

	dx := d.X()
	dy := d.Y()
	for i := 0; i < wf.wordLength; i++ {
		nx, ny := x+dx*i, y+dy*i
		grid := wf.array
		r := wf.word[i]
		if nx < 0 || ny < 0 || ny >= len(grid) || nx >= len(grid[ny]) || grid[ny][nx].Value != r {
			return false
		}
	}

	return true
}

func (wf *xMasFinder) markRelevantFields(x int, y int, d types.Direction) {
	dx := d.X()
	dy := d.Y()
	for i := 0; i < wf.wordLength; i++ {
		nx, ny := x+dx*i, y+dy*i
		wf.array[ny][nx].Relevance++
	}
}

func (wf *xMasFinder) Find() int {

	result := 0

	type check struct {
		x, y int
		d    types.Direction
	}

	directions := []types.Direction{types.DiagonalLtr, types.DiagonalRtl, types.DiagonalLtrReverse, types.DiagonalRtlReverse}
	checks := slices.Collect(shared.Map(directions, func(d types.Direction) check {
		return check{x: d.X() * -1, y: d.Y() * -1, d: d}
	}))

	const expectedMatches = 2

	for y := 1; y < len(wf.array)-1; y++ {
		for x := 1; x < len(wf.array[y])-1; x++ {

			matches := make([]check, 0, expectedMatches)
			for _, c := range checks {
				if wf.findWord(x+c.x, y+c.y, c.d) {
					matches = append(matches, c)
				}
			}

			if len(matches) == cap(matches) {
				result++
				for _, m := range matches {
					wf.markRelevantFields(x+m.x, y+m.y, m.d)
				}
			}
		}
	}

	return result
}

var _ types.WordFinder = (*xMasFinder)(nil)

func NewXMasFinder(word []rune, array [][]types.Field) types.WordFinder {
	if len(word)%3 != 0 {
		panic("word length must be a multiple of 3")
	}
	return &xMasFinder{
		BaseFinder: NewBaseFinder(word, array),
		wordLength: len(word),
	}
}
