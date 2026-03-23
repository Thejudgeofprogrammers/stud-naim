package resume

import "context"

type ResumeService interface {
	GetResume(ctx context.Context, userID string) (string, error)
	UploadResume(ctx context.Context, userID string, fileName string) error
	DeleteResume(ctx context.Context, userID string) error
}
