package todo

type IToDoRepository interface {
	Create(text string, userId string) *ToDo
	ListAll() []*ToDo
}
