package models

import (
	"database/sql"
	"errors"
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
	stmt := `SELECT id, topic, duration, study_date, notes FROM study_sessions
	WHERE id = ?`

	row := m.DB.QueryRow(stmt, id)

	s := &Session{}
	err := row.Scan(&s.ID, &s.Topic, &s.Duration, &s.StudyDate, &s.Notes)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}

func (m *SessionModel) GetLatestSessions() ([]*Session, error) {
	stmt := `SELECT id, topic, duration, study_date, notes FROM study_sessions
	ORDER BY id desc LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	sessions := []*Session{}

	for rows.Next() {
		s := &Session{}
		err = rows.Scan(&s.ID, &s.Topic, &s.Duration, &s.StudyDate, &s.Notes)
		if err != nil {
			return nil, err
		}

		sessions = append(sessions, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return sessions, nil
}
