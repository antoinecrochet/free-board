package service

import (
	"github.com/antoinecrochet/free-board/internal/core/model"
	"github.com/antoinecrochet/free-board/internal/core/port"
)

type Board struct {
	timeSheetPort port.TimeSheetPort
}

func NewBoard(timeSheetPort port.TimeSheetPort) *Board {
	return &Board{timeSheetPort: timeSheetPort}
}

func (b *Board) SaveTimeSheet(userId int64, day string, hours float64) error {
	return b.timeSheetPort.Save(&model.TimeSheet{UserID: userId, Day: day, Hours: hours})
}
