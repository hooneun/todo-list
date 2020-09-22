package dblayer

import (
	"github.com/hooneun/todo-list/app/models"
)

// DBLayer ...
type DBLayer interface {
	AddUser(models.User) (models.User, error)
	GetUserByID(int) (models.User, error)
}
