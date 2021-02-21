package queue

import "time"

// CQueue : Channel Queue
type CQueue struct {
	capacity int
	q        chan interface{}
}

// SimpleFIFOQueue supports only the basic operations
type SimpleFIFOQueue interface {
	Insert(interface{}) error
	Remove() (interface{}, error)
	Size() int
	Capacity() int
}

// Insert inserts the item into queue
func (q *CQueue) Insert(item interface{}) error {
	if len(q.q) < q.capacity {
		q.q <- item
		return nil
	}
	return &Error{
		time.Now(),
		"CQueue at max capacity",
		"Insert",
	}
}

// Remove removes the oldest inserted element
func (q *CQueue) Remove() (interface{}, error) {
	if len(q.q) > 0 {
		return <-q.q, nil
	}
	return nil, &Error{
		time.Now(),
		"CQueue empty. Can't remove",
		"Remove",
	}
}

// Size returns the size of queue
func (q *CQueue) Size() int {
	return len(q.q)
}

// Capacity returns the allotted capacity for the queue
func (q *CQueue) Capacity() int {
	return q.capacity
}

// CreateCQueue creates a channel queue
func CreateCQueue(capacity int) SimpleFIFOQueue {
	return &CQueue{
		capacity: capacity,
		q:        make(chan interface{}, capacity),
	}
}
