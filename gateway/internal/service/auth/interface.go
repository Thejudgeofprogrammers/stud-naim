package auth

import (
	"context"
	"gateway/internal/service/jwt"
)

type AuthService interface {
	Register(ctx context.Context, email, password, role string) error
	Login(ctx context.Context, email, password string) (*jwt.AuthTokens, error)
	Refresh(ctx context.Context, refreshToken string) (*jwt.AuthTokens, error)
}
