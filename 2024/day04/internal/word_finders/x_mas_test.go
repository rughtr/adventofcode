package word_finders

import (
	"aoc24-04/internal/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_XMasFinder_Find_all_x_mas_words_from_sample(t *testing.T) {

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

	const word = "MAS"
	sut := NewXMasFinder([]rune(word), array)

	// Act
	actual := sut.Find()

	expected := "" +
		".M.S......" + "\n" +
		"..A..MSMS." + "\n" +
		".M.S.MAA.." + "\n" +
		"..A.ASMSM." + "\n" +
		".M.S.M...." + "\n" +
		".........." + "\n" +
		"S.S.S.S.S." + "\n" +
		".A.A.A.A.." + "\n" +
		"M.M.M.M.M." + "\n" +
		".........." + "\n"

	dump := sut.Dump()

	// Assert
	assert.Equal(t, 9, actual)
	assert.Equal(t, expected, dump)
}
