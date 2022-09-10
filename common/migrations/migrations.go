package migrations

import (
	"log"

	"github.com/shreeyashnaik/Course-Management-Backend/common/db"
	"github.com/shreeyashnaik/Course-Management-Backend/pkg/models"
)

func MigrateModels() {
	database := db.GetDB()

	database.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	// Define role as enum
	database.Exec("CREATE TYPE role AS ENUM ('superadmin', 'admin', 'employee');")

	// Setup Join table for viewed courses
	if err := database.SetupJoinTable(
		&models.User{},
		"ViewedCourses",
		&models.ViewedCourses{},
	); err != nil {
		log.Panic(err)
	}

	if err := database.AutoMigrate(
		&models.User{},
		&models.Course{},
	); err != nil {
		log.Panic(err)
	}

}
