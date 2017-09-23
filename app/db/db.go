package db

import (
	"fmt"

	"github.com/bryutus/caspian-serverside/app/conf"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db gorm.DB

func Connect() *gorm.DB {

	var config conf.Config
	err := conf.LoardConf(&config)
	if err != nil {
		panic(err)
	}

	driver := config.Database.Driver
	connect := fmt.Sprintf("%s:%s@%s/%s?charset=%s&parseTime=%s", config.Database.User, config.Database.Pass, config.Database.Protocol, config.Database.Database, config.Database.Charset, config.Database.ParseTime)

	db, err := gorm.Open(driver, connect)
	if err != nil {
		panic("Failed to connect to database.")
	}

	return db
}
