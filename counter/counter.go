package counter

import (
	"sync"
)

// AtomicCounter is a thread safe counter
type AtomicCounter struct {
	mu      sync.Mutex
	counter uint64
}

// Counter is a generic interface for counter
type Counter interface {
	Increment()
	Decrement()
	Get() uint64
}

// Increment the counter by 1 unit
func (c *AtomicCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter++
}

// Decrement the counter by 1 unit
func (c *AtomicCounter) Decrement() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.counter > 0 {
		c.counter--
	}
}

// Get the current value of the counter
func (c *AtomicCounter) Get() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.counter
}

// CreateAtomicCounter creates it
func CreateAtomicCounter() *AtomicCounter {
	return &AtomicCounter{
		counter: 0,
	}
}
