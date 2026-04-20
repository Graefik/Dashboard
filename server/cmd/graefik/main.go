package main

import (
	"context"
	"log"

	httpadapter "back-dashboard/internal/adapters/http"
	"back-dashboard/internal/adapters/memstore"
	gormdb "back-dashboard/internal/adapters/persistence/gorm"
	"back-dashboard/internal/adapters/persistence/gorm/migrations"
	prometheusclient "back-dashboard/internal/adapters/prometheus"
	"back-dashboard/internal/adapters/security"
	traefikclient "back-dashboard/internal/adapters/traefik"
	"back-dashboard/internal/app/auth"
	metricsapp "back-dashboard/internal/app/metrics"
	"back-dashboard/internal/app/observe"
	"back-dashboard/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	// ── Infrastructure ────────────────────────────────────────────
	db, err := gormdb.Open(gormdb.Config{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		User:     cfg.DB.User,
		Password: cfg.DB.Password,
		Name:     cfg.DB.Name,
	})
	if err != nil {
		log.Fatalf("db: %v", err)
	}

	if err := migrations.Run(db); err != nil {
		log.Fatalf("migrations: %v", err)
	}

	// ── Adapters outbound ────────────────────────────────────────
	userRepo := gormdb.NewUserRepo(db)
	hasher := security.NewBcryptHasher(0)
	passGen := security.NewPasswordGenerator()

	tokens, err := security.NewJWTIssuer(cfg.JWTSecret, cfg.JWTTTL)
	if err != nil {
		log.Fatalf("jwt: %v", err)
	}

	traefikCli, err := traefikclient.New(traefikclient.Config{
		BaseURL:  cfg.Traefik.BaseURL,
		Username: cfg.Traefik.Username,
		Password: cfg.Traefik.Password,
		Timeout:  cfg.Traefik.Timeout,
	})
	if err != nil {
		log.Fatalf("traefik client: %v", err)
	}

	// ── Use cases (auth, observe) ────────────────────────────────
	bootstrap := auth.NewBootstrap(userRepo, hasher, passGen, cfg.AdminUsername, cfg.AdminPassLength)
	if err := bootstrap.EnsureAdmin(context.Background()); err != nil {
		log.Fatalf("bootstrap admin: %v", err)
	}

	loginService := auth.NewLoginService(userRepo, hasher, tokens)
	observeService := observe.NewService(traefikCli)

	// ── Metrics (optionnel : activé uniquement si URL fournie) ──
	var metricsHandler *httpadapter.MetricsHandler
	if cfg.Metrics.URL != "" {
		promCli, err := prometheusclient.New(prometheusclient.Config{
			URL:      cfg.Metrics.URL,
			Username: cfg.Metrics.Username,
			Password: cfg.Metrics.Password,
			Timeout:  cfg.Metrics.Timeout,
		})
		if err != nil {
			log.Fatalf("prometheus client: %v", err)
		}

		capacity := cfg.Metrics.RetentionMinutes * 60 / int(cfg.Metrics.ScrapeInterval.Seconds())
		store := memstore.NewRingBuffer(capacity)

		scraper := metricsapp.NewScraper(promCli, store, cfg.Metrics.ScrapeInterval)
		scraper.Start(context.Background())

		metricsService := metricsapp.NewService(store)
		metricsHandler = httpadapter.NewMetricsHandler(metricsService)

		log.Printf("metrics: scrape %s every %s (retention %dm)",
			cfg.Metrics.URL, cfg.Metrics.ScrapeInterval, cfg.Metrics.RetentionMinutes)
	} else {
		log.Printf("metrics: désactivé (TRAEFIK_METRICS_URL non défini)")
	}

	// ── Router ───────────────────────────────────────────────────
	r := httpadapter.NewRouter(httpadapter.Deps{
		Auth:         httpadapter.NewAuthHandler(loginService),
		Observe:      httpadapter.NewObserveHandler(observeService),
		Metrics:      metricsHandler,
		System:       httpadapter.NewSystemHandler(cfg.Traefik.BaseURL),
		JWTSecret:    cfg.JWTSecret,
		AllowOrigins: cfg.AllowOrigins,
	})

	log.Printf("Graefik écoute sur :%s", cfg.HTTPPort)
	if err := r.Run(":" + cfg.HTTPPort); err != nil {
		log.Fatalf("http: %v", err)
	}
}
