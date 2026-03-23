package repository

import (
	"context"
	"gateway/internal/domain"
)

type ProfileRepository interface {
	CreateStudent(ctx context.Context, p *domain.StudentProfile) error
	GetStudent(ctx context.Context, userID string) (*domain.StudentProfile, error)
	UpdateStudent(ctx context.Context, p *domain.StudentProfile) error
	ListStudents(ctx context.Context) ([]*domain.StudentProfile, error)

	CreateEmployer(ctx context.Context, p *domain.EmployerProfile) error
	GetEmployer(ctx context.Context, userID string) (*domain.EmployerProfile, error)
	UpdateEmployer(ctx context.Context, p *domain.EmployerProfile) error
	ListEmployers(ctx context.Context) ([]*domain.EmployerProfile, error)
}
