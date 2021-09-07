package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	generated1 "github.com/leducgiachoang/app-gorm/generated"
	"github.com/leducgiachoang/app-gorm/models"
	"github.com/leducgiachoang/app-gorm/app"
	"github.com/google/uuid"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input models.NewTodo) (*models.Todo, error) {
	todo := models.Todo{
		ID:   string(uuid.NewString()),
		Text: input.Text,
	}

	if err := app.DB.FirstOrCreate(&todo, models.Todo{Text: todo.Text, Done: true}).Error; err != nil{
		return nil, err
	}

	return &todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	var todos []*models.Todo

	if err := app.DB.Find(&todos).Error; err != nil{
		return nil, err
	}

	return todos, nil
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
