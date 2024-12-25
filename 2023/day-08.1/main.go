package main

import (
	"fmt"
	"os"
	"path"
)

func main() {

	cwd, _ := os.Getwd()
	inputPath := path.Join(cwd, "input.txt")
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		panic("the input file does not exist")
	}

	m := createNavigator(inputPath)
	steps := 0
	for m.NavigateNext() {
		steps++
	}
	fmt.Printf("Steps: %d\n", steps)

	os.Exit(0)

}
