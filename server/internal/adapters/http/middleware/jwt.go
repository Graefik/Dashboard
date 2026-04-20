package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"back-dashboard/internal/adapters/http/dto"
)

const (
	CtxUsernameKey = "auth.username"
	CtxUserIDKey   = "auth.userID"
)

// JWTAuth retourne un middleware qui valide un token Bearer JWT
// signé avec le secret fourni. En cas d'échec, renvoie 401 JSON.
func JWTAuth(secret string) gin.HandlerFunc {
	secretBytes := []byte(secret)

	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			abort(c, "token manquant", "MISSING_TOKEN")
			return
		}

		token := strings.TrimPrefix(header, "Bearer ")
		if token == header {
			abort(c, "format Authorization invalide", "INVALID_AUTH_HEADER")
			return
		}

		claims := &jwt.RegisteredClaims{}
		parsed, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return secretBytes, nil
		})
		if err != nil || !parsed.Valid {
			abort(c, "token invalide ou expiré", "INVALID_TOKEN")
			return
		}

		// Claim "Username" custom non présent dans RegisteredClaims,
		// mais re-parsé pour éviter complexité : on stocke le Subject.
		c.Set(CtxUserIDKey, claims.Subject)
		c.Next()
	}
}

func abort(c *gin.Context, message, code string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
		Error: message,
		Code:  code,
	})
}
