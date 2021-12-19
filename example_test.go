package semaphore_test

import (
	"fmt"
	"github.com/gocurr/semaphore"
	"sync"
	"time"
)

func ExampleSemaphore_Acquire() {
	s := semaphore.New(1)
	permit := s.Acquire()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		p := s.Acquire()
		fmt.Println("goroutine acquired")

		time.Sleep(1 * time.Second)
		fmt.Println("goroutine is releasing")
		p.Release()
		wg.Done()
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("main is releasing")
	permit.Release()

	s.Acquire().Release()
	fmt.Println("main acquired")
	wg.Wait()

	// Output: main is releasing
	// goroutine acquired
	// goroutine is releasing
	// main acquired
}

func ExampleSemaphore_TryAcquire() {
	s := semaphore.New(1)
	permit := s.Acquire()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		p, err := s.TryAcquire()
		if err != nil {
			fmt.Println("cannot acquire")
			wg.Done()
			return
		}

		fmt.Println("goroutine is releasing")
		p.Release()
		wg.Done()
	}()

	time.Sleep(1 * time.Second)
	permit.Release()

	wg.Wait()

	// Output: cannot acquire
}

func ExampleSemaphore_TryAcquireTimeout() {
	s := semaphore.New(1)
	permit := s.Acquire()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		p, err := s.TryAcquireTimeout(1 * time.Second)
		if err != nil {
			fmt.Println("goroutine cannot acquire")
			wg.Done()
			return
		}

		fmt.Println("goroutine Acquired")
		p.Release()
		wg.Done()
	}()

	time.Sleep(2 * time.Second)
	permit.Release()

	wg.Wait()

	// Output: goroutine cannot acquire

}
