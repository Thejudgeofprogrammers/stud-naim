package domain

import (
	"time"
)

type Role string

const (
	RoleStudent  Role = "student"
	RoleEmployer Role = "employer"
	RoleCurator  Role = "curator"
	RoleAdmin    Role = "admin"
)

type User struct {
	ID        string
	Email     string
	Password  string
	Role      Role
	CreatedAt time.Time
}

type StudentProfile struct {
	ID        string
	UserID    string
	FullName  string
	Skills    []string
	About     string
	ResumeID  *string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type EmployerProfile struct {
	ID             string
	UserID         string
	CompanyName    string
	Description    string
	Representative string
	Verified       bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
