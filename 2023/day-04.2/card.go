package main

type cardRecord struct {
	Number         int
	WinningNumbers map[int]struct{}
	Numbers        []int
}

func newCardRecord(number int, winningNumbers []int, numbers []int) *cardRecord {
	winningNumbersSet := make(map[int]struct{})
	for _, n := range winningNumbers {
		winningNumbersSet[n] = struct{}{}
	}
	return &cardRecord{
		Number:         number,
		WinningNumbers: winningNumbersSet,
		Numbers:        numbers,
	}
}

func (r *cardRecord) GetNumMatches() int {
	result := 0
	for _, number := range r.Numbers {
		_, isWinner := r.WinningNumbers[number]
		if isWinner {
			result++
		}
	}
	return result
}
