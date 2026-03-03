package storage

import "github.com/OlegRozh/Final-progect-todo/structs"

type TaskStorage interface {
	GetTask(id string) (*structs.Task, error)
	GetListTasks(limit int) ([]*structs.Task, error)
	CreateTask(date, title, comment, repeat string) (int64, error)
	UpdateTask(task *structs.Task) error
	DeleteTask(id string) error
	UpdateTaskDate(id, newDate string) error
}
