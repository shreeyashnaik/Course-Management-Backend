package migrations

import (
	"log"

	"github.com/shreeyashnaik/Course-Management-Backend/common/db"
	"github.com/shreeyashnaik/Course-Management-Backend/pkg/models"
)

func MigrateModels() {
	database := db.GetDB()
	database.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	if err := database.AutoMigrate(
		&models.User{},
		&models.Course{},
	); err != nil {
		log.Panic(err)
	}
}
