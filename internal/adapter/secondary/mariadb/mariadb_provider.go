package mariadb

import (
	"database/sql"
	"fmt"

	"github.com/antoinecrochet/free-board/internal/core/model"
	_ "github.com/go-sql-driver/mysql"
)

type MariaDbProvider struct {
	db *sql.DB
}

func NewMariaDbProvider(user string, password string, databaseName string) *MariaDbProvider {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", user, password, databaseName))
	if err != nil {
		panic(err)
	}

	return &MariaDbProvider{db: db}
}

func (m *MariaDbProvider) FindByID(id int64) (*model.TimeSheet, error) {
	row := m.db.QueryRow("SELECT id, user_id, day, hours FROM timesheet WHERE id = ?", id)
	ts := new(model.TimeSheet)
	if err := row.Scan(&ts.ID, &ts.UserID, &ts.Day, &ts.Hours); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return ts, nil
}

func (m *MariaDbProvider) FindByUserID(userId int64, from string, to string) ([]*model.TimeSheet, error) {
	// Set default values for from and to if empty
	from = defaultIfEmpty(from, "0000-01-01")
	to = defaultIfEmpty(to, "9999-12-31")

	rows, err := m.db.Query("SELECT id, user_id, day, hours FROM timesheet WHERE user_id = ? AND day BETWEEN ? AND ?", userId, from, to)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var timeSheets []*model.TimeSheet
	for rows.Next() {
		ts := new(model.TimeSheet)
		if err := rows.Scan(&ts.ID, &ts.UserID, &ts.Day, &ts.Hours); err != nil {
			return nil, err
		}
		timeSheets = append(timeSheets, ts)
	}

	return timeSheets, nil
}

func defaultIfEmpty(value string, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}

func (m *MariaDbProvider) FindByUserIDAndDay(userId int64, day string) (*model.TimeSheet, error) {
	row := m.db.QueryRow("SELECT id, user_id, day, hours FROM timesheet WHERE user_id = ? AND day = ?", userId, day)
	ts := new(model.TimeSheet)
	if err := row.Scan(&ts.ID, &ts.UserID, &ts.Day, &ts.Hours); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return ts, nil
}

func (m *MariaDbProvider) Save(timeSheet *model.TimeSheet) error {
	_, err := m.db.Exec("INSERT INTO timesheet (user_id, day, hours) VALUES (?, ?, ?)", timeSheet.UserID, timeSheet.Day, timeSheet.Hours)
	return err
}

func (m *MariaDbProvider) Update(timeSheet *model.TimeSheet) error {
	_, err := m.db.Exec("UPDATE timesheet SET user_id = ?, day = ?, hours = ? WHERE id = ?", timeSheet.UserID, timeSheet.Day, timeSheet.Hours, timeSheet.ID)
	return err
}

func (m *MariaDbProvider) Delete(id int64) error {
	_, err := m.db.Exec("DELETE FROM timesheet WHERE id = ?", id)
	return err
}
