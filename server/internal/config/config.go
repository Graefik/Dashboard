package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	HTTPPort        string
	AllowOrigins    []string
	DB              DB
	AdminUsername   string
	AdminPassLength int
	JWTSecret       string
	JWTTTL          time.Duration
}

type DB struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func Load() (*Config, error) {
	cfg := &Config{
		HTTPPort:        getEnv("HTTP_PORT", "8080"),
		AllowOrigins:    splitCSV(getEnv("CORS_ALLOW_ORIGINS", "http://localhost:5173,http://localhost:5178")),
		AdminUsername:   getEnv("ADMIN_USERNAME", "admin"),
		AdminPassLength: getEnvInt("ADMIN_PASSWORD_LENGTH", 32),
		JWTSecret:       getEnv("JWT_SECRET", ""),
		JWTTTL:          time.Duration(getEnvInt("JWT_TTL_HOURS", 24)) * time.Hour,
		DB: DB{
			Host:     getEnv("DB_HOST", "mysql"),
			Port:     getEnv("DB_PORT", "3306"),
			User:     getEnv("DB_USER", "dashboard"),
			Password: getEnv("DB_PASSWORD", "dashboard_pwd"),
			Name:     getEnv("DB_NAME", "dashboard_dev"),
		},
	}

	if cfg.JWTSecret == "" {
		return nil, fmt.Errorf("JWT_SECRET est requis (>= 16 caractères)")
	}
	if len(cfg.JWTSecret) < 16 {
		return nil, fmt.Errorf("JWT_SECRET doit faire >= 16 caractères")
	}
	return cfg, nil
}

func getEnv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

func getEnvInt(k string, def int) int {
	v := os.Getenv(k)
	if v == "" {
		return def
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return def
	}
	return n
}

func splitCSV(s string) []string {
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		if t := strings.TrimSpace(p); t != "" {
			out = append(out, t)
		}
	}
	return out
}
