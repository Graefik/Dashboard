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
	Traefik         Traefik
	Metrics         Metrics
}

type Traefik struct {
	BaseURL  string
	Username string
	Password string
	Timeout  time.Duration
}

type Metrics struct {
	URL              string
	Username         string
	Password         string
	Timeout          time.Duration
	ScrapeInterval   time.Duration
	RetentionMinutes int
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
		Traefik: Traefik{
			BaseURL:  getEnv("TRAEFIK_API_URL", ""),
			Username: getEnv("TRAEFIK_API_USERNAME", ""),
			Password: getEnv("TRAEFIK_API_PASSWORD", ""),
			Timeout:  time.Duration(getEnvInt("TRAEFIK_API_TIMEOUT_SEC", 5)) * time.Second,
		},
		Metrics: Metrics{
			URL:              getEnv("TRAEFIK_METRICS_URL", ""),
			Username:         getEnv("TRAEFIK_METRICS_USERNAME", ""),
			Password:         getEnv("TRAEFIK_METRICS_PASSWORD", ""),
			Timeout:          time.Duration(getEnvInt("TRAEFIK_METRICS_TIMEOUT_SEC", 5)) * time.Second,
			ScrapeInterval:   time.Duration(getEnvInt("METRICS_SCRAPE_INTERVAL_SEC", 10)) * time.Second,
			RetentionMinutes: getEnvInt("METRICS_RETENTION_MINUTES", 60),
		},
	}

	if cfg.JWTSecret == "" {
		return nil, fmt.Errorf("JWT_SECRET est requis (>= 16 caractères)")
	}
	if len(cfg.JWTSecret) < 16 {
		return nil, fmt.Errorf("JWT_SECRET doit faire >= 16 caractères")
	}
	if cfg.Traefik.BaseURL == "" {
		return nil, fmt.Errorf("TRAEFIK_API_URL est requis")
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
