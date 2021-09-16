package gosem

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"testing"
	"time"
)

func Test_Sem(t *testing.T) {
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
}

func Test_Sem_TryAcquire(t *testing.T) {
	rand.Seed(time.Now().Unix())

	semPool := New(runtime.NumCPU())

	var wg sync.WaitGroup

	for i := 0; i < 80; i++ {
		wg.Add(1)

		go func(i int) {
			semaphore, err := semPool.TryAcquireTimeout(9 * time.Second)
			if err != nil {
				fmt.Printf("semaphore is nil\n")
				wg.Done()
				return
			}
			// mock slow operation
			n := rand.Intn(3)
			time.Sleep(time.Duration(n) * time.Second)
			fmt.Printf("%d released\n", i)
			semaphore.Release()
			wg.Done()
		}(i)
	}

	wg.Wait()
}

func Test_timeout(t *testing.T) {
	pool := New(2)
	_ = pool.Acquire()
	_ = pool.Acquire()
	_, err := pool.TryAcquireTimeout(2 * time.Second)
	if err != nil {
		t.Errorf("%v", err)
	}

	time.Sleep(10 * time.Second)
}
