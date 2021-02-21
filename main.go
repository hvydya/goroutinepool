package main

import (
	"os"
	"sync"
	"time"

	"github.com/hvydya/goroutinepool/grpool"
)

var lock sync.Mutex

func sample() {
	time.Sleep(2 * time.Second)
	lock.Lock()
	defer lock.Unlock()
	f, err := os.OpenFile("./testing.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString("text\n"); err != nil {
		panic(err)
	}
}

func main() {
	pool := grpool.CreateExecutor(5, 1000)
	for i := 0; i < 20; i++ {
		pool.Submit(sample)
	}
	time.Sleep(7 * time.Second)
	pool.Shutdown()
	// fmt.Printf("size: %d\n", pool.TaskQueue.Size())
}
