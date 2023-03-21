package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// To create a connection to database
var (
	DBConn *gorm.DB // gorm is an ORM to connect to the database
)
