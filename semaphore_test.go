package semaphore

import (
	log "github.com/sirupsen/logrus"
	"sync"
	"testing"
	"time"
)

func Test_Sem(t *testing.T) {
	semaphore := New(1)
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
}

func Test_Sem_TryAcquire(t *testing.T) {
	semaphore := New(1)
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
}

func Test_timeout(t *testing.T) {
	semaphore := New(1)
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
}
