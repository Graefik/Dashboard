package security

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTIssuer struct {
	secret []byte
	ttl    time.Duration
	issuer string
}

func NewJWTIssuer(secret string, ttl time.Duration) (*JWTIssuer, error) {
	if len(secret) < 16 {
		return nil, errors.New("jwt: secret must be at least 16 chars")
	}
	if ttl <= 0 {
		ttl = 24 * time.Hour
	}
	return &JWTIssuer{
		secret: []byte(secret),
		ttl:    ttl,
		issuer: "graefik",
	}, nil
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (j *JWTIssuer) Issue(userID uint, username string) (string, time.Time, error) {
	now := time.Now().UTC()
	exp := now.Add(j.ttl)

	claims := Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.issuer,
			Subject:   fmt.Sprintf("%d", userID),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := t.SignedString(j.secret)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("sign token: %w", err)
	}
	return signed, exp, nil
}
