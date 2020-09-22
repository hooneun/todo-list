package dblayer

import (
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
