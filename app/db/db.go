package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db gorm.DB

func Connect() *gorm.DB {

	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "app"
	OPTION := "charset=utf8&parseTime=True"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + OPTION

	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic("Failed to connect to database.")
	}

	return db
}
