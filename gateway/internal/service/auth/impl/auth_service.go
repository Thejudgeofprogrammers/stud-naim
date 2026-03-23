package impl

import (
	"context"
	"gateway/internal/domain"
	"gateway/internal/repository"
	"gateway/internal/service/auth"
	"gateway/internal/service/jwt"
	"gateway/internal/service/refresh"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	refreshService refresh.RefreshService
	jwtService     jwt.JWTService
	userRepo       repository.UserRepository
}

func NewAuthService(
	refreshService refresh.RefreshService,
	jwtService jwt.JWTService,
	userRepo repository.UserRepository,
) auth.AuthService {
	return &authService{
		refreshService: refreshService,
		jwtService:     jwtService,
		userRepo:       userRepo,
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

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	var userRole domain.Role
	switch role {
	case string(domain.RoleStudent):
		userRole = domain.RoleStudent
	case string(domain.RoleEmployer):
		userRole = domain.RoleEmployer
	default:
		return domain.ErrInvalidRole
	}

	user := &domain.User{
		Email:     email,
		Password:  string(hash),
		Role:      userRole,
		CreatedAt: time.Now(),
	}

	return s.userRepo.Create(ctx, user)
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
