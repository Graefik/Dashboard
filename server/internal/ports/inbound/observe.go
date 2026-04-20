package inbound

import (
	"context"

	"back-dashboard/internal/domain/traefik"
)

// ObserveUseCase expose les données Traefik aux adapters driving.
type ObserveUseCase interface {
	GetOverview(ctx context.Context) (*traefik.Overview, error)
	ListRouters(ctx context.Context) ([]traefik.Router, error)
	ListServices(ctx context.Context) ([]traefik.Service, error)
}
