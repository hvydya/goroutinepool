package goroutinepool

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

// TODO : Write tests for concurrency testing of the queue

func TestCreateQueue(t *testing.T) {
	capacity := 5
	q := CreateQueue(capacity)
	if q.Size() != 0 {
		t.Errorf("Queue must have size of zero when initialized. But found %v", q.Size())
	}
	if q.Capacity() != capacity {
		t.Errorf("Queue must have capacity of %v but found %v", capacity, q.Capacity())
	}
}

func TestCapacity(t *testing.T) {
	capacity := 5
	q := CreateQueue(capacity)
	if q.Size() != 0 || q.Capacity() != capacity {
		t.Errorf("Queue must have size 0 and capacity %v. But found %v and %v", capacity, q.Size(), q.Capacity())
	}
}

func TestSizeNoConcurrency(t *testing.T) {
	capacity := 5
	q := CreateQueue(capacity)
	testItems := []interface{}{1, 2, 3, 4, 5}
	arr := make([]interface{}, 0, capacity)
	for _, v := range testItems {
		q.Insert(v)
		arr = append(arr, v)
		if q.Size() != len(arr) || !reflect.DeepEqual(q.q, arr) {
			t.Errorf("Queue must return size %v and must contain %v. But found size %v and items %v", len(arr), arr, q.Size(), q.q)
		}
	}
}

func TestInsertNoConcurrency(t *testing.T) {
	capacity := 5
	q := CreateQueue(capacity)
	testItems := []interface{}{1, 2, 3, 4, 5}
	arr := make([]interface{}, 0, capacity)
	for _, v := range testItems {
		err := q.Insert(v)
		arr = append(arr, v)
		if err != nil || q.Size() != len(arr) || !reflect.DeepEqual(q.q, arr) {
			t.Errorf("Queue must return size %v and must contain %v. But found size %v and items %v", len(arr), arr, q.Size(), q.q)
		}
	}
}

func TestInsertErrorNoConcurrency(t *testing.T) {
	capacity := 0
	q := CreateQueue(capacity)
	err := q.Insert(1)
	if err == nil {
		t.Errorf("Queue must return error but found nil")
	}
}

func TestRemoveNoConcurrency(t *testing.T) {
	capacity := 5
	q := CreateQueue(capacity)
	testItems := []interface{}{1, 2, 3, 4, 5}
	arr := make([]interface{}, 0, capacity)
	for _, v := range testItems {
		q.Insert(v)
		arr = append(arr, v)
	}

	for range testItems {
		item, err := q.Remove()
		expectedItem := arr[0]
		arr = arr[1:]
		if err != nil || item != expectedItem || q.Size() != len(arr) {
			t.Errorf("Queue must return removed %v and size must be %v. But found %v and %v", expectedItem, len(arr), item, q.Size())
		}
	}
}

func TestRemoveErrorNoConcurrency(t *testing.T) {
	capacity := 0
	q := CreateQueue(capacity)
	removed, err := q.Remove()
	if err == nil || removed != nil {
		t.Errorf("Queue must throw error when size is 0 and Remove is called")
	}
}

func TestPeekNoConcurrency(t *testing.T) {
	capacity := 5
	q := CreateQueue(capacity)
	testItems := []interface{}{1, 2, 3, 4, 5}
	for _, v := range testItems {
		q.Insert(v)
		arr := []interface{}{v}
		peeked, err := q.Peek()
		if err != nil || peeked != v || !reflect.DeepEqual(q.q, arr) {
			t.Errorf("Queue must return peek %v and must contain %v. But found peek %v and items %v", v, arr, peeked, q.q)
		}
		q.Remove()
	}
}

func TestPeekErrorNoConcurrency(t *testing.T) {
	capacity := 0
	q := CreateQueue(capacity)
	peeked, err := q.Peek()
	if err == nil || peeked != nil {
		t.Errorf("Queue must throw error when size is 0 and Peek is called")
	}
}

func TestQueueError(t *testing.T) {
	now := time.Now()
	errorMsg := "Testing Queue Error"
	funcStr := "TestQueueError"
	msg := fmt.Sprintf("at %v, due to %s in %s", now, errorMsg, funcStr)

	err := &QueueError{
		now,
		errorMsg,
		funcStr,
	}
	if msg != fmt.Sprint(err) {
		t.Errorf("Expected err message to be %s, but found %s", msg, fmt.Sprint(err))
	}
}
