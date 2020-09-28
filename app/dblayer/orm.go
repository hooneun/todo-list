package dblayer

import (
	"github.com/hooneun/todo-list/app/helper"
	"github.com/hooneun/todo-list/app/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ORM ...
type ORM struct {
	*gorm.DB
}

// NewORM ...
func NewORM() (*ORM, error) {
	dns := "root@tcp(127.0.0.1:3306)/todo_list?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	return &ORM{
		DB: db,
	}, err
}

// AddUser create usre
func (db *ORM) AddUser(user models.User) (models.User, error) {
	helper.MakeHash(&user.Password)
	err := db.Create(&user).Error
	user.Password = ""

	return user, err
}
