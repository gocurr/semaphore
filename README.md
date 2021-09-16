# gosem

To download, run:

```bash
go get -u github.com/gocurr/gosem
```

Import it in your program as:

```go
import "github.com/gocurr/gosem"
```

It requires Go 1.11 or later due to usage of Go Modules.

- Acquire:

```go
pool := gosem.New(1)
semaphore := pool.Acquire()

var wg sync.WaitGroup
wg.Add(1)

go func() {
    log.Info("begin of TryAcquireTimeout")
    sem := pool.Acquire()
    log.Info("end of TryAcquireTimeout")
    
    log.Infof("Acquired")
    sem.Release()
    wg.Done()
}()

time.Sleep(5 * time.Second)
semaphore.Release()

wg.Wait()
```

- TryAcquire:

```go
pool := gosem.gosem.New(1)
semaphore := pool.Acquire()

var wg sync.WaitGroup
wg.Add(1)

go func() {
    log.Info("begin of TryAcquireTimeout")
    sem, err := pool.TryAcquire()
    log.Info("end of TryAcquireTimeout")
    if err != nil {
        log.Errorf("%v", err)
        wg.Done()
        return
    }
    
    log.Infof("Acquired")
    sem.Release()
    wg.Done()
}()

time.Sleep(1 * time.Second)
semaphore.Release()

wg.Wait()
```

- TryAcquireTimeout:

```go
pool := gosem.gosem.New(1)
semaphore := pool.Acquire()

var wg sync.WaitGroup
wg.Add(1)

go func() {
    log.Info("begin of TryAcquireTimeout")
    sem, err := pool.TryAcquireTimeout(10 * time.Second)
    log.Info("end of TryAcquireTimeout")
    if err != nil {
        log.Errorf("%v", err)
        wg.Done()
        return
    }

    log.Infof("Acquired")
    sem.Release()
    wg.Done()
}()

time.Sleep(15 * time.Second)
semaphore.Release()

wg.Wait()
```