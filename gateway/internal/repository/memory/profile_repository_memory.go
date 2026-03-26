package memory

import (
	"context"
	"gateway/internal/domain"
	"gateway/internal/repository"
	"sync"
	"time"

	"github.com/google/uuid"
)

type ProfileRepositoryMemory struct {
	mu sync.RWMutex

	students  map[string]*domain.StudentProfile
	employers map[string]*domain.EmployerProfile
}

func NewProfileRepositoryMemory() repository.ProfileRepository {
	return &ProfileRepositoryMemory{
		students:  make(map[string]*domain.StudentProfile),
		employers: make(map[string]*domain.EmployerProfile),
	}
}

func (r *ProfileRepositoryMemory) CreateStudent(ctx context.Context, p *domain.StudentProfile) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.students[p.UserID]; exists {
		return domain.ErrUserAlreadyExists
	}

	p.ID = uuid.NewString()
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	r.students[p.UserID] = p
	return nil
}

func (r *ProfileRepositoryMemory) GetStudent(ctx context.Context, userID string) (*domain.StudentProfile, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	p, ok := r.students[userID]
	if !ok {
		return nil, domain.ErrUserNotFound
	}

	return p, nil
}

func (r *ProfileRepositoryMemory) UpdateStudent(ctx context.Context, p *domain.StudentProfile) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	existing, ok := r.students[p.UserID]
	if !ok {
		return domain.ErrUserNotFound
	}

	p.ID = existing.ID
	p.CreatedAt = existing.CreatedAt
	p.UpdatedAt = time.Now()

	r.students[p.UserID] = p
	return nil
}

func (r *ProfileRepositoryMemory) ListStudents(ctx context.Context) ([]*domain.StudentProfile, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var res []*domain.StudentProfile
	for _, s := range r.students {
		res = append(res, s)
	}

	return res, nil
}

func (r *ProfileRepositoryMemory) CreateEmployer(ctx context.Context, p *domain.EmployerProfile) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.employers[p.UserID]; exists {
		return domain.ErrUserAlreadyExists
	}

	p.ID = uuid.NewString()
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	r.employers[p.UserID] = p
	return nil
}

func (r *ProfileRepositoryMemory) GetEmployer(ctx context.Context, userID string) (*domain.EmployerProfile, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	p, ok := r.employers[userID]
	if !ok {
		return nil, domain.ErrUserNotFound
	}

	return p, nil
}

func (r *ProfileRepositoryMemory) UpdateEmployer(ctx context.Context, p *domain.EmployerProfile) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	existing, ok := r.employers[p.UserID]
	if !ok {
		return domain.ErrUserNotFound
	}

	p.ID = existing.ID
	p.CreatedAt = existing.CreatedAt
	p.Verified = existing.Verified
	p.UpdatedAt = time.Now()

	r.employers[p.UserID] = p
	return nil
}

func (r *ProfileRepositoryMemory) ListEmployers(ctx context.Context) ([]*domain.EmployerProfile, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var res []*domain.EmployerProfile
	for _, e := range r.employers {
		res = append(res, e)
	}

	return res, nil
}
