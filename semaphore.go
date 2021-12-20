package semaphore

import (
	"errors"
	"sync"
	"time"
)

type Semaphore chan *Permit

type Permit struct {
	once      sync.Once // the only chance to release
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
	// release only once
	p.once.Do(func() {
		// push a new Permit to channel
		p.semaphore <- &Permit{semaphore: p.semaphore}
	})
}

func (s Semaphore) Avails() int {
	return len(s)
}
