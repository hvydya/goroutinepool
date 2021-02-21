package queue

import (
	"sync"
	"time"
)

// TSQueue : thread safe queue
type TSQueue struct {
	mu       sync.Mutex
	capacity int
	q        []interface{}
}

// Queue link here
type Queue interface {
	Insert(interface{}) error
	Remove() (interface{}, error)
	Peek() (interface{}, error)
	Size() int
	Capacity() int
}

// Insert inserts the item into the queue
func (q *TSQueue) Insert(item interface{}) error {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.q) < int(q.capacity) {
		q.q = append(q.q, item)
		return nil
	}
	return &Error{
		time.Now(),
		"TSQueue at max capacity",
		"Insert",
	}
}

// Remove removes the oldest element from the queue
func (q *TSQueue) Remove() (interface{}, error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.q) > 0 {
		item := q.q[0]
		q.q = q.q[1:]
		return item, nil
	}
	return nil, &Error{
		time.Now(),
		"TSQueue is empty",
		"Remove",
	}
}

// Peek returns the oldest element without removing from the queue
func (q *TSQueue) Peek() (interface{}, error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	qlen := len(q.q)
	if qlen > 0 {
		return q.q[0], nil
	}
	return nil, &Error{
		time.Now(),
		"TSQueue is empty",
		"Peek",
	}
}

// Size returns the current size of the queue
func (q *TSQueue) Size() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.q)
}

// Capacity returns the capacity of the queue
func (q *TSQueue) Capacity() int {
	return q.capacity
}

// CreateQueue creates an empty queue with desired capacity
func CreateQueue(capacity int) Queue {
	return &TSQueue{
		capacity: capacity,
		q:        make([]interface{}, 0, capacity),
	}
}
