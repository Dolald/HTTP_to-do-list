package repository

import (
	todo "todolist"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type ToDoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAllLists(userId int) ([]todo.TodoList, error)
	GetById(userId, listId int) (todo.TodoList, error)
	DeleteList(userId, listId int) error
	UpdateList(userId, listId int, list todo.UpdateListInput) error
}

type ToDoItem interface {
	Create(listId int, input todo.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todo.TodoItem, error)
	GetById(userId, listId, itemId int) (todo.TodoItem, error)
	DeleteList(userId, listId, itemId int) error
	UpdateList(userId, listId, itemId int, list todo.UpdateListInput) error
}

type Repository struct {
	Authorization
	ToDoList
	ToDoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		ToDoList:      NewToDoListPostgres(db),
		//	ToDoItem:      NewToDoItem(db),
	}
}
