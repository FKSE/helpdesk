package model

import "time"

const (
	RoleSuperAdmin = iota
	RoleAdmin      = iota
	RoleSupporter  = iota
	RoleCustomer   = iota
)

type (
	// Help desk user
	User struct {
		ID        int
		Username  string
		FirstName string
		LastName  string
		Email     string
		Password  string
		Salt      string
		Role      int
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	// Products
	Product struct {
		ID           int
		AccountingId string
		CreatedAt    time.Time
		UpdatedAt    time.Time
	}
	// Projects
	Project struct {
		ID          int
		Name        string
		Description string
		Products    []Product
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
)
