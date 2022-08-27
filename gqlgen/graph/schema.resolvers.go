package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/jupemara/graphql-languages/gqlgen/graph/generated"
	"github.com/jupemara/graphql-languages/gqlgen/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	v := memoryRepository.Create(input.Text, input.UserID)
	return &model.Todo{
		Text: v.Text(),
		ID:   v.UserId(),
	}, nil
}

// DeleteTodo is the resolver for the deleteTodo field.
func (r *mutationResolver) DeleteTodo(ctx context.Context, id string) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: DeleteTodo - deleteTodo"))
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	vs := todosUsecase.Execute()
	todos := []*model.Todo{}
	for _, v := range vs {
		todos = append(todos, &model.Todo{
			ID:   v.UserId(),
			Text: v.Text(),
		})
	}
	return todos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
