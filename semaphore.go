package semaphore

import "errors"

type SemPool chan *Semaphore

type Semaphore struct {
	pool *SemPool
}

func New(size int) SemPool {
	if size <= 0 {
		panic(errors.New("negative cannot be a size"))
	}
	pool := make(chan *Semaphore, size)
	for i := 0; i < size; i++ {
		pool <- &Semaphore{pool: (*SemPool)(&pool)}
	}
	return pool
}

func (s SemPool) Acquire() *Semaphore {
	return <-s
}

func (s *Semaphore) Release() {
	*s.pool <- s
}
