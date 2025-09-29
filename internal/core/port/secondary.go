package port

import "github.com/antoinecrochet/free-board/internal/core/model"

//go:generate mockgen -source=secondary.go -destination=mock/secondary.go

type TimeSheetPort interface {
	FindByID(id int64) (*model.TimeSheet, error)
	FindByUserID(userId int64) ([]*model.TimeSheet, error)
	FindByUserIDAndDay(userId int64, day string) (*model.TimeSheet, error)
	Save(timeSheet *model.TimeSheet) error
	Update(timeSheet *model.TimeSheet) error
	Delete(id int64) error
}
