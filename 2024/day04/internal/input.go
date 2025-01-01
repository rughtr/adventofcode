package internal

import (
	"aoc24-04/internal/types"
	"bufio"
	"os"
)

func LoadInput(f *os.File) [][]types.Field {

	fields := make([][]types.Field, 0)

	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		row := make([]types.Field, 0)
		runes := []rune(line)
		for i := 0; i < len(runes); i++ {
			r := types.Field{runes[i], 0}
			row = append(row, r)
		}
		fields = append(fields, row)
	}

	return fields
}
