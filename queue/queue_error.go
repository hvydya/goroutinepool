package queue

import (
	"fmt"
	"time"
)

// Error is thrown when any queue function and some error occurs
type Error struct {
	When time.Time
	What string
	Func string
}

func (e *Error) Error() string {
	return fmt.Sprintf("at %v, due to %s in %s", e.When, e.What, e.Func)
}
