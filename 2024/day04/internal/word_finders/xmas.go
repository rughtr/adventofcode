package word_finders

import (
	"aoc24-04/internal/types"
)

type xmasFinder struct {
	BaseFinder
}

var _ types.WordFinder = (*xmasFinder)(nil)

func NewXmasFinder(word []rune, array [][]types.Field) types.WordFinder {
	return &xmasFinder{
		BaseFinder: NewBaseFinder(word, array),
	}
}

func (wf *xmasFinder) Find() int {

	result := 0

	for y := 0; y < len(wf.array); y++ {
		for x := 0; x < len(wf.array[y]); x++ {

			runeAt := wf.array[y][x].Value

			first := wf.word[0]
			if runeAt == first {
				result += wf.findWord(x, y, types.HorizontalForward)
				result += wf.findWord(x, y, types.HorizontalReverse)
				result += wf.findWord(x, y, types.VerticalDown)
				result += wf.findWord(x, y, types.VerticalUp)
				result += wf.findWord(x, y, types.DiagonalLtr)
				result += wf.findWord(x, y, types.DiagonalLtrReverse)
				result += wf.findWord(x, y, types.DiagonalRtl)
				result += wf.findWord(x, y, types.DiagonalRtlReverse)
			}
		}
	}

	return result
}

func (wf *xmasFinder) findWord(x int, y int, dir types.Direction) int {

	word := wf.word
	wordLen := len(word)

	dx := dir.X()
	dy := dir.Y()

	if x < 0 || y < 0 || y >= len(wf.array) || x >= len(wf.array[y]) {
		return 0
	}

	// Check bounds (depending on direction)
	xb, yb := x+dx*(wordLen-1), y+dy*(wordLen-1)
	if xb < 0 || xb >= len(wf.array[y]) || yb < 0 || yb >= len(wf.array) {
		return 0
	}

	for i := 0; i < wordLen; i++ {
		runeAt := wf.array[y+dy*i][x+dx*i].Value
		if runeAt != word[i] {
			return 0
		}
	}

	for i := 0; i < wordLen; i++ {
		wf.array[y+dy*i][x+dx*i].Relevance++
	}

	return 1
}
