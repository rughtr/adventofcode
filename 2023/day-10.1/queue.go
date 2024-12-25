package main

import (
	"container/list"
	"sync"
)

type Queue[TValue any] struct {
	list *list.List
	m    sync.RWMutex
}

func MakeQueue[TValue any]() *Queue[TValue] {
	return &Queue[TValue]{
		list: list.New(),
		m:    sync.RWMutex{},
	}
}

func (q *Queue[TValue]) Any() bool {
	q.m.RLock()
	defer q.m.RUnlock()
	return q.list.Len() > 0
}

func (q *Queue[TValue]) Enqueue(value TValue) {
	q.m.Lock()
	defer q.m.Unlock()
	q.list.PushBack(value)
}

func (q *Queue[TValue]) Dequeue() TValue {
	q.m.Lock()
	defer q.m.Unlock()
	if q.list.Len() == 0 {
		panic("no more items")
	}
	elem := q.list.Front()
	q.list.Remove(elem)
	value, _ := elem.Value.(TValue)
	return value
}
