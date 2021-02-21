package grpool

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/hvydya/goroutinepool/counter"
	"github.com/hvydya/goroutinepool/queue"
)

// ExecutorError describes a general error in an executor
type ExecutorError struct {
	When time.Time
	What string
	Func string
}

func (e *ExecutorError) Error() string {
	return fmt.Sprintf("at %v, due to %s in %s", e.When, e.What, e.Func)
}

// BasicExecutor is a generic exeutor
type BasicExecutor struct {
	mu        sync.Mutex
	TaskQueue queue.Queue
	poolSize  int
	Running   *counter.AtomicCounter
	shutdown  context.CancelFunc
}

// Task represents any task that needs to be parallelized
type Task func()

// Executor is a generic interface of a goroutinepool
type Executor interface {
	Submit(Task) error
	start(*context.Context)
	Shutdown()
}

// Submit a task to the executor
func (exe *BasicExecutor) Submit(task Task) error {
	if err := exe.TaskQueue.Insert(task); err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func (exe *BasicExecutor) start(ctx *context.Context) {
	for {
		select {
		case <-(*ctx).Done():
			return
		default:
			if exe.TaskQueue.Size() > 0 && exe.Running.Get() < uint64(exe.poolSize) {
				funcInterface, err := exe.TaskQueue.Remove()
				if err == nil {
					run, ok := funcInterface.(Task)
					if ok {
						exe.Running.Increment()
						go func() {
							run()
							exe.Running.Decrement()
						}()
					} else {
						err := &ExecutorError{
							When: time.Now(),
							What: "Couldn't cast queue item to Task type",
							Func: "start",
						}

						fmt.Println(err)
					}
				} else {
					err := errors.New(err.Error())
					fmt.Println(err)
				}
			}
		}
	}
}

// Shutdown the goroutinepool
func (exe *BasicExecutor) Shutdown() {
	exe.shutdown()
}

// CreateExecutor creates a goroutinepool
func CreateExecutor(poolSize, queueCapacity int) Executor {
	queue := queue.CreateQueue(queueCapacity)
	context, cancel := context.WithCancel(context.Background())
	var executor Executor = &BasicExecutor{
		TaskQueue: queue,
		poolSize:  poolSize,
		Running:   counter.CreateAtomicCounter(),
		shutdown:  cancel,
	}
	go executor.start(&context)
	return executor
}

// TODO : make this unbounded queue.

// CreateFixedPool creates a pool with specified poolSize and fixed queue size of 100.
func CreateFixedPool(poolSize int) *BasicExecutor {
	queue := queue.CreateQueue(1000)
	context, cancel := context.WithCancel(context.Background())
	executor := &BasicExecutor{
		TaskQueue: queue,
		poolSize:  poolSize,
		Running:   counter.CreateAtomicCounter(),
	}
	executor.shutdown = cancel
	go executor.start(&context)
	return executor
}
