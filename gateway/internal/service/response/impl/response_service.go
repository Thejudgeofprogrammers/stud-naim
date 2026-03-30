package impl

import (
	"context"
	"gateway/internal/domain"
	"gateway/internal/repository/memory"
	opportunitySvc "gateway/internal/service/opportunity"
	"gateway/internal/service/response"
)

type responseService struct {
	respRepo *memory.ResponseRepositoryMemory
	oppSvc   opportunitySvc.OpportunityService
}

func NewResponseService(respRepo *memory.ResponseRepositoryMemory, oppSvc opportunitySvc.OpportunityService) response.ResponseService {
	return &responseService{
		respRepo: respRepo,
		oppSvc:   oppSvc,
	}
}

func (s *responseService) Create(ctx context.Context, userID, oppID, message string) error {
	_, err := s.oppSvc.Get(ctx, oppID)
	if err != nil {
		return err
	}
	resp := &domain.Response{
		UserID:        userID,
		OpportunityID: oppID,
		Status:        domain.ResponseNew,
		Message:       message,
	}
	return s.respRepo.Create(resp)
}

func (s *responseService) ListByUser(ctx context.Context, userID string) ([]domain.Response, error) {
	return s.respRepo.ListByUser(userID)
}

func (s *responseService) UpdateStatus(ctx context.Context, responseID string, status domain.ResponseStatus) error {
	resp, err := s.respRepo.GetByID(responseID)
	if err != nil {
		return err
	}
	resp.Status = status
	return s.respRepo.Update(resp)
}