package port

import "github.com/antoinecrochet/free-board/internal/core/model"

//go:generate mockgen -source=service.go -destination=mock/service.go

type BoardManager interface {
	GetTimeSheets(userId int64, from string, to string) ([]*model.TimeSheet, error)
	GetTimeSheet(userId int64, timeSheetID int64) (*model.TimeSheet, error)
	SaveTimeSheet(userId int64, day string, hours float64) (int64, error)
	UpdateTimeSheetHours(userId int64, timeSheetID int64, hours float64) error
	DeleteTimeSheet(userId int64, timeSheetID int64) error
}
