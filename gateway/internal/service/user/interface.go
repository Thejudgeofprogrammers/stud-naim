package user

import (
	"context"
	"gateway/internal/transport/http_gin/dto"
)

type UserService interface {
	GetStudent(ctx context.Context, id string) (*dto.StudentResponse, error)
	UpdateStudent(ctx context.Context, id string, req dto.UpdateStudentRequest) error

	GetEmployer(ctx context.Context, id string) (*dto.EmployerResponse, error)
	UpdateEmployer(ctx context.Context, id string, req dto.UpdateEmployerRequest) error

	ListStudents(ctx context.Context) ([]dto.StudentResponse, error)
	ListEmployers(ctx context.Context) ([]dto.EmployerResponse, error)
}
