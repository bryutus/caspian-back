package main

import (
	"github.com/bryutus/caspian-back/app/db"
	"github.com/bryutus/caspian-back/app/models"
)

func main() {
	db := db.Connect()
	defer db.Close()

	db.DropTableIfExists(&models.History{}, &models.Resource{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.History{}, &models.Resource{})
}
