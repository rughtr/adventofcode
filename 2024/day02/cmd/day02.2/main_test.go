package main

import (
	"aoc24-02/internal"
	shared "aoc24-shared"
	"fmt"
	"slices"
	"strings"
	"testing"
)

type testCase struct {
	r              internal.Report
	expectedResult bool
}

func Test_Report_isSafe_with_dampener_returns_expected_value(t *testing.T) {

	cases := []testCase{
		{r: internal.NewReport([]int{7, 6, 4, 2, 1}), expectedResult: true},
		{r: internal.NewReport([]int{1, 2, 7, 8, 9}), expectedResult: false},
		{r: internal.NewReport([]int{9, 7, 6, 2, 1}), expectedResult: false},
		{r: internal.NewReport([]int{1, 3, 2, 4, 5}), expectedResult: true},
		{r: internal.NewReport([]int{8, 6, 4, 4, 1}), expectedResult: true},
		{r: internal.NewReport([]int{1, 3, 6, 7, 9}), expectedResult: true},
		{r: internal.NewReport([]int{55, 58, 59, 57, 60, 61, 68}), expectedResult: false},
		{r: internal.NewReport([]int{71, 69, 70, 71, 72, 75}), expectedResult: true},
	}

	for i, tc := range cases {
		levels := strings.Join(slices.Collect(shared.Map(tc.r.Levels, func(value int) string {
			return fmt.Sprintf("%d", value)
		})), "_")
		t.Run(fmt.Sprintf("isSafe_with_levels_%s_returns_%v", levels, tc.expectedResult), func(t *testing.T) {
			_, _, result := isSafe(tc.r)
			if result != tc.expectedResult {
				t.Errorf("Test case %d failed: input = %v, expected = %v, got = %v", i+1, tc.r, tc.expectedResult, result)
			}
		})
	}
}
