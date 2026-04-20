package prometheusclient

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/prometheus/common/expfmt"
	"github.com/prometheus/common/model"

	"back-dashboard/internal/domain/metrics"
)

// prometheus/common@0.67+ requiert une ValidationScheme explicite, sinon le
// parser panique dès le premier # HELP. LegacyValidation = comportement v0.62-.
func init() {
	model.NameValidationScheme = model.LegacyValidation
}

type Config struct {
	URL      string // ex: http://127.0.0.1:8082/metrics
	Username string
	Password string
	Timeout  time.Duration
}

type Client struct {
	cfg  Config
	http *http.Client
}

func New(cfg Config) (*Client, error) {
	if strings.TrimSpace(cfg.URL) == "" {
		return nil, errors.New("prometheus: URL required")
	}
	if cfg.Timeout <= 0 {
		cfg.Timeout = 5 * time.Second
	}
	return &Client{
		cfg:  cfg,
		http: &http.Client{Timeout: cfg.Timeout},
	}, nil
}

func (c *Client) ScrapeOnce(ctx context.Context) (*metrics.Snapshot, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.cfg.URL, nil)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}
	if c.cfg.Username != "" {
		req.SetBasicAuth(c.cfg.Username, c.cfg.Password)
	}
	req.Header.Set("Accept", string(expfmt.NewFormat(expfmt.TypeTextPlain)))

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("scrape: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("scrape: status %d", resp.StatusCode)
	}

	// Sécurité : certaines versions de prometheus/common redéfinissent la
	// scheme après le init() ; on la re-force juste avant le parse.
	model.NameValidationScheme = model.LegacyValidation
	parser := expfmt.TextParser{}
	families, err := parser.TextToMetricFamilies(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parse: %w", err)
	}

	snap := &metrics.Snapshot{At: time.Now().UTC()}
	parseSnapshot(snap, families)
	return snap, nil
}
