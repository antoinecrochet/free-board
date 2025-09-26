package port

import "github.com/antoinecrochet/free-board/internal/core/model"

//go:generate mockgen -source=secondary.go -destination=mock/secondary.go

type TimeSheetPort interface {
	FindByUserID(userId int64) ([]*model.TimeSheet, error)
	Save(timeSheet *model.TimeSheet) error
}
