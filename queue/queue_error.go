package queue

import (
	"fmt"
	"time"
)

// QueueError is thrown when any queue function and some error occurs
type QueueError struct {
	When time.Time
	What string
	Func string
}

func (e *QueueError) Error() string {
	return fmt.Sprintf("at %v, due to %s in %s", e.When, e.What, e.Func)
}
