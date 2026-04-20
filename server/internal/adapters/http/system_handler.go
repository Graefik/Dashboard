package httpadapter

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"back-dashboard/internal/adapters/http/dto"
)

type SystemHandler struct {
	traefikURL string
}

func NewSystemHandler(traefikURL string) *SystemHandler {
	return &SystemHandler{traefikURL: traefikURL}
}

func (h *SystemHandler) Info(c *gin.Context) {
	c.JSON(http.StatusOK, dto.SystemInfoResponse{
		TraefikURL: h.traefikURL,
	})
}
