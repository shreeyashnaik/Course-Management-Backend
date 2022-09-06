package utils

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// This function is responsible for setting default environment variables and importing env.
func ImportEnv() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetDefault("PORT", 3000)
	viper.SetDefault("MIGRATE", false)
	viper.SetDefault("SERVER_TYPE", "core")
	viper.SetDefault("ENVIRONMENT", "development")

	// checks if environment variables match any of the existing keys and loads them.
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found ignoring error
		} else {
			log.Panicln(fmt.Errorf("fatal error config file: %s", err))
		}
	}
}
