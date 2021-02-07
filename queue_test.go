package goroutinepool

import (
	"reflect"
	"testing"
)

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
		q.Insert(v)
		arr = append(arr, v)
		if q.Size() != len(arr) || !reflect.DeepEqual(q.q, arr) {
			t.Errorf("Queue must return size %v and must contain %v. But found size %v and items %v", len(arr), arr, q.Size(), q.q)
		}
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
		item := q.Remove()
		expectedItem := arr[0]
		arr = arr[1:]
		if item != expectedItem || q.Size() != len(arr) {
			t.Errorf("Queue must return removed %v and size must be %v. But found %v and %v", expectedItem, len(arr), item, q.Size())
		}
	}
}

func TestPeekNoConcurrency(t *testing.T) {
	capacity := 5
	q := CreateQueue(capacity)
	testItems := []interface{}{1, 2, 3, 4, 5}
	for _, v := range testItems {
		q.Insert(v)
		arr := []interface{}{v}
		if q.Peek() != v || !reflect.DeepEqual(q.q, arr) {
			t.Errorf("Queue must return peek %v and must contain %v. But found peek %v and items %v", v, arr, q.Peek(), q.q)
		}
		q.Remove()
	}
}
