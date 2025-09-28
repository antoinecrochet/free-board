package port

import "github.com/antoinecrochet/free-board/internal/core/model"

//go:generate mockgen -source=service.go -destination=mock/service.go

type BoardManager interface {
	GetTimeSheets(userId int64) ([]*model.TimeSheet, error)
	SaveTimeSheet(userId int64, day string, hours float64) error
	UpdateTimeSheet(userId int64, day string, hours float64) error
}
