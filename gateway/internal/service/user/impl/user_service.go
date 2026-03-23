package impl

import (
	"context"
	"gateway/internal/repository"
	"gateway/internal/service/user"
	"gateway/internal/transport/http_gin/dto"
)

type userService struct {
	userRepo    repository.UserRepository
	profileRepo repository.ProfileRepository
}

func NewUserService(
	userRepo repository.UserRepository,
	profileRepo repository.ProfileRepository,
) user.UserService {
	return &userService{
		userRepo:    userRepo,
		profileRepo: profileRepo,
	}
}

func (s *userService) GetStudent(ctx context.Context, id string) (*dto.StudentResponse, error) {
	return &dto.StudentResponse{
		ID:     id,
		Name:   "Test Student",
		Skills: []string{"Go", "Docker"},
		About:  "Backend developer",
	}, nil
}

func (s *userService) UpdateStudent(ctx context.Context, id string, req dto.UpdateStudentRequest) error {
	return nil
}

func (s *userService) GetEmployer(ctx context.Context, id string) (*dto.EmployerResponse, error) {
	return &dto.EmployerResponse{
		ID:             id,
		CompanyName:    "Test Company",
		Description:    "IT Company",
		Representative: "John Doe",
	}, nil
}

func (s *userService) UpdateEmployer(ctx context.Context, id string, req dto.UpdateEmployerRequest) error {
	return nil
}

func (s *userService) ListStudents(ctx context.Context) ([]dto.StudentResponse, error) {
	return []dto.StudentResponse{
		{
			ID:     "1",
			Name:   "Student 1",
			Skills: []string{"Go"},
		},
	}, nil
}

func (s *userService) ListEmployers(ctx context.Context) ([]dto.EmployerResponse, error) {
	return []dto.EmployerResponse{
		{
			ID:          "1",
			CompanyName: "Company 1",
		},
	}, nil
}
