package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/leducgiachoang/app-gorm/models"
)

func (r *mutationResolver) CreateMail(ctx context.Context, input models.NewMail) (*models.Mail, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateMail(ctx context.Context, id string, input models.EditMail) (*models.Mail, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteMail(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SendOrderMail(ctx context.Context, id string, mailID string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Types(ctx context.Context) ([]string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Schedules(ctx context.Context) ([]string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Mails(ctx context.Context, filter *models.MailFilter) (*models.Mails, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MailByID(ctx context.Context, id string) (*models.Mail, error) {
	panic(fmt.Errorf("not implemented"))
}
