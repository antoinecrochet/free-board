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

func (b *Board) GetTimeSheets(username string, from string, to string) ([]*model.TimeSheet, error) {
	timeSheets, err := b.timeSheetPort.FindByUsername(username, from, to)
	if err != nil {
		slog.Error("Error getting timesheets for user", "username", username, "error", err)
		return nil, err
	}
	return timeSheets, nil
}

func (b *Board) GetTimeSheet(username string, timeSheetID int64) (*model.TimeSheet, error) {
	timeSheet, err := b.timeSheetPort.FindByID(timeSheetID)
	if err != nil {
		slog.Error("Error finding timesheet by id", "timeSheetID", timeSheetID, "error", err)
		return nil, err
	}
	if timeSheet == nil || timeSheet.Username != username {
		return nil, &model.NotFoundError{Code: "timesheet-not-found"}
	}
	return timeSheet, nil
}

func (b *Board) SaveTimeSheet(username string, day string, hours float64) (int64, error) {
	timeSheet, err := b.timeSheetPort.FindByUsernameAndDay(username, day)
	if err != nil {
		slog.Error("Error finding timesheet for user", "username", username, "day", day, "error", err)
		return -1, err
	}
	if timeSheet != nil {
		return -1, &model.AlreadExistsError{Code: "timesheet-already-exists"}
	}

	return b.timeSheetPort.Save(&model.TimeSheet{Username: username, Day: day, Hours: hours})
}

func (b *Board) UpdateTimeSheetHours(username string, timeSheetID int64, hours float64) error {
	timeSheet, err := b.timeSheetPort.FindByID(timeSheetID)
	if err != nil {
		slog.Error("Error finding timesheet by id", "timeSheetID", timeSheetID, "error", err)
		return err
	}
	if timeSheet == nil || timeSheet.Username != username {
		return &model.NotFoundError{Code: "timesheet-not-found"}
	}

	// Update hours
	timeSheet.Hours = hours
	return b.timeSheetPort.Update(timeSheet)
}

func (b *Board) DeleteTimeSheet(username string, timeSheetID int64) error {
	timeSheet, err := b.timeSheetPort.FindByID(timeSheetID)
	if err != nil {
		slog.Error("Error finding timesheet by id", "timeSheetID", timeSheetID, "error", err)
		return err
	}
	if timeSheet == nil || timeSheet.Username != username {
		return &model.NotFoundError{Code: "timesheet-not-found"}
	}

	return b.timeSheetPort.Delete(timeSheetID)
}
