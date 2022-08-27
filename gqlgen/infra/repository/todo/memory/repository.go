package memory

import (
	"github.com/jupemara/graphql-languages/gqlgen/domain/model/todo"
)

type ToDoMemoryRepository struct{}

var db = []*todo.ToDo{}

func (x *ToDoMemoryRepository) Create(text string, userId string) *todo.ToDo {
	todo := todo.NewToDo(text, userId)
	db = append(db, todo)
	return todo
}

func (x *ToDoMemoryRepository) ListAll() []*todo.ToDo {
	return db
}
