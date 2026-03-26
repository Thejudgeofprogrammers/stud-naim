package resume

import "context"

type ResumeService interface {
	GetResume(ctx context.Context, userID string) (string, error)
	UploadResume(ctx context.Context, userID string, fileURL string) error
	DeleteResume(ctx context.Context, userID string) error
}
