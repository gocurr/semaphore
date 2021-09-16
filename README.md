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

- Acquire:

```go
sem := semaphore.New(1)
permit := sem.Acquire()

var wg sync.WaitGroup
wg.Add(1)

go func() {
    log.Info("begin of Acquire")
    p := sem.Acquire()
    log.Info("end of Acquire")
    
    log.Infof("Acquired")
    p.Release()
    wg.Done()
}()

time.Sleep(5 * time.Second)
permit.Release()

// double release, will panic
//sem.Release()

wg.Wait()
```

- TryAcquire:

```go
sem := semaphore.New(1)
permit := sem.Acquire()

var wg sync.WaitGroup
wg.Add(1)

go func() {
    log.Info("begin of TryAcquire")
    p, err := sem.TryAcquire()
    log.Info("end of TryAcquire")
    if err != nil {
        log.Errorf("%v", err)
        wg.Done()
        return
    }

    log.Infof("Acquired")
    p.Release()
    wg.Done()
}()

time.Sleep(1 * time.Second)
permit.Release()

wg.Wait()
```

- TryAcquireTimeout:

```go
sem := semaphore.New(1)
permit := sem.Acquire()

var wg sync.WaitGroup
wg.Add(1)

go func() {
    log.Info("begin of TryAcquireTimeout")
    p, err := sem.TryAcquireTimeout(10 * time.Second)
    log.Info("end of TryAcquireTimeout")
    if err != nil {
        log.Errorf("%v", err)
        wg.Done()
        return
    }
    
    log.Infof("Acquired")
    p.Release()
    wg.Done()
}()

time.Sleep(15 * time.Second)
permit.Release()

wg.Wait()
```