package prometheusclient

import (
	"math"
	"sort"

	dto "github.com/prometheus/client_model/go"

	"back-dashboard/internal/domain/metrics"
)

// parseSnapshot remplit le snapshot à partir des metric families Prometheus
// exposées par Traefik. Traefik v3 :
//   - traefik_entrypoint_requests_total{code, entrypoint, method, protocol}
//   - traefik_entrypoint_request_duration_seconds_bucket{code, entrypoint, method, protocol, le}
//   - traefik_router_requests_total{code, method, protocol, router, service}
//   - traefik_router_request_duration_seconds_bucket{...}
//   - traefik_service_requests_total{code, method, protocol, service}
//   - traefik_service_request_duration_seconds_bucket{...}
//   - traefik_open_connections{entrypoint, method, protocol}
func parseSnapshot(snap *metrics.Snapshot, families map[string]*dto.MetricFamily) {
	epMap := map[string]*metrics.EntryPointMetrics{}
	rtMap := map[string]*metrics.RouterMetrics{}
	svcMap := map[string]*metrics.ServiceMetrics{}

	// ── Requests counters ───────────────────────────────────────
	if f := families["traefik_entrypoint_requests_total"]; f != nil {
		for _, m := range f.Metric {
			ep := labelValue(m, "entrypoint")
			code := labelValue(m, "code")
			if ep == "" || code == "" || m.Counter == nil {
				continue
			}
			e := ensureEP(epMap, ep)
			e.Requests[code] += uint64(m.Counter.GetValue())
		}
	}

	if f := families["traefik_router_requests_total"]; f != nil {
		for _, m := range f.Metric {
			rt := labelValue(m, "router")
			code := labelValue(m, "code")
			if rt == "" || code == "" || m.Counter == nil {
				continue
			}
			r := ensureRouter(rtMap, rt, labelValue(m, "service"))
			r.Requests[code] += uint64(m.Counter.GetValue())
		}
	}

	if f := families["traefik_service_requests_total"]; f != nil {
		for _, m := range f.Metric {
			svc := labelValue(m, "service")
			code := labelValue(m, "code")
			if svc == "" || code == "" || m.Counter == nil {
				continue
			}
			s := ensureService(svcMap, svc)
			s.Requests[code] += uint64(m.Counter.GetValue())
		}
	}

	// ── Duration histograms (on somme par entrypoint/router/service) ──
	if f := families["traefik_entrypoint_request_duration_seconds"]; f != nil {
		for _, m := range f.Metric {
			ep := labelValue(m, "entrypoint")
			if ep == "" || m.Histogram == nil {
				continue
			}
			addHistogram(&ensureEP(epMap, ep).Latency, m.Histogram)
		}
	}

	if f := families["traefik_router_request_duration_seconds"]; f != nil {
		for _, m := range f.Metric {
			rt := labelValue(m, "router")
			if rt == "" || m.Histogram == nil {
				continue
			}
			addHistogram(&ensureRouter(rtMap, rt, labelValue(m, "service")).Latency, m.Histogram)
		}
	}

	if f := families["traefik_service_request_duration_seconds"]; f != nil {
		for _, m := range f.Metric {
			svc := labelValue(m, "service")
			if svc == "" || m.Histogram == nil {
				continue
			}
			addHistogram(&ensureService(svcMap, svc).Latency, m.Histogram)
		}
	}

	// ── Open connections (gauge) ───────────────────────────────
	if f := families["traefik_open_connections"]; f != nil {
		total := 0
		for _, m := range f.Metric {
			if m.Gauge != nil {
				total += int(m.Gauge.GetValue())
			}
		}
		snap.OpenConnections = total
	}

	snap.EntryPoints = toEPSlice(epMap)
	snap.Routers = toRouterSlice(rtMap)
	snap.Services = toServiceSlice(svcMap)
}

// ── helpers ────────────────────────────────────────────────────

func labelValue(m *dto.Metric, name string) string {
	for _, l := range m.Label {
		if l.GetName() == name {
			return l.GetValue()
		}
	}
	return ""
}

func ensureEP(m map[string]*metrics.EntryPointMetrics, name string) *metrics.EntryPointMetrics {
	e, ok := m[name]
	if !ok {
		e = &metrics.EntryPointMetrics{Name: name, Requests: map[string]uint64{}}
		m[name] = e
	}
	return e
}

func ensureRouter(m map[string]*metrics.RouterMetrics, name, service string) *metrics.RouterMetrics {
	r, ok := m[name]
	if !ok {
		r = &metrics.RouterMetrics{Name: name, Service: service, Requests: map[string]uint64{}}
		m[name] = r
	}
	return r
}

func ensureService(m map[string]*metrics.ServiceMetrics, name string) *metrics.ServiceMetrics {
	s, ok := m[name]
	if !ok {
		s = &metrics.ServiceMetrics{Name: name, Requests: map[string]uint64{}}
		m[name] = s
	}
	return s
}

// addHistogram : agrège un dto.Histogram dans le domain. On additionne bucket par bucket.
func addHistogram(dst *metrics.Histogram, src *dto.Histogram) {
	if src == nil {
		return
	}
	dst.Sum += src.GetSampleSum()
	dst.Count += src.GetSampleCount()

	existing := map[float64]uint64{}
	for _, b := range dst.Buckets {
		existing[b.Le] = b.Count
	}
	for _, b := range src.Bucket {
		le := b.GetUpperBound()
		if math.IsInf(le, 1) {
			le = math.Inf(1)
		}
		existing[le] += b.GetCumulativeCount()
	}
	out := make([]metrics.Bucket, 0, len(existing))
	for le, c := range existing {
		out = append(out, metrics.Bucket{Le: le, Count: c})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Le < out[j].Le })
	dst.Buckets = out
}

func toEPSlice(m map[string]*metrics.EntryPointMetrics) []metrics.EntryPointMetrics {
	out := make([]metrics.EntryPointMetrics, 0, len(m))
	for _, v := range m {
		out = append(out, *v)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out
}

func toRouterSlice(m map[string]*metrics.RouterMetrics) []metrics.RouterMetrics {
	out := make([]metrics.RouterMetrics, 0, len(m))
	for _, v := range m {
		out = append(out, *v)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out
}

func toServiceSlice(m map[string]*metrics.ServiceMetrics) []metrics.ServiceMetrics {
	out := make([]metrics.ServiceMetrics, 0, len(m))
	for _, v := range m {
		out = append(out, *v)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out
}
