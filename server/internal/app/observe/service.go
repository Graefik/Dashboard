package observe

import (
	"context"

	"back-dashboard/internal/domain/traefik"
	"back-dashboard/internal/ports/outbound"
)

type Service struct {
	client outbound.TraefikClient
}

func NewService(client outbound.TraefikClient) *Service {
	return &Service{client: client}
}

func (s *Service) GetOverview(ctx context.Context) (*traefik.Overview, error) {
	return s.client.GetOverview(ctx)
}

func (s *Service) ListRouters(ctx context.Context) ([]traefik.Router, error) {
	return s.client.ListRouters(ctx)
}

func (s *Service) ListServices(ctx context.Context) ([]traefik.Service, error) {
	return s.client.ListServices(ctx)
}
