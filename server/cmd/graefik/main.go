package main

import (
	"context"
	"log"

	httpadapter "back-dashboard/internal/adapters/http"
	gormdb "back-dashboard/internal/adapters/persistence/gorm"
	"back-dashboard/internal/adapters/persistence/gorm/migrations"
	"back-dashboard/internal/adapters/security"
	"back-dashboard/internal/app/auth"
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
	hasher := security.NewBcryptHasher(0) // bcrypt.DefaultCost
	passGen := security.NewPasswordGenerator()

	tokens, err := security.NewJWTIssuer(cfg.JWTSecret, cfg.JWTTTL)
	if err != nil {
		log.Fatalf("jwt: %v", err)
	}

	// ── Use cases ────────────────────────────────────────────────
	bootstrap := auth.NewBootstrap(userRepo, hasher, passGen, cfg.AdminUsername, cfg.AdminPassLength)
	if err := bootstrap.EnsureAdmin(context.Background()); err != nil {
		log.Fatalf("bootstrap admin: %v", err)
	}

	loginService := auth.NewLoginService(userRepo, hasher, tokens)

	// ── Adapters inbound ─────────────────────────────────────────
	authHandler := httpadapter.NewAuthHandler(loginService)

	r := httpadapter.NewRouter(httpadapter.Deps{
		Auth:         authHandler,
		AllowOrigins: cfg.AllowOrigins,
	})

	log.Printf("Graefik écoute sur :%s", cfg.HTTPPort)
	if err := r.Run(":" + cfg.HTTPPort); err != nil {
		log.Fatalf("http: %v", err)
	}
}
