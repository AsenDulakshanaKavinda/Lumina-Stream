package utils

// This file contains functions related to handling configurations, such as loading environment variables and configuration files.
// It uses the "github.com/joho/godotenv" package to load environment variables from a .env file and the "github.com/spf13/viper" package to manage configuration files.

/* 
utils.LoadCofigs()

config := viper.GetStringMapString("example")
fmt.Println(config)
fmt.Println(config["key1"])
*/


import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)


func LoadCofigs() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error while loading .env file: %v", err)
	}
	config_env := os.Getenv("ENV")
	config_name := "config." + config_env + ".yaml"

	viper.AddConfigPath("./configs")
	viper.SetConfigName(config_name)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error while loading configs: %v", err)
	}
}