package gosem

import (
	"errors"
	"time"
)

type SemPool chan *Semaphore

type Semaphore struct {
	pool *SemPool
}

func New(permits int) SemPool {
	if permits <= 0 {
		panic(errors.New("negative or 0 cannot be a permits"))
	}
	pool := make(chan *Semaphore, permits)
	for i := 0; i < permits; i++ {
		pool <- &Semaphore{pool: (*SemPool)(&pool)}
	}
	return pool
}

func (s SemPool) Acquire() *Semaphore {
	return <-s
}

func (s SemPool) TryAcquire() (*Semaphore, error) {
	select {
	case ret := <-s:
		return ret, nil
	default:
		return nil, errors.New("no permits available")
	}
}

func (s SemPool) TryAcquireTimeout(timeout time.Duration) (*Semaphore, error) {
	for {
		select {
		case ret := <-s:
			return ret, nil
		case <-time.After(timeout):
			return nil, errors.New("timeout")
		}
	}
}

func (s SemPool) Available() int {
	return len(s)
}

func (s *Semaphore) Release() {
	*s.pool <- s
}
