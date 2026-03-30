package memory

import (
	"gateway/internal/domain"
	"sync"
	"time"

	"github.com/google/uuid"
)

type ResponseRepositoryMemory struct {
	mu        sync.RWMutex
	responses map[string]domain.Response // key = id
	byUser    map[string][]string        // userID -> []responseID
	byOpp     map[string][]string        // oppID -> []responseID
}

func NewResponseRepositoryMemory() *ResponseRepositoryMemory {
	return &ResponseRepositoryMemory{
		responses: make(map[string]domain.Response),
		byUser:    make(map[string][]string),
		byOpp:     make(map[string][]string),
	}
}

func (r *ResponseRepositoryMemory) Create(resp *domain.Response) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	resp.ID = uuid.New().String()
	resp.AppliedAt = time.Now()
	r.responses[resp.ID] = *resp
	r.byUser[resp.UserID] = append(r.byUser[resp.UserID], resp.ID)
	r.byOpp[resp.OpportunityID] = append(r.byOpp[resp.OpportunityID], resp.ID)
	return nil
}

func (r *ResponseRepositoryMemory) Update(resp *domain.Response) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.responses[resp.ID]; !ok {
		return domain.ErrNotFound
	}
	r.responses[resp.ID] = *resp
	return nil
}

func (r *ResponseRepositoryMemory) GetByID(id string) (*domain.Response, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if resp, ok := r.responses[id]; ok {
		return &resp, nil
	}
	return nil, domain.ErrNotFound
}

func (r *ResponseRepositoryMemory) ListByUser(userID string) ([]domain.Response, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	ids := r.byUser[userID]
	respList := make([]domain.Response, 0, len(ids))
	for _, id := range ids {
		respList = append(respList, r.responses[id])
	}
	return respList, nil
}