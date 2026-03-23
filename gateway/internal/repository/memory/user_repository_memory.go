package memory


import (
	"context"
	"gateway/internal/domain"
	"gateway/internal/repository"
	"sync"

	"github.com/google/uuid"
)

type UserRepositoryMemory struct {
	mu sync.RWMutex

	usersByID    map[string]*domain.User
	usersByEmail map[string]*domain.User
}

func NewUserRepositoryMemory() repository.UserRepository {
	return &UserRepositoryMemory{
		usersByID:    make(map[string]*domain.User),
		usersByEmail: make(map[string]*domain.User),
	}
}

func (r *UserRepositoryMemory) Create(ctx context.Context, user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.usersByEmail[user.Email]; exists {
		return domain.ErrUserAlreadyExists
	}

	user.ID = uuid.NewString()

	r.usersByID[user.ID] = user
	r.usersByEmail[user.Email] = user

	return nil
}

func (r *UserRepositoryMemory) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	u, ok := r.usersByEmail[email]
	if !ok {
		return nil, domain.ErrUserNotFound
	}

	return u, nil
}

func (r *UserRepositoryMemory) GetByID(ctx context.Context, id string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	u, ok := r.usersByID[id]
	if !ok {
		return nil, domain.ErrUserNotFound
	}

	return u, nil
}