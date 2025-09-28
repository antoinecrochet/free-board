package api

type ErrorResponse struct {
	Error string `json:"error"`
}

type CreateTimeSheetRequest struct {
	Day   string  `json:"day" binding:"required,datetime=2006-01-02"`
	Hours float64 `json:"hours" binding:"required,gte=0,lte=24"`
}

type TimeSheetResponse struct {
	Day   string  `json:"day"`
	Hours float64 `json:"hours"`
}

type TimeSheetsResponse struct {
	TimeSheets []TimeSheetResponse `json:"timesheets"`
}
