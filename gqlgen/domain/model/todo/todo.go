package todo

type ToDo struct {
	text   string
	userId string
}

func NewToDo(text string, userId string) *ToDo {
	return &ToDo{
		text:   text,
		userId: userId,
	}
}

func (x *ToDo) Text() string {
	return x.text
}

func (x *ToDo) UserId() string {
	return x.userId
}
