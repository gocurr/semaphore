package gosem

import (
	"errors"
	"time"
)

type SemPool chan *Semaphore

type Semaphore struct {
	enabled bool
	pool    *SemPool
}

func New(permits int) SemPool {
	if permits <= 0 {
		panic(errors.New("negative or 0 cannot be a permits"))
	}
	pool := make(chan *Semaphore, permits)
	for i := 0; i < permits; i++ {
		pool <- &Semaphore{pool: (*SemPool)(&pool), enabled: true}
	}
	return pool
}

func (pool SemPool) Acquire() *Semaphore {
	semaphore := <-pool
	semaphore.enable()
	return semaphore
}

func (pool SemPool) TryAcquire() (*Semaphore, error) {
	select {
	case semaphore := <-pool:
		semaphore.enable()
		return semaphore, nil
	default:
		return nil, errors.New("no permits enabled")
	}
}

func (pool SemPool) TryAcquireTimeout(timeout time.Duration) (*Semaphore, error) {
	select {
	case semaphore := <-pool:
		semaphore.enable()
		return semaphore, nil
	case <-time.After(timeout):
		return nil, errors.New("timeout")
	}
}

func (s *Semaphore) Release() {
	// check state
	if !s.enabled {
		panic(errors.New("double release"))
	}

	s.disable()
	*s.pool <- s
}

func (s *Semaphore) enable() {
	s.enabled = true
}

func (s *Semaphore) disable() {
	s.enabled = false
}
