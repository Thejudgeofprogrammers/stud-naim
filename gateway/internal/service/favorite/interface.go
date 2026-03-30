package favorite

import "context"

type FavoriteService interface {
	Add(ctx context.Context, userID, oppID string) error
	Remove(ctx context.Context, userID, oppID string) error
	List(ctx context.Context, userID string) ([]interface{}, error) // вернёт список объектов Opportunity
	IsFavorite(ctx context.Context, userID, oppID string) (bool, error)
}