package api

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/antoinecrochet/free-board/internal/core/model"
	"github.com/gin-gonic/gin"
)

func (a *Application) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "UP",
	})
}

func (a *Application) GetTimeSheets(c *gin.Context) {
	var params GetTimeSheetsQueryParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username := c.GetString("username")
	timeSheets, err := a.board.GetTimeSheets(username, params.From, params.To)
	if err != nil {
		slog.Error("Error while getting timesheets for user", "username", username, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal-server-error"})
		return
	}
	c.JSON(http.StatusOK, TimeSheetsResponse{TimeSheets: MapTimeSheetArrayToApi(timeSheets)})
}

func (a *Application) GetTimeSheet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid-id"})
		return
	}

	username := c.GetString("username")
	timeSheet, err := a.board.GetTimeSheet(username, int64(id))
	if err != nil {
		if _, ok := err.(*model.NotFoundError); ok {
			c.JSON(http.StatusNotFound, ErrorResponse{Error: err.(*model.NotFoundError).ErrorCode()})
			return
		}
		slog.Error("Error while getting timesheet for user", "username", username, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal-server-error"})
		return
	}
	c.JSON(http.StatusOK, MapTimeSheetToApi(timeSheet))
}

func (a *Application) CreateTimeSheet(c *gin.Context) {
	var req CreateTimeSheetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username := c.GetString("username")
	id, err := a.board.SaveTimeSheet(username, req.Day, req.Hours)
	if err != nil {
		if _, ok := err.(*model.AlreadExistsError); ok {
			c.JSON(http.StatusConflict, ErrorResponse{Error: err.(*model.AlreadExistsError).ErrorCode()})
			return
		}
		slog.Error("Error while saving timesheet for user", "username", username, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal-server-error"})
		return
	}

	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}

	c.Header("Location", fmt.Sprintf("%s://%s/timesheets/%d", scheme, c.Request.Host, id))
	c.Status(http.StatusCreated)
}

func (a *Application) PatchTimeSheet(c *gin.Context) {
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

	username := c.GetString("username")
	if err := a.board.UpdateTimeSheetHours(username, int64(id), req.Hours); err != nil {
		if _, ok := err.(*model.NotFoundError); ok {
			c.JSON(http.StatusNotFound, ErrorResponse{Error: err.(*model.NotFoundError).ErrorCode()})
			return
		}
		slog.Error("Error while updating timesheet for user", "username", username, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal-server-error"})
		return
	}
	c.Status(http.StatusNoContent)
}

func (a *Application) DeleteTimeSheet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid-id"})
		return
	}

	username := c.GetString("username")
	if err := a.board.DeleteTimeSheet(username, int64(id)); err != nil {
		if _, ok := err.(*model.NotFoundError); ok {
			c.JSON(http.StatusNotFound, ErrorResponse{Error: err.(*model.NotFoundError).ErrorCode()})
			return
		}
		slog.Error("Error while deleting timesheet for user", "username", username, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal-server-error"})
		return
	}
	c.Status(http.StatusNoContent)
}
