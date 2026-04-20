package memstore

import (
	"sync"

	"back-dashboard/internal/domain/metrics"
)

// RingBuffer : store circulaire thread-safe pour snapshots.
// Une fois plein, chaque Push écrase le plus ancien.
type RingBuffer struct {
	mu   sync.RWMutex
	buf  []*metrics.Snapshot
	head int // index du prochain slot à écrire
	size int // nombre d'éléments actuellement stockés
	cap  int
}

func NewRingBuffer(capacity int) *RingBuffer {
	if capacity <= 0 {
		capacity = 360 // 1h @ 10s
	}
	return &RingBuffer{
		buf: make([]*metrics.Snapshot, capacity),
		cap: capacity,
	}
}

func (r *RingBuffer) Push(s *metrics.Snapshot) {
	if s == nil {
		return
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	r.buf[r.head] = s
	r.head = (r.head + 1) % r.cap
	if r.size < r.cap {
		r.size++
	}
}

func (r *RingBuffer) All() []*metrics.Snapshot {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if r.size == 0 {
		return nil
	}
	out := make([]*metrics.Snapshot, 0, r.size)
	start := (r.head - r.size + r.cap) % r.cap
	for i := 0; i < r.size; i++ {
		out = append(out, r.buf[(start+i)%r.cap])
	}
	return out
}

func (r *RingBuffer) Latest() *metrics.Snapshot {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if r.size == 0 {
		return nil
	}
	idx := (r.head - 1 + r.cap) % r.cap
	return r.buf[idx]
}

func (r *RingBuffer) Oldest() *metrics.Snapshot {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if r.size == 0 {
		return nil
	}
	idx := (r.head - r.size + r.cap) % r.cap
	return r.buf[idx]
}

func (r *RingBuffer) Len() int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.size
}
