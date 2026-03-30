package memory

import (
	"sync"
)

type FavoriteRepositoryMemory struct {
	mu        sync.RWMutex
	favorites map[string]map[string]bool // userID -> map[opportunityID]bool
}

func NewFavoriteRepositoryMemory() *FavoriteRepositoryMemory {
	return &FavoriteRepositoryMemory{
		favorites: make(map[string]map[string]bool),
	}
}

func (r *FavoriteRepositoryMemory) Add(userID, oppID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.favorites[userID]; !ok {
		r.favorites[userID] = make(map[string]bool)
	}
	r.favorites[userID][oppID] = true
	return nil
}

func (r *FavoriteRepositoryMemory) Remove(userID, oppID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if m, ok := r.favorites[userID]; ok {
		delete(m, oppID)
		if len(m) == 0 {
			delete(r.favorites, userID)
		}
	}
	return nil
}

func (r *FavoriteRepositoryMemory) List(userID string) ([]string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	ids := make([]string, 0)
	if m, ok := r.favorites[userID]; ok {
		for id := range m {
			ids = append(ids, id)
		}
	}
	return ids, nil
}

func (r *FavoriteRepositoryMemory) IsFavorite(userID, oppID string) (bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if m, ok := r.favorites[userID]; ok {
		return m[oppID], nil
	}
	return false, nil
}