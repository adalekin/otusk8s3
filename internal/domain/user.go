package domain

import "time"

type User struct {
	ID        int64
	Name      string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
