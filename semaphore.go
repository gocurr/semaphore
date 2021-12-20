package semaphore

import (
	"errors"
	"time"
)

type Semaphore chan *Permit

type Permit struct {
	semaphore Semaphore
}

func New(permits int) Semaphore {
	if permits < 1 {
		panic("permits must be greater than 0")
	}
	semaphore := make(chan *Permit, permits)
	for i := 0; i < permits; i++ {
		semaphore <- &Permit{semaphore: semaphore}
	}
	return semaphore
}

func (s Semaphore) Acquire() *Permit {
	return <-s
}

func (s Semaphore) TryAcquire() (*Permit, error) {
	select {
	case p := <-s:
		return p, nil
	default:
		return nil, errors.New("no permits can be released")
	}
}

func (s Semaphore) TryAcquireTimeout(timeout time.Duration) (*Permit, error) {
	select {
	case p := <-s:
		return p, nil
	case <-time.After(timeout):
		return nil, errors.New("try to acquire timeout")
	}
}

func (p *Permit) Release() {
	if p.semaphore == nil {
		panic("double release")
	}
	// push a new Permit to channel
	p.semaphore <- &Permit{semaphore: p.semaphore}
	// evict the old one
	p.semaphore = nil
}

func (s Semaphore) Avails() int {
	return len(s)
}
