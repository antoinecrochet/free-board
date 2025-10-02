package api

import (
	"github.com/antoinecrochet/free-board/internal/core/model"
)

func MapTimeSheetArrayToApi(timeSheets []*model.TimeSheet) []TimeSheetResponse {
	responses := make([]TimeSheetResponse, len(timeSheets))
	for i, ts := range timeSheets {
		responses[i] = MapTimeSheetToApi(ts)
	}
	return responses
}

func MapTimeSheetToApi(timeSheet *model.TimeSheet) TimeSheetResponse {
	return TimeSheetResponse{
		ID:    timeSheet.ID,
		Day:   timeSheet.Day,
		Hours: timeSheet.Hours,
	}
}
