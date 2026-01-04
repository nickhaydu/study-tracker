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
	return nil, nil
}
