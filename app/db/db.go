package db

import (
	"github.com/bryutus/caspian-serverside/app/conf"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db gorm.DB

func Connect() *gorm.DB {

	db, err := gorm.Open(conf.GetDbDriver(), conf.GetDbConnect())
	if err != nil {
		panic("Failed to connect to database.")
	}

	return db
}
