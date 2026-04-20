package metrics

import "time"

// Overview est la vue agrégée consommée par l'UI.
// Produite à partir de l'historique de Snapshots par le use case metrics.
type Overview struct {
	At          time.Time
	KPIs        KPIs
	Requests    TimeSeries
	Latency     LatencySeries
	StatusCodes StatusCodes
	Routers     []RouterRate // tous les routers avec du trafic, triés par RPS
	Empty       bool // true tant que le buffer n'a pas assez de données
}

type KPIs struct {
	RPS               float64
	RPSDeltaPct       float64
	LatencyP95Ms      float64
	LatencyP95DeltaPct float64
	Connections       int
	ConnectionsDeltaPct float64
	ErrorRatePct      float64
	ErrorRateDeltaPct float64
}

type TimeSeries struct {
	Timestamps []int64             // unix seconds
	Series     map[string][]float64 // "2xx","3xx","4xx","5xx"
}

type LatencySeries struct {
	Timestamps []int64
	P50        []float64 // ms
	P95        []float64
	P99        []float64
}

type StatusCodes struct {
	C2xx     uint64
	C3xx     uint64
	C4xx     uint64
	C5xx     uint64
	Timeouts uint64 // 5xx peut servir de proxy si pas de timeout dédié
}

type RouterRate struct {
	Name  string
	RPS   float64
	Codes map[string]uint64 // code HTTP → total cumulé (ex: "200":1234,"404":12)
}
