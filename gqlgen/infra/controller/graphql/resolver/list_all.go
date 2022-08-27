package resolver

import (
	"context"

	"github.com/jupemara/graphql-languages/gqlgen/graph/model"
	"github.com/jupemara/graphql-languages/gqlgen/usecase/todo"
)

type ListAllTodoResolver struct {
	usecase todo.ListAllToDoUsecase
}

func (x *ListAllTodoResolver) Handle(ctx context.Context) ([]*model.Todo, error) {
	vs := x.usecase.Execute()
	todos := []*model.Todo{}
	for _, v := range vs {
		todos = append(todos, &model.Todo{
			ID:   v.UserId(),
			Text: v.Text(),
		})
	}
	return todos, nil
}
