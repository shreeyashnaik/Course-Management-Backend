package main

import (
	"github.com/shreeyashnaik/Course-Management-Backend/common/db"
	"github.com/shreeyashnaik/Course-Management-Backend/common/migrations"
	"github.com/shreeyashnaik/Course-Management-Backend/common/utils"
	"github.com/shreeyashnaik/Course-Management-Backend/config"
	"github.com/shreeyashnaik/Course-Management-Backend/src/core/server"
	"github.com/spf13/viper"
)

func main() {
	// Imports env variables
	utils.ImportEnv()

	// Loads env variables into config variables
	config.LoadConfigVars()

	// Migrate the models to DB as tables
	if viper.GetBool("MIGRATE") {
		migrations.MigrateModels()
	}

	// Initialize DB Services
	db.InitServices()

	server.StartCoreServer()
}
