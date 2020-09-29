package repository
import "test-layered/domain/model"

type ITodoRepository interface {
	GetTodoList() ([]model.Todo, error)
}