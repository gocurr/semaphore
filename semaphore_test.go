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
	fmt.Printf("begin: %d\n", semPool.Available())

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func(i int) {
			semaphore, err := semPool.TryAcquireTimeout(2 * time.Second)
			if err != nil {
				available := semPool.Available()
				fmt.Printf("semaphore is nil, available %d\n", available)
				wg.Done()
				return
			}
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
