package main

import (
	"aoc24-04/internal"
	"aoc24-04/internal/word_finders"
	shared "aoc24-shared"
	"fmt"
	"os"
)

func main() {

	f := shared.OpenInput()
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	fields := internal.LoadInput(f)

	finder := word_finders.NewXmasFinder([]rune("XMAS"), fields)
	count := finder.Find()

	finder.Dump()
	fmt.Printf("Num words: %d", count)
}
