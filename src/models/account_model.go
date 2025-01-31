package models

import "time"

type Account struct {
	ID        int64
	Username  string
	Password  string
	Mobile    string
	Email     string
	IsPremium bool
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
