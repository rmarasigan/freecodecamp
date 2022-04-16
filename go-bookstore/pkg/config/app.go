package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// This is used for the database connection
func Connect() {
	conn, err := gorm.Open("mysql", "root:This is a Very Hard Password 200%@(127.0.0.1:3306)/go_bookstore?charset=utf8&parseTime=True")
	if err != nil {
		panic(err)
	}

	db = conn
}

func GetDB() *gorm.DB {
	return db
}
