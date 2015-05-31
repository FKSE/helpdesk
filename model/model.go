package model

import "time"

type User struct {
	ID        int
	Username  string
	FirstName string
	LastName  string
	Email     string
	Password  string
	Salt      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
