package api

import (
	"fmt"
	"log/slog"

	"github.com/antoinecrochet/free-board/internal/core/port"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Application struct {
	board port.BoardManager
}

func NewApplication(board port.BoardManager) *Application {
	return &Application{board: board}
}

func (a *Application) StartServer(keycloakServerUrl string, keycloakRealm string) (err error) {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Location"},
		AllowCredentials: true,
	}))

	jwkSet, err := GetJWKSet(keycloakServerUrl, keycloakRealm)
	if err != nil {
		slog.Error("Failed to get JWK set", slog.String("error", err.Error()))
		return fmt.Errorf("failed to get JWK set: %w", err)
	}

	router.GET("/health", a.HealthCheck)

	apiV1 := router.Group("/api/v1")
	apiV1.Use(JWTMiddleware(keycloakServerUrl, keycloakRealm, jwkSet))
	{
		apiV1.GET("/timesheets", a.GetTimeSheets)
		apiV1.GET("/timesheets/:id", a.GetTimeSheet)
		apiV1.POST("/timesheets", a.CreateTimeSheet)
		apiV1.PATCH("/timesheets/:id", a.PatchTimeSheet)
		apiV1.DELETE("/timesheets/:id", a.DeleteTimeSheet)
	}

	return router.Run()
}
