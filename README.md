# goroutinepool

Implementation of Threadpool

### Usage
```go
func myTask() {
    // do something
    // ...
}

func main() {
    // This creates a pool with maximum 5 go routines running your tasks concurrently
    pool := grpool.CreateFixedPool(5)
    for i := 0; i < 10; i++ {
        pool.Submit(myTask)
    }
    // do other stuff
}
```

### TODOS
- More tests
- More methods for the interfaces
- Shutdown must kill all dangling or running go routines when it is called (critical).

### Check coverage
```shell
go test -coverprofile cp.out && go tool cover -html cp.out
```