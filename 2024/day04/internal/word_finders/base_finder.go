package word_finders

import (
	"aoc24-04/internal/types"
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

type BaseFinder struct {
	word  []rune
	array [][]types.Field
}

func NewBaseFinder(word []rune, array [][]types.Field) BaseFinder {
	return BaseFinder{
		word:  word,
		array: array,
	}
}

func (wf *BaseFinder) Dump() string {
	buffer := strings.Builder{}
	writer := bufio.NewWriter(&buffer)
	for _, row := range wf.array {
		for _, field := range row {
			if field.Relevance > 0 {
				_, _ = writer.WriteString(string(field.Value))
			} else {
				_, _ = writer.WriteString(".")
			}
		}
		_, _ = writer.WriteString("\n")
	}
	_ = writer.Flush()
	return buffer.String()
}

func (wf *BaseFinder) DumpTo(file string) {

	wd, _ := os.Getwd()
	path := filepath.Join(wd, file)

	f, _ := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	dump := wf.Dump()
	_, _ = f.WriteString(dump)
}
