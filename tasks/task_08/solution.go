package main

import (
	"sync"
	"time"
)

type Clock interface{ Now() time.Time }
type Limiter struct {
	mu     sync.Mutex
	clock  Clock
	rate   float64
	burst  int
	tokens float64
	last   time.Time
}

func NewLimiter(clock Clock, ratePerSec float64, burst int) *Limiter {
	var last time.Time
	if clock != nil {
		last = clock.Now()
	}
	return &Limiter{
		mu:     sync.Mutex{},
		clock:  clock,
		rate:   ratePerSec,
		burst:  burst,
		tokens: float64(burst),
		last:   last,
	}
}
func (l *Limiter) AllowN(n int) bool {
	// сначала блокировка потоков
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.clock == nil || l.burst <= 0 || n <= 0 || n > l.burst {
		return false
	}

	now := l.clock.Now()
	if now.After(l.last) && l.rate > 0 {
	// пополнение токенов если время продвинулось вперед
	// и есть скорость пополнения не отрицательная
		elapsed := now.Sub(l.last).Seconds()
		l.tokens += elapsed * l.rate
		if l.tokens > float64(l.burst) {
			l.tokens = float64(l.burst)
		}
		l.last = now
	}

	if l.tokens >= float64(n) {
		l.tokens -= float64(n)
		return true
	}
	return false
}