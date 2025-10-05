package port

import "github.com/antoinecrochet/free-board/internal/core/model"

//go:generate mockgen -source=secondary.go -destination=mock/secondary.go

type TimeSheetPort interface {
	FindByID(id int64) (*model.TimeSheet, error)
	FindByUsername(username string, from string, to string) ([]*model.TimeSheet, error)
	FindByUsernameAndDay(username string, day string) (*model.TimeSheet, error)
	Save(timeSheet *model.TimeSheet) (int64, error)
	Update(timeSheet *model.TimeSheet) error
	Delete(id int64) error
}
