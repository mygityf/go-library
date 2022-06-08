package queue

import (
	"container/list"
	"sync"
)

// queue
type queueImpl struct {
	list *list.List
	mu   sync.RWMutex
}

// NewQueuePtr new
func NewQueuePtr() *queueImpl {
	return &queueImpl{
		list: list.New(),
	}
}

// PushFront push front
func (q *queueImpl) PushFront(value interface{}) {
	q.mu.Lock()
	q.list.PushFront(value)
	q.mu.Unlock()
}

// PushBack push back
func (q *queueImpl) PushBack(value interface{}) {
	q.mu.Lock()
	q.list.PushBack(value)
	q.mu.Unlock()
}

// PopFront pop front
func (q *queueImpl) PopFront() interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()
	e := q.list.Front()
	if e != nil {
		q.list.Remove(e)
		return e.Value
	}
	return nil
}

// PopBack pop back
func (q *queueImpl) PopBack() interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()
	e := q.list.Front()
	if e != nil {
		q.list.Remove(e)
		return e.Value
	}
	return nil
}

// PeakFront peak front
func (q *queueImpl) PeakFront() interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()
	e := q.list.Front()
	if e != nil {
		return e.Value
	}

	return nil
}

// PeakBack of queue
func (q *queueImpl) PeakBack() interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()
	e := q.list.Back()
	if e != nil {
		return e.Value
	}

	return nil
}

// Len of queue
func (q *queueImpl) Len() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.list.Len()
}

// IsEmpty of queue
func (q *queueImpl) IsEmpty() bool {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.list.Len() == 0
}

// Visit by action
func (q *queueImpl) Visit(action func(ele *list.Element)) {
	if action == nil {
		return
	}
	q.mu.Lock()
	defer q.mu.Unlock()
	for e := q.list.Front(); e != nil; e = e.Next() {
		action(e)
	}
	return
}

// Remove elements
func (q *queueImpl) Remove(elements []*list.Element) {
	q.mu.Lock()
	defer q.mu.Unlock()
	for _, element:= range elements {
		q.list.Remove(element)
	}
}
