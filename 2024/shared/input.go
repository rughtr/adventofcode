package shared

import (
	"os"
	"path"
)

// OpenInput opens the "input.txt" file in the current working directory and returns it as an *os.File.
// It panics if the file does not exist or if an error occurs while opening the file.
func OpenInput() *os.File {

	cwd, _ := os.Getwd()
	inputPath := path.Join(cwd, "input.txt")
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		panic("the input file does not exist")
	}

	f, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}

	return f
}
