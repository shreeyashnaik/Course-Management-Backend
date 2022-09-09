package config

import (
	"fmt"
	"strconv"

	"github.com/spf13/viper"
)

var (
	ENVIRONMENT          = ""
	PORT                 = "8080"
	CORS_ALLOWED_ORIGINS = "http://localhost:3030, http://localhost:3031"
	DB_URI               = ""
	JWT_SECRET_KEY       = ""
)

func LoadConfigVars() {
	ENVIRONMENT = viper.GetString("ENVIRONMENT")
	PORT = strconv.Itoa(viper.GetInt("PORT"))
	username := viper.GetString("DB_USER")
	password := viper.GetString("DB_PASS")
	dbName := viper.GetString("DB_NAME")
	dbHost := viper.GetString("DB_HOST")
	DB_URI = fmt.Sprintf("host=%s user=%s dbname=%s port=5432 sslmode=require password=%s", dbHost, username, dbName, password)
}
