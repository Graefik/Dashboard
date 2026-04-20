package httpadapter

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"back-dashboard/internal/adapters/http/dto"
	dmetrics "back-dashboard/internal/domain/metrics"
	"back-dashboard/internal/ports/inbound"
)

type MetricsHandler struct {
	uc inbound.MetricsUseCase
}

func NewMetricsHandler(uc inbound.MetricsUseCase) *MetricsHandler {
	return &MetricsHandler{uc: uc}
}

func (h *MetricsHandler) Overview(c *gin.Context) {
	ov, err := h.uc.GetOverview(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: "metrics indisponibles",
			Code:  "METRICS_UNAVAILABLE",
		})
		return
	}
	c.JSON(http.StatusOK, toMetricsResponse(ov))
}

func toMetricsResponse(ov *dmetrics.Overview) dto.MetricsOverviewResponse {
	routers := make([]dto.RouterRateDTO, 0, len(ov.Routers))
	for _, r := range ov.Routers {
		routers = append(routers, dto.RouterRateDTO{
			Name:  r.Name,
			RPS:   r.RPS,
			Codes: r.Codes,
		})
	}
	return dto.MetricsOverviewResponse{
		At:    ov.At.Unix(),
		Empty: ov.Empty,
		KPIs: dto.KPIsDTO{
			RPS:                 ov.KPIs.RPS,
			RPSDeltaPct:         ov.KPIs.RPSDeltaPct,
			LatencyP95Ms:        ov.KPIs.LatencyP95Ms,
			LatencyP95DeltaPct:  ov.KPIs.LatencyP95DeltaPct,
			Connections:         ov.KPIs.Connections,
			ConnectionsDeltaPct: ov.KPIs.ConnectionsDeltaPct,
			ErrorRatePct:        ov.KPIs.ErrorRatePct,
			ErrorRateDeltaPct:   ov.KPIs.ErrorRateDeltaPct,
		},
		Requests: dto.RequestsDTO{
			Timestamps: ov.Requests.Timestamps,
			Series:     ov.Requests.Series,
		},
		Latency: dto.LatencyDTO{
			Timestamps: ov.Latency.Timestamps,
			P50:        ov.Latency.P50,
			P95:        ov.Latency.P95,
			P99:        ov.Latency.P99,
		},
		StatusCodes: dto.StatusCodesDTO{
			C2xx:     ov.StatusCodes.C2xx,
			C3xx:     ov.StatusCodes.C3xx,
			C4xx:     ov.StatusCodes.C4xx,
			C5xx:     ov.StatusCodes.C5xx,
			Timeouts: ov.StatusCodes.Timeouts,
		},
		Routers: routers,
	}
}
