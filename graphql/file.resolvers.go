package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/leducgiachoang/app-gorm/models"
)

func (r *mutationResolver) CreateSignedURL(ctx context.Context, filename string, fileMime string) (*models.FileSigned, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateFile(ctx context.Context, input models.NewFile) (*models.File, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateFileUUID(ctx context.Context, fileID string, uuid string) (*models.File, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetFileURL(ctx context.Context, fileID string) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetFileURLFromKey(ctx context.Context, key string) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetFileThumbnailURL(ctx context.Context, fileID string) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}
