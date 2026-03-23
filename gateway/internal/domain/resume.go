package domain

import "time"

type Resume struct {
	ID        string
	UserID    string
	FileURL   string
	CreatedAt time.Time
}
