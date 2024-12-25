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

	a := buildAlmanacFromFile(inputPath)
	seedInfos := a.resolveAllSeedInfos()

	location := int64(0)
	for _, info := range seedInfos {
		if location == 0 || location > info.Location {
			location = info.Location
		}
	}

	fmt.Printf("Lowest location: %d\n", location)

	os.Exit(0)
}
