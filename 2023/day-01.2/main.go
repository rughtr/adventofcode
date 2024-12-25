package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
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

	numberWords := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}

	firstDigit := func(s string, direction int) int {
		runes := []rune(s)
		length := len(runes)
		for i := 0; i < length; i++ {
			for word, digit := range numberWords {
				wordLength := len(word)
				var selection string
				if direction < 0 && i >= wordLength {
					selection = string(runes[length-i : length-i+wordLength])
				} else if direction >= 0 && i+wordLength < length {
					selection = string(runes[i : i+wordLength])
				}
				if strings.Compare(selection, word) == 0 {
					return digit
				}
			}
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
