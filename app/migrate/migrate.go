package main

import (
	"github.com/bryutus/caspian-serverside/app/db"
	"github.com/bryutus/caspian-serverside/app/models"
)

func main() {
	db, _ := db.Connect()
	defer db.Close()

	db.DropTableIfExists(&models.Resource{}, &models.History{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.History{}, &models.Resource{}).AddForeignKey("history_id", "histories(id)", "RESTRICT", "RESTRICT")
}
