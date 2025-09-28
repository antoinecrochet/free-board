package service

import (
	"log/slog"

	"github.com/antoinecrochet/free-board/internal/core/model"
	"github.com/antoinecrochet/free-board/internal/core/port"
)

type Board struct {
	timeSheetPort port.TimeSheetPort
}

func NewBoard(timeSheetPort port.TimeSheetPort) *Board {
	return &Board{timeSheetPort: timeSheetPort}
}

func (b *Board) GetTimeSheets(userId int64) ([]*model.TimeSheet, error) {
	timeSheets, err := b.timeSheetPort.FindByUserID(userId)
	if err != nil {
		slog.Error("Error getting timesheets for user", "userId", userId, "error", err)
		return nil, err
	}
	return timeSheets, nil
}

func (b *Board) GetTimeSheet(userId int64, timeSheetID int64) (*model.TimeSheet, error) {
	timeSheet, err := b.timeSheetPort.FindByID(timeSheetID)
	if err != nil {
		slog.Error("Error finding timesheet by id", "timeSheetID", timeSheetID, "error", err)
		return nil, err
	}
	if timeSheet == nil || timeSheet.UserID != userId {
		return nil, &model.NotFoundError{Code: "timesheet-not-found"}
	}
	return timeSheet, nil
}

func (b *Board) SaveTimeSheet(userId int64, day string, hours float64) error {
	timeSheet, err := b.timeSheetPort.FindByUserIDAndDay(userId, day)
	if err != nil {
		slog.Error("Error finding timesheet for user", "userId", userId, "day", day, "error", err)
		return err
	}
	if timeSheet != nil {
		return &model.AlreadExistsError{Code: "timesheet-already-exists"}
	}

	return b.timeSheetPort.Save(&model.TimeSheet{UserID: userId, Day: day, Hours: hours})
}

func (b *Board) UpdateTimeSheetHours(userId int64, timeSheetID int64, hours float64) error {
	timeSheet, err := b.timeSheetPort.FindByID(timeSheetID)
	if err != nil {
		slog.Error("Error finding timesheet by id", "timeSheetID", timeSheetID, "error", err)
		return err
	}
	if timeSheet == nil || timeSheet.UserID != userId {
		return &model.NotFoundError{Code: "timesheet-not-found"}
	}

	// Update hours
	timeSheet.Hours = hours
	return b.timeSheetPort.Update(timeSheet)
}
