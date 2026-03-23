package memory

import (
	"context"
	"gateway/internal/domain"
	"gateway/internal/repository"
	"sync"
	"time"

	"github.com/google/uuid"
)

type ResumeRepositoryMemory struct {
	mu sync.RWMutex

	resumes map[string]*domain.Resume
}

func NewResumeRepositoryMemory() repository.ResumeRepository {
	return &ResumeRepositoryMemory{
		resumes: make(map[string]*domain.Resume),
	}
}

func (r *ResumeRepositoryMemory) Create(ctx context.Context, res *domain.Resume) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	res.ID = uuid.NewString()
	res.CreatedAt = time.Now()

	r.resumes[res.UserID] = res
	return nil
}

func (r *ResumeRepositoryMemory) GetByUserID(ctx context.Context, userID string) (*domain.Resume, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	res, ok := r.resumes[userID]
	if !ok {
		return nil, domain.ErrResumeNotFound
	}

	return res, nil
}

func (r *ResumeRepositoryMemory) Update(ctx context.Context, res *domain.Resume) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.resumes[res.UserID] = res
	return nil
}

func (r *ResumeRepositoryMemory) Delete(ctx context.Context, userID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.resumes, userID)
	return nil
}
