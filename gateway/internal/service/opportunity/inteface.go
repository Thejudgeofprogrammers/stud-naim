package opportunity

import (
	"context"
	"gateway/internal/domain"
)

type OpportunityService interface {
	Create(ctx context.Context, o *domain.Opportunity, role domain.Role) error

	Get(ctx context.Context, id string) (*domain.Opportunity, error)

	Update(ctx context.Context, o *domain.Opportunity, userID string) error

	Delete(ctx context.Context, id string, userID string) error

	List(ctx context.Context) ([]*domain.Opportunity, error)

	Filter(
		ctx context.Context,
		tag string,
		format domain.WorkFormat,
	) ([]*domain.Opportunity, error)
}
