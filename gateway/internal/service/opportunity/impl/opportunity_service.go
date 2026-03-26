package impl

import (
	"context"
	"gateway/internal/domain"
	"gateway/internal/repository"
	"gateway/internal/service/opportunity"
	"time"
)

type opportunityService struct {
	repo repository.OpportunityRepository
}

func NewOpportunityService(repo repository.OpportunityRepository) opportunity.OpportunityService {
	return &opportunityService{
		repo: repo,
	}
}

func (s *opportunityService) Create(ctx context.Context, o *domain.Opportunity, role domain.Role) error {

	// только работодатель может создавать
	if role != domain.RoleEmployer {
		return domain.ErrInvalidRole
	}

	// базовая валидация
	if o.Title == "" || o.Description == "" {
		return domain.ErrInvalidData
	}

	// срок жизни (по умолчанию 30 дней)
	if o.ExpiredAt.IsZero() {
		o.ExpiredAt = time.Now().Add(30 * 24 * time.Hour)
	}

	return s.repo.Create(ctx, o)
}

func (s *opportunityService) Get(ctx context.Context, id string) (*domain.Opportunity, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *opportunityService) Update(ctx context.Context, o *domain.Opportunity, userID string) error {
	existing, err := s.repo.GetByID(ctx, o.ID)
	if err != nil {
		return err
	}

	if existing.CompanyID != userID {
		return domain.ErrForbidden
	}

	return s.repo.Update(ctx, o)
}

func (s *opportunityService) Delete(ctx context.Context, id string, userID string) error {
	existing, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if existing.CompanyID != userID {
		return domain.ErrForbidden
	}

	return s.repo.Delete(ctx, id)
}

func (s *opportunityService) List(ctx context.Context) ([]*domain.Opportunity, error) {
	return s.repo.List(ctx)
}

func (s *opportunityService) Filter(
	ctx context.Context,
	tag string,
	format domain.WorkFormat,
) ([]*domain.Opportunity, error) {
	return s.repo.Filter(ctx, tag, format)
}
