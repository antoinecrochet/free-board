package api

import (
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

func (a *Application) StartServer() (err error) {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Location"},
		AllowCredentials: true,
	}))

	router.GET("/health", a.HealthCheck)

	router.GET("/timesheets", a.GetTimeSheets)
	router.GET("/timesheets/:id", a.GetTimeSheet)
	router.POST("/timesheets", a.CreateTimeSheet)
	router.PATCH("/timesheets/:id", a.PatchTimeSheet)
	router.DELETE("/timesheets/:id", a.DeleteTimeSheet)

	return router.Run()
}
