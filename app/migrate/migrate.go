package main

import (
	"github.com/bryutus/caspian-serverside/app/db"
	"github.com/bryutus/caspian-serverside/app/models"
)

func main() {
	db := db.Connect()
	defer db.Close()

	db.DropTableIfExists(&models.History{}, &models.Resource{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.History{}, &models.Resource{})
}
