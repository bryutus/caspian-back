package db

import (
	"fmt"

	"github.com/bryutus/caspian-serverside/app/conf"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db gorm.DB

func Connect() (*gorm.DB, error) {

	db, err := gorm.Open(conf.GetDbDriver(), conf.GetDbConnect())
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to database: %s", err.Error())
	}

	return db, nil
}
