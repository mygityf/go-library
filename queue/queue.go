package queue

import (
	"container/list"
	"sync"
)

type Queue struct {
	list *list.List
	mu   *sync.RWMutex
}

func NewQueuePtr() *Queue {
	list := list.New()
	return &Queue{
		list: list,
		mu:   &sync.RWMutex{},
	}
}

func (q *Queue) PushFront(value interface{}) {
	q.mu.Lock()
	q.list.PushFront(value)
	q.mu.Unlock()
}

func (q *Queue) PushBack(value interface{}) {
	q.mu.Lock()
	q.list.PushBack(value)
	q.mu.Unlock()
}

func (q *Queue) PopFront() interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()
	e := q.list.Front()
	if e != nil {
		q.list.Remove(e)
		return e.Value
	}
	return nil
}

func (q *Queue) PopBack() interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()
	e := q.list.Front()
	if e != nil {
		q.list.Remove(e)
		return e.Value
	}
	return nil
}

func (q *Queue) PeakFront() interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()
	e := q.list.Front()
	if e != nil {
		return e.Value
	}

	return nil
}

func (q *Queue) PeakBack() interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()
	e := q.list.Back()
	if e != nil {
		return e.Value
	}

	return nil
}
func (q *Queue) Len() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.list.Len()
}

func (q *Queue) Empty() bool {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.list.Len() == 0
}
