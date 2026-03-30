package response

import (
	"context"
	"gateway/internal/domain"
)

type ResponseService interface {
	Create(ctx context.Context, userID, oppID, message string) error
	ListByUser(ctx context.Context, userID string) ([]domain.Response, error)
	UpdateStatus(ctx context.Context, responseID string, status domain.ResponseStatus) error
}