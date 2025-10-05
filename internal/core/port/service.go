package port

import "github.com/antoinecrochet/free-board/internal/core/model"

//go:generate mockgen -source=service.go -destination=mock/service.go

type BoardManager interface {
	GetTimeSheets(username string, from string, to string) ([]*model.TimeSheet, error)
	GetTimeSheet(username string, timeSheetID int64) (*model.TimeSheet, error)
	SaveTimeSheet(username string, day string, hours float64) (int64, error)
	UpdateTimeSheetHours(username string, timeSheetID int64, hours float64) error
	DeleteTimeSheet(username string, timeSheetID int64) error
}
