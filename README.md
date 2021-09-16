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
semaphore := gosem.New(1)
permit := semaphore.Acquire()

var wg sync.WaitGroup
wg.Add(1)

go func() {
log.Info("begin of Acquire")
p := semaphore.Acquire()
log.Info("end of Acquire")

log.Infof("Acquired")
p.Release()
wg.Done()
}()

time.Sleep(5 * time.Second)
permit.Release()

// double release, will panic
//semaphore.Release()

wg.Wait()
```

- TryAcquire:

```go
semaphore := gosem.New(1)
permit := semaphore.Acquire()

var wg sync.WaitGroup
wg.Add(1)

go func() {
log.Info("begin of TryAcquire")
p, err := semaphore.TryAcquire()
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
semaphore := gosem.New(1)
permit := semaphore.Acquire()

var wg sync.WaitGroup
wg.Add(1)

go func() {
log.Info("begin of TryAcquireTimeout")
p, err := semaphore.TryAcquireTimeout(10 * time.Second)
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