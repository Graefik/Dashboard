package httpadapter

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"back-dashboard/internal/adapters/http/dto"
	"back-dashboard/internal/domain/traefik"
	"back-dashboard/internal/ports/inbound"
)

type ObserveHandler struct {
	observe inbound.ObserveUseCase
}

func NewObserveHandler(observe inbound.ObserveUseCase) *ObserveHandler {
	return &ObserveHandler{observe: observe}
}

func (h *ObserveHandler) Overview(c *gin.Context) {
	ov, err := h.observe.GetOverview(c.Request.Context())
	if err != nil {
		respondTraefikError(c, err)
		return
	}
	c.JSON(http.StatusOK, toOverviewResponse(ov))
}

func (h *ObserveHandler) Routers(c *gin.Context) {
	routers, err := h.observe.ListRouters(c.Request.Context())
	if err != nil {
		respondTraefikError(c, err)
		return
	}
	out := make([]dto.RouterDTO, 0, len(routers))
	for _, r := range routers {
		out = append(out, toRouterDTO(&r))
	}
	c.JSON(http.StatusOK, out)
}

func (h *ObserveHandler) Services(c *gin.Context) {
	services, err := h.observe.ListServices(c.Request.Context())
	if err != nil {
		respondTraefikError(c, err)
		return
	}
	out := make([]dto.ServiceDTO, 0, len(services))
	for _, s := range services {
		out = append(out, toServiceDTO(&s))
	}
	c.JSON(http.StatusOK, out)
}

func respondTraefikError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, traefik.ErrUnauthorized):
		c.JSON(http.StatusBadGateway, dto.ErrorResponse{
			Error: "identifiants Traefik invalides",
			Code:  "TRAEFIK_UNAUTHORIZED",
		})
	case errors.Is(err, traefik.ErrUnavailable):
		c.JSON(http.StatusBadGateway, dto.ErrorResponse{
			Error: "instance Traefik injoignable",
			Code:  "TRAEFIK_UNAVAILABLE",
		})
	case errors.Is(err, traefik.ErrNotFound):
		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Error: "ressource Traefik introuvable",
			Code:  "TRAEFIK_NOT_FOUND",
		})
	default:
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: "erreur interne",
			Code:  "INTERNAL",
		})
	}
}

func toOverviewResponse(ov *traefik.Overview) dto.OverviewResponse {
	return dto.OverviewResponse{
		HTTP:    toSectionDTO(&ov.HTTP),
		TCP:     toSectionDTO(&ov.TCP),
		UDP:     toSectionDTO(&ov.UDP),
		Features: dto.FeaturesDTO{
			Tracing:   ov.Features.Tracing,
			Metrics:   ov.Features.Metrics,
			AccessLog: ov.Features.AccessLog,
			Hub:       ov.Features.Hub,
		},
		Version: ov.Version,
	}
}

func toSectionDTO(s *traefik.Section) dto.SectionDTO {
	return dto.SectionDTO{
		Routers:     dto.CountsDTO(s.Routers),
		Services:    dto.CountsDTO(s.Services),
		Middlewares: dto.CountsDTO(s.Middlewares),
	}
}

func toRouterDTO(r *traefik.Router) dto.RouterDTO {
	out := dto.RouterDTO{
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
		tls := dto.RouterTLSDTO{CertResolver: r.TLS.CertResolver}
		for _, d := range r.TLS.Domains {
			tls.Domains = append(tls.Domains, dto.CertDomainDTO{Main: d.Main, SANs: d.SANs})
		}
		out.TLS = &tls
	}
	return out
}

func toServiceDTO(s *traefik.Service) dto.ServiceDTO {
	return dto.ServiceDTO{
		Name:         s.Name,
		Type:         s.Type,
		Provider:     s.Provider,
		Status:       s.Status,
		UsedBy:       s.UsedBy,
		ServerStatus: s.ServerStatus,
	}
}
