package traefikclient

import "back-dashboard/internal/domain/traefik"

func toDomainOverview(d *overviewDTO, version string) *traefik.Overview {
	return &traefik.Overview{
		HTTP:    toDomainSection(&d.HTTP),
		TCP:     toDomainSection(&d.TCP),
		UDP:     toDomainSection(&d.UDP),
		Features: traefik.Features{
			Tracing:   d.Features.Tracing,
			Metrics:   d.Features.Metrics,
			AccessLog: d.Features.AccessLog,
			Hub:       d.Features.Hub,
		},
		Version: version,
	}
}

func toDomainSection(s *sectionDTO) traefik.Section {
	return traefik.Section{
		Routers:     toDomainCounts(&s.Routers),
		Services:    toDomainCounts(&s.Services),
		Middlewares: toDomainCounts(&s.Middlewares),
	}
}

func toDomainCounts(c *countsDTO) traefik.Counts {
	return traefik.Counts{
		Total:    c.Total,
		Warnings: c.Warnings,
		Errors:   c.Errors,
	}
}

func toDomainRouter(r *routerDTO) traefik.Router {
	out := traefik.Router{
		Name:        r.Name,
		Rule:        r.Rule,
		EntryPoints: r.EntryPoints,
		Service:     r.Service,
		Middlewares: r.Middlewares,
		Status:      r.Status,
		Using:       r.Using,
		Provider:    r.Provider,
	}
	if r.TLS != nil {
		tls := traefik.RouterTLS{CertResolver: r.TLS.CertResolver}
		for _, d := range r.TLS.Domains {
			tls.Domains = append(tls.Domains, traefik.CertDomain{Main: d.Main, SANs: d.SANs})
		}
		out.TLS = &tls
	}
	return out
}

func toDomainService(s *serviceDTO) traefik.Service {
	return traefik.Service{
		Name:         s.Name,
		Type:         s.Type,
		Provider:     s.Provider,
		Status:       s.Status,
		UsedBy:       s.UsedBy,
		ServerStatus: s.ServerStatus,
	}
}
