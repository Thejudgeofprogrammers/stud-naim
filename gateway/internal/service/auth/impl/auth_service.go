package impl

import (
	"context"
	"gateway/internal/domain"
	"gateway/internal/repository"
	"gateway/internal/service/auth"
	"gateway/internal/service/jwt"
	"gateway/internal/service/refresh"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	refreshService refresh.RefreshService
	jwtService     jwt.JWTService
	userRepo       repository.UserRepository
	profileRepo    repository.ProfileRepository
}

func NewAuthService(
	refreshService refresh.RefreshService,
	jwtService jwt.JWTService,
	userRepo repository.UserRepository,
	profileRepo repository.ProfileRepository,
) auth.AuthService {
	return &authService{
		refreshService: refreshService,
		jwtService:     jwtService,
		userRepo:       userRepo,
		profileRepo:    profileRepo,
	}
}

func parseRole(role string) (domain.Role, error) {
	switch role {
	case string(domain.RoleStudent):
		return domain.RoleStudent, nil
	case string(domain.RoleEmployer):
		return domain.RoleEmployer, nil
	case string(domain.RoleCurator):
		return domain.RoleCurator, nil
	default:
		return "", domain.ErrInvalidRole
	}
}

func (s *authService) Refresh(ctx context.Context, refreshToken string) (*jwt.AuthTokens, error) {
	userID, err := s.refreshService.Validate(ctx, refreshToken)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	newAccess, err := s.jwtService.GenerateAccessToken(user.ID, jwt.AccessRole(user.Role))
	if err != nil {
		return nil, err
	}

	newRefresh, err := s.refreshService.Create(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &jwt.AuthTokens{
		AccessToken:  newAccess,
		RefreshToken: newRefresh,
	}, nil
}

func (s *authService) Register(ctx context.Context, email, password, role string) error {
	_, err := s.userRepo.GetByEmail(ctx, email)
	if err == nil {
		return domain.ErrUserAlreadyExists
	}

	userRole, err := parseRole(role)
	if err != nil {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &domain.User{
		ID:        uuid.NewString(),
		Email:     email,
		Password:  string(hash),
		Role:      userRole,
		CreatedAt: time.Now(),
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return err
	}

	switch userRole {

	case domain.RoleStudent:
		err = s.profileRepo.CreateStudent(ctx, &domain.StudentProfile{
			UserID:   user.ID,
			FullName: "",
			Skills:   []string{},
			About:    "",
		})
		if err != nil {
			return err
		}

	case domain.RoleEmployer:
		err = s.profileRepo.CreateEmployer(ctx, &domain.EmployerProfile{
			UserID:         user.ID,
			CompanyName:    "",
			Description:    "",
			Representative: "",
			Verified:       false,
		})
		if err != nil {
			return err
		}

	case domain.RoleCurator:
		// пока без профиля (добавить)
	}

	return nil
}

func (s *authService) Login(ctx context.Context, email, password string) (*jwt.AuthTokens, error) {
	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, domain.ErrUserNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, domain.ErrInvalidCredentials
	}

	access, err := s.jwtService.GenerateAccessToken(user.ID, jwt.AccessRole(user.Role))
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.refreshService.Create(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return &jwt.AuthTokens{
		AccessToken:  access,
		RefreshToken: refreshToken,
	}, nil
}
