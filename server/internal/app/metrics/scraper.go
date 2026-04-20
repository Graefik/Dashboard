package metrics

import (
	"context"
	"log"
	"time"

	"back-dashboard/internal/ports/outbound"
)

// Scraper lance une goroutine qui pousse un snapshot à intervalle régulier.
type Scraper struct {
	src      outbound.MetricsScraper
	store    outbound.MetricsStore
	interval time.Duration
}

func NewScraper(src outbound.MetricsScraper, store outbound.MetricsStore, interval time.Duration) *Scraper {
	if interval <= 0 {
		interval = 10 * time.Second
	}
	return &Scraper{src: src, store: store, interval: interval}
}

// Start lance le scrape en arrière-plan jusqu'à ctx.Done().
// Un scrape immédiat est exécuté avant de commencer le ticker.
func (s *Scraper) Start(ctx context.Context) {
	go func() {
		s.tick(ctx)
		t := time.NewTicker(s.interval)
		defer t.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-t.C:
				s.tick(ctx)
			}
		}
	}()
}

func (s *Scraper) tick(ctx context.Context) {
	c, cancel := context.WithTimeout(ctx, s.interval)
	defer cancel()
	snap, err := s.src.ScrapeOnce(c)
	if err != nil {
		log.Printf("metrics scrape: %v", err)
		return
	}
	s.store.Push(snap)
}
