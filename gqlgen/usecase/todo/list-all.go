package todo

import (
	"github.com/jupemara/graphql-languages/gqlgen/domain/model/todo"
)

func NewListAllToDoUsecase(
	repository todo.IToDoRepository,
) *ListAllToDoUsecase {
	return &ListAllToDoUsecase{
		repository: repository,
	}
}

type ListAllToDoUsecase struct {
	repository todo.IToDoRepository
}

func (x *ListAllToDoUsecase) Execute() []*todo.ToDo {
	return x.repository.ListAll()
}
