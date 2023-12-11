package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

type intList []int

func (s intList) all(cb func(x int) bool) bool {
	if len(s) == 0 {
		return true
	}
	b := cb(s[0])
	for i := 1; b && i < len(s); i++ {
		if cb(s[i]) != b {
			return false
		}
	}
	return b
}

func collectPlaceholders(numbers intList) intList {
	placeholders := make(intList, 0)
	n := numbers[0]
	for i := 1; i < len(numbers); i++ {
		diff := numbers[i] - n
		placeholders = append(placeholders, diff)
		n = numbers[i]
	}
	return placeholders
}

func (s intList) extrapolateNextValue() int {

	calculationStack := MakeStack[intList]()
	calculationStack.Push(s)

	collectionStack := MakeStack[intList]()
	collectionStack.Push(s)

	for collectionStack.Any() {
		numbers := collectionStack.Pop()
		if numbers.all(func(x int) bool {
			return x == 0
		}) {
			break
		}
		placeholders := collectPlaceholders(numbers)
		collectionStack.Push(placeholders)
		calculationStack.Push(placeholders)
	}

	result := 0
	for calculationStack.Any() {
		p0 := calculationStack.Pop()
		result = p0[len(p0)-1]
		p1, ok := calculationStack.TryPeek()
		if ok {
			p1 = calculationStack.Pop()
			extrapolated := p1[len(p1)-1] + result
			p1 = append(p1, extrapolated)
			calculationStack.Push(p1)
		}
	}

	return result
}

func readSequences(filename string) []intList {

	sequences := make([]intList, 0)

	file, _ := os.Open(filename)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	reader := bufio.NewScanner(file)
	for reader.Scan() {
		line := reader.Text()
		numberStrings := strings.Split(line, " ")
		numbers := make([]int, 0)
		for _, next := range numberStrings {
			number, err := strconv.Atoi(next)
			if err != nil {
				panic("cannot parse int")
			}
			numbers = append(numbers, number)
		}
		sequences = append(sequences, numbers)
	}

	return sequences
}

func main() {

	cwd, _ := os.Getwd()
	inputPath := path.Join(cwd, "input.txt")
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		panic("the input file does not exist")
	}

	sum := 0
	sequences := readSequences(inputPath)
	for _, numbers := range sequences {
		extrapolated := numbers.extrapolateNextValue()
		sum += extrapolated
	}

	fmt.Printf("%d\n", sum)
}
