package metrics

import "time"

// Snapshot : image à un instant T de l'état des métriques Traefik.
// Tous les compteurs sont CUMULATIFS (monotones croissants) ;
// les rates se calculent en comparant deux snapshots.
type Snapshot struct {
	At              time.Time
	EntryPoints     []EntryPointMetrics
	Routers         []RouterMetrics
	Services        []ServiceMetrics
	OpenConnections int // somme sur tous entrypoints/méthodes
}

type EntryPointMetrics struct {
	Name     string
	Requests map[string]uint64 // code HTTP ("200", "404", "500"…) → total cumulé
	Latency  Histogram
}

type RouterMetrics struct {
	Name     string
	Service  string
	Requests map[string]uint64
	Latency  Histogram
}

type ServiceMetrics struct {
	Name     string
	Requests map[string]uint64
	Latency  Histogram
}

// Histogram : données brutes d'un histogramme Prometheus.
type Histogram struct {
	Buckets []Bucket // triés par Le croissant
	Sum     float64  // somme cumulative des durées
	Count   uint64   // nombre cumulatif d'observations
}

type Bucket struct {
	Le    float64 // borne sup (seconds), +Inf = math.Inf(1)
	Count uint64  // count cumulatif jusqu'à cette borne
}
