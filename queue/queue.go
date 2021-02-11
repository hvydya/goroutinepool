package queue

import (
	"sync"
	"time"
)

// Queue is an implementation of LinkedBlockingQueue
type Queue struct {
	mu       sync.Mutex
	capacity int
	q        []interface{}
}

// BlockingQueue link here
type BlockingQueue interface {
	Insert()
	Remove()
	Peek()
	Size()
	Capacity()
}

// Insert inserts the item into the queue
func (q *Queue) Insert(item interface{}) error {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.q) < int(q.capacity) {
		q.q = append(q.q, item)
		return nil
	}
	return &QueueError{
		time.Now(),
		"Queue at max capacity",
		"Insert",
	}
}

// Remove removes the oldest element from the queue
func (q *Queue) Remove() (interface{}, error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.q) > 0 {
		item := q.q[0]
		q.q = q.q[1:]
		return item, nil
	}
	return nil, &QueueError{
		time.Now(),
		"Queue is empty",
		"Remove",
	}
}

// Peek returns the oldest element without removing from the queue
func (q *Queue) Peek() (interface{}, error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	qlen := len(q.q)
	if qlen > 0 {
		return q.q[0], nil
	}
	return nil, &QueueError{
		time.Now(),
		"Queue is empty",
		"Peek",
	}
}

// Size returns the current size of the queue
func (q *Queue) Size() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.q)
}

// Capacity returns the capacity of the queue
func (q *Queue) Capacity() int {
	return q.capacity
}

// CreateQueue creates an empty queue with desired capacity
func CreateQueue(capacity int) *Queue {
	return &Queue{
		capacity: capacity,
		q:        make([]interface{}, 0, capacity),
	}
}
