package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/bryutus/caspian-back/app/models"
)

var db gorm.DB

func main() {
	db := connect()
	defer db.Close()

	db.DropTableIfExists(&models.History{}, &models.Resource{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.History{}, &models.Resource{})
}

func connect() *gorm.DB {
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
