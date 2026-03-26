package impl

import (
	"context"
	"gateway/internal/domain"
	"gateway/internal/repository"
	"gateway/internal/service/resume"
)

type resumeService struct {
	repo repository.ResumeRepository
}

func NewResumeService(repo repository.ResumeRepository) resume.ResumeService {
	return &resumeService{
		repo: repo,
	}
}

func (s *resumeService) GetResume(ctx context.Context, userID string) (string, error) {
	res, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
		return "", err
	}

	return res.FileURL, nil
}

func (s *resumeService) UploadResume(ctx context.Context, userID string, fileURL string) error {
	res, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
		if err == domain.ErrResumeNotFound {
			return s.repo.Create(ctx, &domain.Resume{
				UserID:  userID,
				FileURL: fileURL,
			})
		}
		return err
	}

	res.FileURL = fileURL
	return s.repo.Update(ctx, res)
}

func (s *resumeService) DeleteResume(ctx context.Context, userID string) error {
	return s.repo.Delete(ctx, userID)
}
