package impl

import (
	"context"
	"gateway/internal/repository/memory"
	"gateway/internal/service/favorite"
	opportunitySvc "gateway/internal/service/opportunity"
)

type favoriteService struct {
	favRepo *memory.FavoriteRepositoryMemory
	oppSvc  opportunitySvc.OpportunityService
}

func NewFavoriteService(favRepo *memory.FavoriteRepositoryMemory, oppSvc opportunitySvc.OpportunityService) favorite.FavoriteService {
	return &favoriteService{
		favRepo: favRepo,
		oppSvc:  oppSvc,
	}
}

func (s *favoriteService) Add(ctx context.Context, userID, oppID string) error {
	_, err := s.oppSvc.Get(ctx, oppID)
	if err != nil {
		return err
	}
	return s.favRepo.Add(userID, oppID)
}

func (s *favoriteService) Remove(ctx context.Context, userID, oppID string) error {
	return s.favRepo.Remove(userID, oppID)
}

func (s *favoriteService) List(ctx context.Context, userID string) ([]interface{}, error) {
	oppIDs, err := s.favRepo.List(userID)
	if err != nil {
		return nil, err
	}
	opps := make([]interface{}, 0, len(oppIDs))
	for _, id := range oppIDs {
		opp, err := s.oppSvc.Get(ctx, id)
		if err != nil {
			continue // пропускаем, если opportunity удалена
		}
		opps = append(opps, opp)
	}
	return opps, nil
}

func (s *favoriteService) IsFavorite(ctx context.Context, userID, oppID string) (bool, error) {
	return s.favRepo.IsFavorite(userID, oppID)
}