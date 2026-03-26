package repository

import (
	"context"
	"gateway/internal/domain"
)

type OpportunityRepository interface {
	Create(ctx context.Context, o *domain.Opportunity) error
	GetByID(ctx context.Context, id string) (*domain.Opportunity, error)
	Update(ctx context.Context, o *domain.Opportunity) error
	Delete(ctx context.Context, id string) error

	List(ctx context.Context) ([]*domain.Opportunity, error)
	Filter(ctx context.Context, tag string, format domain.WorkFormat) ([]*domain.Opportunity, error)
}
