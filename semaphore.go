package semaphore

import (
	"errors"
	"time"
)

type Semaphore chan *Permit

type Permit struct {
	canRelease bool
	semaphore  *Semaphore
}

func New(permits int) Semaphore {
	if permits < 1 {
		panic(errors.New("permits must be greater than 0"))
	}
	semaphore := make(chan *Permit, permits)
	for i := 0; i < permits; i++ {
		semaphore <- &Permit{semaphore: (*Semaphore)(&semaphore), canRelease: true}
	}
	return semaphore
}

func (s Semaphore) Acquire() *Permit {
	semaphore := <-s
	return semaphore
}

func (s Semaphore) TryAcquire() (*Permit, error) {
	select {
	case semaphore := <-s:
		return semaphore, nil
	default:
		return nil, errors.New("no permits can release")
	}
}

func (s Semaphore) TryAcquireTimeout(timeout time.Duration) (*Permit, error) {
	select {
	case semaphore := <-s:
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

	p.canRelease = false
	*p.semaphore <- p
	p.canRelease = true
}
