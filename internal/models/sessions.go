package models

import (
	"database/sql"
	"fmt"
	"time"
)

type Session struct {
	ID        int
	Topic     string
	Duration  int
	StudyDate time.Time
	Notes     string
}

type SessionModel struct {
	DB *sql.DB
}

func (m *SessionModel) InsertSession(topic string, duration int, notes string) (int, error) {
	stmt := `INSERT INTO study_sessions(topic, duration, study_date, notes)
	VALUES(?, ?, UTC_TIMESTAMP(), ?)`

	result, err := m.DB.Exec(stmt, topic, duration, notes)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	fmt.Println(int(id))

	return int(id), nil
}

func (m *SessionModel) GetSession(id int) (*Session, error) {
	return nil, nil
}

func (m *SessionModel) GetLatestSession() ([]*Session, error) {
	return nil, nil
}
