package port

//go:generate mockgen -source=service.go -destination=mock/service.go

type BoardManager interface {
	SaveTimeSheet(userId int64, day string, hours float64) error
}
