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
	profile, err := s.profileRepo.GetStudent(ctx, id)
	if err != nil {
		return nil, err
	}

	return &dto.StudentResponse{
		ID:     profile.UserID,
		Name:   profile.FullName,
		Skills: profile.Skills,
		About:  profile.About,
		Resume: "", // потом подтянем из resumeService
	}, nil
}

func (s *userService) UpdateStudent(ctx context.Context, id string, req dto.UpdateStudentRequest) error {
	profile, err := s.profileRepo.GetStudent(ctx, id)
	if err != nil {
		return err
	}

	profile.FullName = req.Name
	profile.Skills = req.Skills
	profile.About = req.About

	return s.profileRepo.UpdateStudent(ctx, profile)
}

func (s *userService) GetEmployer(ctx context.Context, id string) (*dto.EmployerResponse, error) {
	profile, err := s.profileRepo.GetEmployer(ctx, id)
	if err != nil {
		return nil, err
	}

	return &dto.EmployerResponse{
		ID:             profile.UserID,
		CompanyName:    profile.CompanyName,
		Description:    profile.Description,
		Representative: profile.Representative,
		Vacancies:      []string{}, // позже добавим
	}, nil
}

func (s *userService) UpdateEmployer(ctx context.Context, id string, req dto.UpdateEmployerRequest) error {
	profile, err := s.profileRepo.GetEmployer(ctx, id)
	if err != nil {
		return err
	}

	profile.CompanyName = req.CompanyName
	profile.Description = req.Description
	profile.Representative = req.Representative

	return s.profileRepo.UpdateEmployer(ctx, profile)
}

func (s *userService) ListStudents(ctx context.Context) ([]dto.StudentResponse, error) {
	profiles, err := s.profileRepo.ListStudents(ctx)
	if err != nil {
		return nil, err
	}

	var result []dto.StudentResponse

	for _, p := range profiles {
		result = append(result, dto.StudentResponse{
			ID:     p.UserID,
			Name:   p.FullName,
			Skills: p.Skills,
			About:  p.About,
			Resume: "",
		})
	}

	return result, nil
}

func (s *userService) ListEmployers(ctx context.Context) ([]dto.EmployerResponse, error) {
	profiles, err := s.profileRepo.ListEmployers(ctx)
	if err != nil {
		return nil, err
	}

	var result []dto.EmployerResponse

	for _, p := range profiles {
		result = append(result, dto.EmployerResponse{
			ID:          p.UserID,
			CompanyName: p.CompanyName,
		})
	}

	return result, nil
}
