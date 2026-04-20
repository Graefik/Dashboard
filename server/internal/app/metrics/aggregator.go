package metrics

import (
	"context"
	"math"
	"sort"
	"strings"
	"time"

	dmetrics "back-dashboard/internal/domain/metrics"
	"back-dashboard/internal/ports/outbound"
)

type Service struct {
	store outbound.MetricsStore
}

func NewService(store outbound.MetricsStore) *Service {
	return &Service{store: store}
}

const (
	deltaWindow    = 5 * time.Minute  // comparaison pour les deltas KPI
	timelineWindow = 60 * time.Minute // axe X du chart timeline
	timelineStep   = 1 * time.Minute
	latencyWindow  = 4 * time.Hour
	latencyStep    = 5 * time.Minute
)

func (s *Service) GetOverview(_ context.Context) (*dmetrics.Overview, error) {
	snaps := s.store.All()
	out := &dmetrics.Overview{At: time.Now().UTC()}
	if len(snaps) < 2 {
		out.Empty = true
		return out, nil
	}

	latest := snaps[len(snaps)-1]
	out.At = latest.At

	out.KPIs = computeKPIs(snaps)
	out.Requests = buildRequestsSeries(snaps, timelineWindow, timelineStep)
	out.Latency = buildLatencySeries(snaps, latencyWindow, latencyStep)
	out.StatusCodes = computeStatusCodes(latest)
	out.Routers = computeRouters(snaps)
	return out, nil
}

// ── KPIs ────────────────────────────────────────────────────────────────

func computeKPIs(snaps []*dmetrics.Snapshot) dmetrics.KPIs {
	latest := snaps[len(snaps)-1]
	first := snaps[0]

	// RPS "récent" : sur les N derniers snapshots (≈ 1 min)
	nowWin := snaps[maxIdx(len(snaps)-6, 0):]
	prevWin := windowEndingAt(snaps, latest.At.Add(-deltaWindow), 6)

	rpsNow := rps(nowWin)
	rpsPrev := rps(prevWin)

	errNow := errorRate(nowWin)
	errPrev := errorRate(prevWin)

	p95Now := percentileOverWindow(nowWin, 0.95) * 1000
	p95Prev := percentileOverWindow(prevWin, 0.95) * 1000

	connNow := latest.OpenConnections
	connPrev := snapshotAt(snaps, latest.At.Add(-deltaWindow))
	connPrevVal := connNow
	if connPrev != nil {
		connPrevVal = connPrev.OpenConnections
	}

	_ = first
	return dmetrics.KPIs{
		RPS:                 rpsNow,
		RPSDeltaPct:         pctDelta(rpsNow, rpsPrev),
		LatencyP95Ms:        p95Now,
		LatencyP95DeltaPct:  pctDelta(p95Now, p95Prev),
		Connections:         connNow,
		ConnectionsDeltaPct: pctDelta(float64(connNow), float64(connPrevVal)),
		ErrorRatePct:        errNow * 100,
		ErrorRateDeltaPct:   pctDelta(errNow, errPrev),
	}
}

func rps(snaps []*dmetrics.Snapshot) float64 {
	if len(snaps) < 2 {
		return 0
	}
	first, last := snaps[0], snaps[len(snaps)-1]
	dt := last.At.Sub(first.At).Seconds()
	if dt <= 0 {
		return 0
	}
	diff := totalRequests(last) - totalRequests(first)
	return float64(diff) / dt
}

func errorRate(snaps []*dmetrics.Snapshot) float64 {
	if len(snaps) < 2 {
		return 0
	}
	first, last := snaps[0], snaps[len(snaps)-1]
	totalDiff := totalRequests(last) - totalRequests(first)
	if totalDiff == 0 {
		return 0
	}
	errDiff := errorRequests(last) - errorRequests(first)
	return float64(errDiff) / float64(totalDiff)
}

func totalRequests(s *dmetrics.Snapshot) uint64 {
	var sum uint64
	for _, ep := range s.EntryPoints {
		for _, c := range ep.Requests {
			sum += c
		}
	}
	return sum
}

func errorRequests(s *dmetrics.Snapshot) uint64 {
	var sum uint64
	for _, ep := range s.EntryPoints {
		for code, c := range ep.Requests {
			if strings.HasPrefix(code, "5") {
				sum += c
			}
		}
	}
	return sum
}

// ── Timeline des requêtes par famille de code ───────────────────────────

func buildRequestsSeries(snaps []*dmetrics.Snapshot, window, step time.Duration) dmetrics.TimeSeries {
	series := map[string][]float64{
		"2xx": {},
		"3xx": {},
		"4xx": {},
		"5xx": {},
	}
	timestamps := []int64{}
	if len(snaps) < 2 {
		return dmetrics.TimeSeries{Timestamps: timestamps, Series: series}
	}

	end := snaps[len(snaps)-1].At
	start := end.Add(-window)

	for t := start; !t.After(end); t = t.Add(step) {
		a := snapshotAt(snaps, t.Add(-step))
		b := snapshotAt(snaps, t)
		if a == nil || b == nil || a == b {
			continue
		}
		dt := b.At.Sub(a.At).Seconds()
		if dt <= 0 {
			continue
		}
		timestamps = append(timestamps, t.Unix())
		for _, fam := range []string{"2", "3", "4", "5"} {
			d := familyDiff(a, b, fam)
			series[fam+"xx"] = append(series[fam+"xx"], float64(d)/dt)
		}
	}
	return dmetrics.TimeSeries{Timestamps: timestamps, Series: series}
}

func familyDiff(a, b *dmetrics.Snapshot, prefix string) uint64 {
	var sum uint64
	for _, ep := range b.EntryPoints {
		for code, c := range ep.Requests {
			if strings.HasPrefix(code, prefix) {
				sum += c
			}
		}
	}
	for _, ep := range a.EntryPoints {
		for code, c := range ep.Requests {
			if strings.HasPrefix(code, prefix) {
				if c > sum {
					sum = 0
					continue
				}
				sum -= c
			}
		}
	}
	return sum
}

// ── Latency percentiles ────────────────────────────────────────────────

func buildLatencySeries(snaps []*dmetrics.Snapshot, window, step time.Duration) dmetrics.LatencySeries {
	out := dmetrics.LatencySeries{}
	if len(snaps) < 2 {
		return out
	}
	end := snaps[len(snaps)-1].At
	start := end.Add(-window)
	for t := start; !t.After(end); t = t.Add(step) {
		a := snapshotAt(snaps, t.Add(-step))
		b := snapshotAt(snaps, t)
		if a == nil || b == nil || a == b {
			continue
		}
		out.Timestamps = append(out.Timestamps, t.Unix())
		out.P50 = append(out.P50, percentileDelta(a, b, 0.50)*1000)
		out.P95 = append(out.P95, percentileDelta(a, b, 0.95)*1000)
		out.P99 = append(out.P99, percentileDelta(a, b, 0.99)*1000)
	}
	return out
}

// percentileOverWindow : percentile sur l'ensemble d'un sous-ensemble de snaps
// (first → last), utilisé pour le KPI p95.
func percentileOverWindow(snaps []*dmetrics.Snapshot, p float64) float64 {
	if len(snaps) < 2 {
		return 0
	}
	return percentileDelta(snaps[0], snaps[len(snaps)-1], p)
}

// percentileDelta : calcule le p entre deux snapshots à partir des histogrammes
// delta (cumulatifs) sur les entrypoints.
func percentileDelta(a, b *dmetrics.Snapshot, p float64) float64 {
	// Agrège tous les buckets entrypoints
	bucketsA := aggregateBuckets(a.EntryPoints)
	bucketsB := aggregateBuckets(b.EntryPoints)

	les := sortedKeys(bucketsB)
	var total uint64
	for _, le := range les {
		total += bucketsB[le] - safeDiff(bucketsA[le], bucketsB[le])
	}
	if total == 0 {
		return 0
	}
	threshold := uint64(math.Ceil(float64(total) * p))
	var cumul uint64
	var prevLe float64
	var prevCumul uint64
	for _, le := range les {
		c := bucketsB[le] - safeDiff(bucketsA[le], bucketsB[le])
		cumul += c
		if cumul >= threshold {
			if math.IsInf(le, 1) {
				return prevLe
			}
			// interpolation linéaire entre prevLe et le
			if cumul-prevCumul == 0 {
				return le
			}
			frac := float64(threshold-prevCumul) / float64(cumul-prevCumul)
			return prevLe + frac*(le-prevLe)
		}
		prevLe = le
		prevCumul = cumul
	}
	return prevLe
}

// safeDiff : a bucket peut décroître si Traefik reset (reload) — on ignore.
func safeDiff(older, newer uint64) uint64 {
	if older > newer {
		return 0
	}
	return older
}

func aggregateBuckets(eps []dmetrics.EntryPointMetrics) map[float64]uint64 {
	out := map[float64]uint64{}
	for _, ep := range eps {
		for _, b := range ep.Latency.Buckets {
			out[b.Le] += b.Count
		}
	}
	return out
}

func sortedKeys(m map[float64]uint64) []float64 {
	keys := make([]float64, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Float64s(keys)
	return keys
}

// ── Status codes distribution ───────────────────────────────────────────

func computeStatusCodes(snap *dmetrics.Snapshot) dmetrics.StatusCodes {
	out := dmetrics.StatusCodes{}
	for _, ep := range snap.EntryPoints {
		for code, c := range ep.Requests {
			switch {
			case strings.HasPrefix(code, "2"):
				out.C2xx += c
			case strings.HasPrefix(code, "3"):
				out.C3xx += c
			case strings.HasPrefix(code, "4"):
				out.C4xx += c
			case strings.HasPrefix(code, "5"):
				out.C5xx += c
			}
		}
	}
	return out
}

// ── Routers (tous ceux avec du trafic, triés par RPS) ──────────────────

func computeRouters(snaps []*dmetrics.Snapshot) []dmetrics.RouterRate {
	if len(snaps) < 2 {
		return nil
	}
	first := snaps[0]
	last := snaps[len(snaps)-1]
	dt := last.At.Sub(first.At).Seconds()
	if dt <= 0 {
		return nil
	}

	startMap := map[string]uint64{}
	for _, r := range first.Routers {
		startMap[r.Name] = sumUint(r.Requests)
	}

	out := []dmetrics.RouterRate{}
	for _, r := range last.Routers {
		endTotal := sumUint(r.Requests)
		startTotal := startMap[r.Name]
		if endTotal < startTotal {
			continue
		}
		diff := endTotal - startTotal
		if diff == 0 {
			continue
		}
		out = append(out, dmetrics.RouterRate{
			Name:  r.Name,
			RPS:   float64(diff) / dt,
			Codes: cloneCodes(r.Requests),
		})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].RPS > out[j].RPS })
	return out
}

func sumUint(m map[string]uint64) uint64 {
	var s uint64
	for _, v := range m {
		s += v
	}
	return s
}

// cloneCodes évite de partager la map du snapshot (qui peut évoluer
// si un nouveau scrape arrive entre-temps).
func cloneCodes(src map[string]uint64) map[string]uint64 {
	if src == nil {
		return nil
	}
	out := make(map[string]uint64, len(src))
	for k, v := range src {
		out[k] = v
	}
	return out
}

// ── utils temporels ────────────────────────────────────────────────────

// snapshotAt : renvoie le snapshot dont l'At est le plus proche de target
// sans le dépasser. nil si aucun.
func snapshotAt(snaps []*dmetrics.Snapshot, target time.Time) *dmetrics.Snapshot {
	var chosen *dmetrics.Snapshot
	for _, s := range snaps {
		if s.At.After(target) {
			break
		}
		chosen = s
	}
	return chosen
}

func windowEndingAt(snaps []*dmetrics.Snapshot, end time.Time, minItems int) []*dmetrics.Snapshot {
	out := []*dmetrics.Snapshot{}
	for _, s := range snaps {
		if s.At.After(end) {
			break
		}
		out = append(out, s)
	}
	if len(out) < minItems {
		return out
	}
	return out[maxIdx(len(out)-minItems, 0):]
}

func pctDelta(now, prev float64) float64 {
	if prev == 0 {
		return 0
	}
	return (now - prev) / prev * 100
}

func maxIdx(a, b int) int {
	if a > b {
		return a
	}
	return b
}
