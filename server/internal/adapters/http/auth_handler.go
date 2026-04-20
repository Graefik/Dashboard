package httpadapter

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"back-dashboard/internal/adapters/http/dto"
	"back-dashboard/internal/domain/user"
	"back-dashboard/internal/ports/inbound"
)

type AuthHandler struct {
	auth inbound.AuthUseCase
}

func NewAuthHandler(auth inbound.AuthUseCase) *AuthHandler {
	return &AuthHandler{auth: auth}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: "payload invalide",
			Code:  "INVALID_PAYLOAD",
		})
		return
	}

	res, err := h.auth.Login(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		if errors.Is(err, user.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
				Error: "identifiants incorrects",
				Code:  "INVALID_CREDENTIALS",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: "erreur interne",
			Code:  "INTERNAL",
		})
		return
	}

	c.JSON(http.StatusOK, dto.LoginResponse{
		Token:     res.Token,
		Username:  res.Username,
		ExpiresAt: res.ExpiresAt,
	})
}
