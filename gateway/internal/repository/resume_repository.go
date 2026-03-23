package repository

import (
	"context"
	"gateway/internal/domain"
)

type ResumeRepository interface {
	Create(ctx context.Context, r *domain.Resume) error
	GetByUserID(ctx context.Context, userID string) (*domain.Resume, error)
	Update(ctx context.Context, r *domain.Resume) error
	Delete(ctx context.Context, userID string) error
}