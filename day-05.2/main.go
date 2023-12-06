package main

import (
	"fmt"
	"os"
	"path"
	"sync"
)

func main() {

	cwd, _ := os.Getwd()
	inputPath := path.Join(cwd, "input.txt")
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		panic("the input file does not exist")
	}

	a := buildAlmanacFromFile(inputPath)

	seedRanges := a.getSeedRanges()

	result := int64(0)
	locations := make([]int64, len(seedRanges))

	wg := sync.WaitGroup{}
	m := sync.Mutex{}

	for index, seedRange := range seedRanges {
		wg.Add(1)
		findLocation := func(index int, from int64, to int64) {
			location := int64(0)
			iter := createSeedIter(from, to)
			for iter.Next() {
				seed := iter.Current()
				info := a.resolveSeedInfo(seed)
				if location == 0 || location > info.Location {
					location = info.Location
				}
			}
			storeLocation := func(value int64) {
				m.Lock()
				defer m.Unlock()
				locations[index] = value
			}
			storeLocation(location)
			wg.Done()
		}
		go findLocation(index, seedRange[0], seedRange[1])
	}
	wg.Wait()

	for _, n := range locations {
		if result == 0 || n < result {
			result = n
		}
	}

	fmt.Printf("Lowest location: %d\n", result)

	os.Exit(0)
}
