package domain

import (
	"time"
)

type Role string

const (
	RoleStudent  Role = "student"
	RoleEmployer Role = "employer"
)

type User struct {
	ID        string
	Email     string
	Password  string
	Role      Role
	CreatedAt time.Time
}

type StudentProfile struct {
	UserID   string
	FullName string
	Skills   []string
	About    string
	ResumeID *string
}

type EmployerProfile struct {
	UserID         string
	CompanyName    string
	Description    string
	Representative string
}
