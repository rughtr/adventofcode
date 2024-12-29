package main

import (
	shared "aoc24-shared"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func main() {

	f := shared.OpenInput()
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	bytes, _ := io.ReadAll(f)

	const pattern = "mul\\(([1-9]\\d{0,2}),([1-9]\\d{0,2})\\)"
	r := regexp.MustCompile(pattern)

	const all = -1
	indices := r.FindAllSubmatchIndex(bytes, all)

	base := 10
	size := 64
	sum := uint64(0)

	for _, i := range indices {
		a, _ := strconv.ParseInt(string(bytes[i[2]:i[3]]), base, size)
		b, _ := strconv.ParseInt(string(bytes[i[4]:i[5]]), base, size)
		p := a * b
		sum += uint64(p)
	}

	fmt.Println(sum)
}