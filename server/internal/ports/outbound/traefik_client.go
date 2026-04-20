package outbound

import (
	"context"

	"back-dashboard/internal/domain/traefik"
)

// TraefikClient parle à l'instance Traefik distante (API JSON).
// Les métriques Prometheus sont hors scope ici (autre adapter).
type TraefikClient interface {
	GetOverview(ctx context.Context) (*traefik.Overview, error)
	ListRouters(ctx context.Context) ([]traefik.Router, error)
	ListServices(ctx context.Context) ([]traefik.Service, error)
}
