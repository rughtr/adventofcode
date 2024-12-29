package main

import (
	"aoc24-02/internal"
	shared "aoc24-shared"
	"fmt"
	"os"
	"slices"
	"strings"
)

func formatReport(r internal.Report) string {
	return strings.Join(slices.Collect(shared.Map(r.Levels, func(value int) string {
		return fmt.Sprintf("%d", value)
	})), " ")
}

func dumpReports(filename string, reports []internal.Report) {
	f, _ := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	for _, r := range reports {
		_, _ = f.WriteString(formatReport(r))

		safe, reason, ok := isSafe(r)
		if ok && safe != nil {
			_, _ = f.WriteString(", safe: ")
			_, _ = f.WriteString(formatReport(*safe))
		}

		if len(reason) > 0 {
			_, _ = f.WriteString(fmt.Sprintf(", reason: %s", reason))
		}
		_, _ = f.WriteString("\n")
	}
	_ = f.Sync()
}
