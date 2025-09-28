package api

import (
	"github.com/antoinecrochet/free-board/internal/core/model"
)

func MapToApi(timeSheets []*model.TimeSheet) []TimeSheetResponse {
	responses := make([]TimeSheetResponse, len(timeSheets))
	for i, ts := range timeSheets {
		responses[i] = TimeSheetResponse{
			Day:   ts.Day,
			Hours: ts.Hours,
		}
	}
	return responses
}
