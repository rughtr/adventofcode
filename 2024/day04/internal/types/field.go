package types

import (
	"bufio"
	"strings"
)

type Field struct {
	Value     rune
	Relevance int
}

func FieldFromRune(value rune) Field {
	return Field{
		Value:     value,
		Relevance: 0,
	}
}

func ReadFromString(lines string) [][]Field {

	array := make([][]Field, 0)

	reader := strings.NewReader(lines)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		row := make([]Field, 0, len(runes))
		for i := 0; i < len(runes); i++ {
			r := FieldFromRune(runes[i])
			row = append(row, r)
		}
		array = append(array, row)
	}

	return array
}

func (f *Field) IncreaseRelevance() {
	f.Relevance++
}
