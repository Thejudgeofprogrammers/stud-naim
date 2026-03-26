package memory

import (
	"context"
	"gateway/internal/domain"
	"gateway/internal/repository"
	"sync"
	"time"

	"github.com/google/uuid"
)

type OpportunityRepositoryMemory struct {
	mu sync.RWMutex

	opportunities map[string]*domain.Opportunity
}

func NewOpportunityRepositoryMemory() repository.OpportunityRepository {
	return &OpportunityRepositoryMemory{
		opportunities: make(map[string]*domain.Opportunity),
	}
}

func (r *OpportunityRepositoryMemory) Create(ctx context.Context, o *domain.Opportunity) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	o.ID = uuid.NewString()
	o.CreatedAt = time.Now()

	r.opportunities[o.ID] = o
	return nil
}

func (r *OpportunityRepositoryMemory) GetByID(ctx context.Context, id string) (*domain.Opportunity, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	o, ok := r.opportunities[id]
	if !ok {
		return nil, domain.ErrUserNotFound
	}

	return o, nil
}

func (r *OpportunityRepositoryMemory) Update(ctx context.Context, o *domain.Opportunity) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	existing, ok := r.opportunities[o.ID]
	if !ok {
		return domain.ErrUserNotFound
	}

	o.CreatedAt = existing.CreatedAt

	r.opportunities[o.ID] = o
	return nil
}

func (r *OpportunityRepositoryMemory) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.opportunities[id]; !ok {
		return domain.ErrUserNotFound
	}

	delete(r.opportunities, id)
	return nil
}

func (r *OpportunityRepositoryMemory) List(ctx context.Context) ([]*domain.Opportunity, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var res []*domain.Opportunity

	for _, o := range r.opportunities {
		res = append(res, o)
	}

	return res, nil
}

func (r *OpportunityRepositoryMemory) Filter(
	ctx context.Context,
	tag string,
	format domain.WorkFormat,
) ([]*domain.Opportunity, error) {

	r.mu.RLock()
	defer r.mu.RUnlock()

	var res []*domain.Opportunity

	for _, o := range r.opportunities {

		if format != "" && o.Format != format {
			continue
		}

		// фильтр по тегу
		if tag != "" {
			found := false
			for _, t := range o.Tags {
				if t == tag {
					found = true
					break
				}
			}
			if !found {
				continue
			}
		}

		res = append(res, o)
	}

	return res, nil
}