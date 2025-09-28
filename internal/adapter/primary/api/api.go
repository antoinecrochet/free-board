package api

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/antoinecrochet/free-board/internal/core/model"
	"github.com/antoinecrochet/free-board/internal/core/port"
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
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "UP",
		})
	})

	router.POST("/timesheets", func(c *gin.Context) {
		var req CreateTimeSheetRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// For simplicity, we use a hardcoded user ID
		if err := a.board.SaveTimeSheet(1, req.Day, req.Hours); err != nil {
			if _, ok := err.(*model.AlreadExistsError); ok {
				c.JSON(http.StatusConflict, ErrorResponse{Error: err.(*model.AlreadExistsError).ErrorCode()})
				return
			}
			slog.Error("Error while saving timesheet", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal-server-error"})
			return
		}
		c.Status(http.StatusCreated)
	})

	router.GET("/timesheets", func(c *gin.Context) {
		// For simplicity, we use a hardcoded user ID
		timeSheets, err := a.board.GetTimeSheets(1)
		if err != nil {
			slog.Error("Error while getting timesheets", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal-server-error"})
			return
		}
		c.JSON(http.StatusOK, TimeSheetsResponse{TimeSheets: MapToApi(timeSheets)})
	})

	router.PATCH("/timesheets/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid-id"})
			return
		}

		var req UpdateTimeSheetHoursRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// For simplicity, we use a hardcoded user ID
		if err := a.board.UpdateTimeSheetHours(1, int64(id), req.Hours); err != nil {
			if _, ok := err.(*model.NotFoundError); ok {
				c.JSON(http.StatusNotFound, ErrorResponse{Error: err.(*model.NotFoundError).ErrorCode()})
				return
			}
			slog.Error("Error while updating timesheet", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal-server-error"})
			return
		}
		c.Status(http.StatusNoContent)
	})

	return router.Run()
}
