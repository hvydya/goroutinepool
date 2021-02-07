package goroutinepool

import (
	"fmt"
	"time"
)

// QInsertError is thrown when Insert is called on the Queue and some error occurs
type QInsertError struct {
	When time.Time
	What string
}

func (e *QInsertError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}
