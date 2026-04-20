package httpadapter

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"back-dashboard/internal/adapters/http/middleware"
)

type Deps struct {
	Auth         *AuthHandler
	Observe      *ObserveHandler
	Metrics      *MetricsHandler
	System       *SystemHandler
	JWTSecret    string
	AllowOrigins []string
}

func NewRouter(deps Deps) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     deps.AllowOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	api := r.Group("/api")
	{
		// ── Public (auth) ───────────────────────────────
		auth := api.Group("/auth")
		auth.POST("/login", deps.Auth.Login)

		// ── Protégé (Bearer JWT) ────────────────────────
		protected := api.Group("")
		protected.Use(middleware.JWTAuth(deps.JWTSecret))
		{
			protected.GET("/overview", deps.Observe.Overview)

			httpGroup := protected.Group("/http")
			httpGroup.GET("/routers", deps.Observe.Routers)
			httpGroup.GET("/services", deps.Observe.Services)

			if deps.Metrics != nil {
				metricsGroup := protected.Group("/metrics")
				metricsGroup.GET("/overview", deps.Metrics.Overview)
			}

			system := protected.Group("/system")
			system.GET("/info", deps.System.Info)
		}
	}

	return r
}
