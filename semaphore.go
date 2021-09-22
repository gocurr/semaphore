package semaphore

import (
	"errors"
	"time"
)

var (
	lessThanOneErr   = errors.New("permits must be greater than 0")
	noPermitsErr     = errors.New("no permits can be released")
	timeoutErr       = errors.New("try to acquire timeout")
	doubleReleaseErr = errors.New("double release")
)

type Semaphore chan *Permit

type Permit struct {
	releasable bool
	semaphore  Semaphore
}

func New(permits int) Semaphore {
	if permits < 1 {
		panic(lessThanOneErr)
	}
	semaphore := make(chan *Permit, permits)
	for i := 0; i < permits; i++ {
		semaphore <- &Permit{semaphore: semaphore}
	}
	return semaphore
}

func (s Semaphore) Acquire() *Permit {
	p := <-s
	p.releasable = true
	return p
}

func (s Semaphore) TryAcquire() (*Permit, error) {
	select {
	case p := <-s:
		p.releasable = true
		return p, nil
	default:
		return nil, noPermitsErr
	}
}

func (s Semaphore) TryAcquireTimeout(timeout time.Duration) (*Permit, error) {
	select {
	case p := <-s:
		p.releasable = true
		return p, nil
	case <-time.After(timeout):
		return nil, timeoutErr
	}
}

func (p *Permit) Release() {
	if !p.releasable {
		panic(doubleReleaseErr)
	}

	p.releasable = false
	p.semaphore <- p
}

func (s Semaphore) Avails() int {
	return len(s)
}
