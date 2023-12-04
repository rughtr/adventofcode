package main

import "sync"

type cardStats struct {
	counters map[int]int
	m        sync.Mutex
}

func newScratchCardStats() *cardStats {
	return &cardStats{
		counters: make(map[int]int),
		m:        sync.Mutex{},
	}
}

func (s *cardStats) CountScratchCard(card int) {
	s.m.Lock()
	defer s.m.Unlock()
	counter, _ := s.counters[card]
	counter++
	s.counters[card] = counter
}

func (s *cardStats) ScratchCardsTotal() int {
	result := 0
	for _, counter := range s.counters {
		result += counter
	}
	return result
}
