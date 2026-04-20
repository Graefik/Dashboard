package dto

type MetricsOverviewResponse struct {
	At          int64            `json:"at"`
	Empty       bool             `json:"empty"`
	KPIs        KPIsDTO          `json:"kpis"`
	Requests    RequestsDTO      `json:"requests"`
	Latency     LatencyDTO       `json:"latency"`
	StatusCodes StatusCodesDTO   `json:"statusCodes"`
	Routers     []RouterRateDTO  `json:"routers"`
}

type KPIsDTO struct {
	RPS                 float64 `json:"rps"`
	RPSDeltaPct         float64 `json:"rpsDeltaPct"`
	LatencyP95Ms        float64 `json:"latencyP95Ms"`
	LatencyP95DeltaPct  float64 `json:"latencyP95DeltaPct"`
	Connections         int     `json:"connections"`
	ConnectionsDeltaPct float64 `json:"connectionsDeltaPct"`
	ErrorRatePct        float64 `json:"errorRatePct"`
	ErrorRateDeltaPct   float64 `json:"errorRateDeltaPct"`
}

type RequestsDTO struct {
	Timestamps []int64              `json:"timestamps"`
	Series     map[string][]float64 `json:"series"`
}

type LatencyDTO struct {
	Timestamps []int64   `json:"timestamps"`
	P50        []float64 `json:"p50"`
	P95        []float64 `json:"p95"`
	P99        []float64 `json:"p99"`
}

type StatusCodesDTO struct {
	C2xx     uint64 `json:"c2xx"`
	C3xx     uint64 `json:"c3xx"`
	C4xx     uint64 `json:"c4xx"`
	C5xx     uint64 `json:"c5xx"`
	Timeouts uint64 `json:"timeouts"`
}

type RouterRateDTO struct {
	Name  string            `json:"name"`
	RPS   float64           `json:"rps"`
	Codes map[string]uint64 `json:"codes,omitempty"`
}
