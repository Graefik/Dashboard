package traefikclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"back-dashboard/internal/domain/traefik"
)

type Config struct {
	BaseURL  string // ex: https://traefik.maillet.bzh
	Username string
	Password string
	Timeout  time.Duration
}

type Client struct {
	cfg  Config
	http *http.Client
}

func New(cfg Config) (*Client, error) {
	if cfg.BaseURL == "" {
		return nil, errors.New("traefik: base URL required")
	}
	cfg.BaseURL = strings.TrimRight(cfg.BaseURL, "/")
	if cfg.Timeout <= 0 {
		cfg.Timeout = 5 * time.Second
	}
	return &Client{
		cfg:  cfg,
		http: &http.Client{Timeout: cfg.Timeout},
	}, nil
}

func (c *Client) GetOverview(ctx context.Context) (*traefik.Overview, error) {
	var dto overviewDTO
	if err := c.getJSON(ctx, "/api/overview", &dto); err != nil {
		return nil, err
	}
	version, _ := c.fetchVersion(ctx)
	return toDomainOverview(&dto, version), nil
}

func (c *Client) ListRouters(ctx context.Context) ([]traefik.Router, error) {
	var dto []routerDTO
	if err := c.getJSON(ctx, "/api/http/routers", &dto); err != nil {
		return nil, err
	}
	out := make([]traefik.Router, 0, len(dto))
	for _, r := range dto {
		out = append(out, toDomainRouter(&r))
	}
	return out, nil
}

func (c *Client) ListServices(ctx context.Context) ([]traefik.Service, error) {
	var dto []serviceDTO
	if err := c.getJSON(ctx, "/api/http/services", &dto); err != nil {
		return nil, err
	}
	out := make([]traefik.Service, 0, len(dto))
	for _, s := range dto {
		out = append(out, toDomainService(&s))
	}
	return out, nil
}

func (c *Client) fetchVersion(ctx context.Context) (string, error) {
	var v versionDTO
	if err := c.getJSON(ctx, "/api/version", &v); err != nil {
		return "", err
	}
	return v.Version, nil
}

func (c *Client) getJSON(ctx context.Context, path string, out any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.cfg.BaseURL+path, nil)
	if err != nil {
		return fmt.Errorf("new request: %w", err)
	}
	if c.cfg.Username != "" {
		req.SetBasicAuth(c.cfg.Username, c.cfg.Password)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("%w: %s", traefik.ErrUnavailable, err.Error())
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		// pass
	case http.StatusUnauthorized:
		return traefik.ErrUnauthorized
	case http.StatusNotFound:
		return traefik.ErrNotFound
	default:
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 1024))
		return fmt.Errorf("%w: status %d body=%s", traefik.ErrMalformed, resp.StatusCode, string(body))
	}

	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		return fmt.Errorf("%w: %s", traefik.ErrMalformed, err.Error())
	}
	return nil
}
