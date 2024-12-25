package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"unicode"
)

func main() {

	cwd, _ := os.Getwd()
	inputPath := path.Join(cwd, "input.txt")
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		panic("the input file does not exist")
	}

	var sum = 0

	iif := func(condition bool, trueValue rune, falseValue rune) rune {
		if condition {
			return trueValue
		}
		return falseValue
	}

	firstDigit := func(s string, direction int) int {
		runes := []rune(s)
		length := len(runes)
		for i := 0; i < length; i++ {
			runeValue := iif(direction < 0, runes[length-1-i], runes[i])
			if unicode.IsDigit(runeValue) {
				return int(runeValue - '0')
			}
		}
		return 0
	}

	const FIRST = 1
	const LAST = -1

	file, _ := os.Open(inputPath)
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		next := reader.Text()
		a := firstDigit(next, FIRST)
		b := firstDigit(next, LAST)
		n := (a * 10) + b
		sum += n
	}

	_, _ = fmt.Printf("%d", sum)

	os.Exit(0)
}
