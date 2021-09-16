package semaphore

import (
	"errors"
	"time"
)

type Semaphore chan *Permit

type Permit struct {
	releaseable bool
	semaphore   *Semaphore
}

func New(permits int) Semaphore {
	if permits < 1 {
		panic(errors.New("permits must be greater than 0"))
	}
	semaphore := make(chan *Permit, permits)
	for i := 0; i < permits; i++ {
		semaphore <- &Permit{semaphore: (*Semaphore)(&semaphore)}
	}
	return semaphore
}

func (s Semaphore) Acquire() *Permit {
	p := <-s
	p.releaseable = true
	return p
}

func (s Semaphore) TryAcquire() (*Permit, error) {
	select {
	case p := <-s:
		p.releaseable = true
		return p, nil
	default:
		return nil, errors.New("no permits can release")
	}
}

func (s Semaphore) TryAcquireTimeout(timeout time.Duration) (*Permit, error) {
	select {
	case p := <-s:
		p.releaseable = true
		return p, nil
	case <-time.After(timeout):
		return nil, errors.New("timeout")
	}
}

func (p *Permit) Release() {
	// check state
	if !p.releaseable {
		panic(errors.New("double release"))
	}

	p.releaseable = false
	*p.semaphore <- p
}
