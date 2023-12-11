package main

import (
	"sync"
)

type Stack[TValue any] struct {
	values []TValue
	m      sync.RWMutex
}

func MakeStack[TValue any]() Stack[TValue] {
	return Stack[TValue]{
		values: make([]TValue, 0),
		m:      sync.RWMutex{},
	}
}

func (s *Stack[TValue]) Any() bool {
	s.m.RLock()
	defer s.m.RUnlock()
	i := len(s.values)
	return i > 0
}

func (s *Stack[TValue]) Push(value ...TValue) {
	s.m.Lock()
	defer s.m.Unlock()
	s.pushValue(value)
}

func (s *Stack[TValue]) pushValue(value []TValue) {
	s.values = append(s.values, value...)
}

func (s *Stack[TValue]) Pop() TValue {
	s.m.Lock()
	defer s.m.Unlock()
	element := s.popValue()
	return element
}

func (s *Stack[TValue]) popValue() TValue {
	lastItemIndex := len(s.values) - 1
	element := s.values[lastItemIndex]
	s.values = s.values[:lastItemIndex]
	return element
}

func (s *Stack[TValue]) TryPeek() (TValue, bool) {
	s.m.RLock()
	defer s.m.RUnlock()
	return s.tryPeekValue()
}

func (s *Stack[TValue]) tryPeekValue() (TValue, bool) {
	lastItemIndex := len(s.values) - 1
	if lastItemIndex >= 0 {
		element := s.values[lastItemIndex]
		return element, true
	}
	var d TValue
	return d, false
}
