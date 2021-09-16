package gosem

import (
	log "github.com/sirupsen/logrus"
	"sync"
	"testing"
	"time"
)

func Test_Sem(t *testing.T) {
	pool := New(1)
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

	// double release, will panic
	//semaphore.Release()

	wg.Wait()
}

func Test_Sem_TryAcquire(t *testing.T) {
	pool := New(1)
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
}

func Test_timeout(t *testing.T) {
	pool := New(1)
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
}
