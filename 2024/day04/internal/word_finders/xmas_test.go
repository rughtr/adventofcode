package word_finders

import (
	"aoc24-04/internal/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_WordFinder_Find_all_xmas_words_from_sample(t *testing.T) {

	// Arrange
	input := "" +
		"MMMSXXMASM" + "\n" +
		"MSAMXMSMSA" + "\n" +
		"AMXSXMAAMM" + "\n" +
		"MSAMASMSMX" + "\n" +
		"XMASAMXAMM" + "\n" +
		"XXAMMXXAMA" + "\n" +
		"SMSMSASXSS" + "\n" +
		"SAXAMASAAA" + "\n" +
		"MAMMMXMMMM" + "\n" +
		"MXMXAXMASX" + "\n"

	array := types.ReadFromString(input)

	const word = "XMAS"
	sut := NewXmasFinder([]rune(word), array)

	// Act
	actual := sut.Find()

	expected := "" +
		"....XXMAS." + "\n" +
		".SAMXMS..." + "\n" +
		"...S..A..." + "\n" +
		"..A.A.MS.X" + "\n" +
		"XMASAMX.MM" + "\n" +
		"X.....XA.A" + "\n" +
		"S.S.S.S.SS" + "\n" +
		".A.A.A.A.A" + "\n" +
		"..M.M.M.MM" + "\n" +
		".X.X.XMASX" + "\n"

	dump := sut.Dump()

	// Assert
	assert.Equal(t, 18, actual)
	assert.Equal(t, expected, dump)
}
