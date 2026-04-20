package outbound

import (
	"context"

	"back-dashboard/internal/domain/metrics"
)

// MetricsScraper fait UN scrape de l'endpoint Prometheus Traefik
// et retourne un snapshot domain-pur.
type MetricsScraper interface {
	ScrapeOnce(ctx context.Context) (*metrics.Snapshot, error)
}

// MetricsStore conserve un historique glissant de snapshots.
// Implémentation typique : ring buffer en mémoire.
type MetricsStore interface {
	Push(s *metrics.Snapshot)
	// All renvoie les snapshots ordonnés du plus ancien au plus récent.
	All() []*metrics.Snapshot
	Latest() *metrics.Snapshot
	Oldest() *metrics.Snapshot
	Len() int
}
