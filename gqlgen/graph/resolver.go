package graph

//go:generate go run github.com/99designs/gqlgen generate

import (
	"github.com/jupemara/graphql-languages/gqlgen/infra/repository/todo/memory"
	"github.com/jupemara/graphql-languages/gqlgen/usecase/todo"

	"github.com/jupemara/graphql-languages/gqlgen/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos []*model.Todo
}

var (
	memoryRepository = &memory.ToDoMemoryRepository{}
	todosUsecase     = todo.NewListAllToDoUsecase(memoryRepository)
)
