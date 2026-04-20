package inbound

import (
	"context"

	"back-dashboard/internal/domain/metrics"
)

type MetricsUseCase interface {
	GetOverview(ctx context.Context) (*metrics.Overview, error)
}
