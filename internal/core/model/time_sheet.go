package model

type TimeSheet struct {
	ID     int64
	UserID int64
	Day    string
	Hours  float64
}
