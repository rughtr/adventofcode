package main

import (
	"aoc24-02/internal"
	"testing"
)

type testCase struct {
	r              internal.Report
	expectedResult bool
}

func Test_Report_IsSafe_returns_expected_value(t *testing.T) {

	cases := []testCase{
		{r: internal.NewReport([]int{7, 6, 4, 2, 1}), expectedResult: true},
		{r: internal.NewReport([]int{1, 2, 7, 8, 9}), expectedResult: false},
		{r: internal.NewReport([]int{9, 7, 6, 2, 1}), expectedResult: false},
		{r: internal.NewReport([]int{1, 3, 2, 4, 5}), expectedResult: false},
		{r: internal.NewReport([]int{8, 6, 4, 4, 1}), expectedResult: false},
		{r: internal.NewReport([]int{1, 3, 6, 7, 9}), expectedResult: true},
	}

	for i, tc := range cases {
		result := isSafe(tc.r)
		if result != tc.expectedResult {
			t.Errorf("Test case %d failed: input = %v, expected = %v, got = %v", i+1, tc.r, tc.expectedResult, result)
		}
	}
}
