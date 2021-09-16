package semaphore

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func Test_Sem(t *testing.T) {
	rand.Seed(time.Now().Unix())
	semPool := New(2)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
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
