# semaphore

To download, run:

```bash
go get -u github.com/gocurr/semaphore
```

Import it in your program as:

```go
import "github.com/gocurr/semaphore"
```

It requires Go 1.11 or later due to usage of Go Modules.

- Usage example:

```go
rand.Seed(time.Now().Unix())

semPool := New(runtime.NumCPU())
var wg sync.WaitGroup

for i := 0; i < 100; i++ {
    wg.Add(1)
    
    go func(i int) {
        semaphore := semPool.Acquire()
        // mock slow operation
        n := rand.Intn(5)
        time.Sleep(time.Duration(n) * time.Second)
        fmt.Printf("%d released\n", i)
        semaphore.Release()
        wg.Done()
    }(i)
}

wg.Wait()
```