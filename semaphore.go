package gosem

import (
	"errors"
	"time"
)

type Semaphore chan *Permit

type Permit struct {
	canRelease bool
	pool       *Semaphore
}

func New(permits int) Semaphore {
	if permits < 1 {
		panic(errors.New("permits must be greater than 0"))
	}
	pool := make(chan *Permit, permits)
	for i := 0; i < permits; i++ {
		pool <- &Permit{pool: (*Semaphore)(&pool)}
	}
	return pool
}

func (s Semaphore) Acquire() *Permit {
	semaphore := <-s
	semaphore.releaseAble()
	return semaphore
}

func (s Semaphore) TryAcquire() (*Permit, error) {
	select {
	case semaphore := <-s:
		semaphore.releaseAble()
		return semaphore, nil
	default:
		return nil, errors.New("no permits can release")
	}
}

func (s Semaphore) TryAcquireTimeout(timeout time.Duration) (*Permit, error) {
	select {
	case semaphore := <-s:
		semaphore.releaseAble()
		return semaphore, nil
	case <-time.After(timeout):
		return nil, errors.New("timeout")
	}
}

func (p *Permit) Release() {
	// check state
	if !p.canRelease {
		panic(errors.New("double release"))
	}

	p.releaseDisable()
	*p.pool <- p
}

func (p *Permit) releaseAble() {
	p.canRelease = true
}

func (p *Permit) releaseDisable() {
	p.canRelease = false
}
