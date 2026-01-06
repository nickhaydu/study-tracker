package models

import (
	"database/sql"
	"time"
)

type Stats struct {
	Topic         string
	TotalHours    float64
	Sessions      int
	AvgPerSession float64
	LastStudied   time.Time
}

type StatsModel struct {
	DB *sql.DB
}

func (m *StatsModel) GetSessionStats() ([]*Stats, error) {
	stmt := `
	SELECT topic, 
	SUM(duration) as total_hours, 
	COUNT(*) as time_studied, 
	AVG(duration) as average_duration,
	MAX(study_date) as last_time
	FROM study_sessions 
	GROUP BY topic`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	stats := []*Stats{}

	for rows.Next() {
		s := &Stats{}
		err := rows.Scan(&s.Topic, &s.TotalHours, &s.Sessions, &s.AvgPerSession, &s.LastStudied)
		if err != nil {
			return nil, err
		}

		stats = append(stats, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return stats, nil
}
