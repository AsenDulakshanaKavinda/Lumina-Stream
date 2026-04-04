package utils

// This file contains the logic to read the configuration from the YAML file
// and unmarshal it into the Config struct. 
// It uses the Viper library to read the configuration file and the godotenv 
// library to load environment variables from a .env file. 
// The GetConfig function is a singleton that ensures that the configuration is loaded only once
// and can be accessed globally throughout the application.
// `config_schemas.go`: contains the definition of the Config struct. 

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)


var (
	instance *Config
	once sync.Once
	config Config
)

func GetConfig() *Config {
    once.Do(func() {
        instance = loadConfigs()
    })
    return instance
}


func loadConfigs() *Config {
	// - read config YAML file based on the ENV variable 
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error while loading .env file: %v", err)
	}
	config_env := os.Getenv("ENV")
	config_name := "config." + config_env 

	// - set up Viper to read the YAML configuration file
	viper.AddConfigPath("./configs")
	viper.SetConfigName(config_name)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error while loading configs: %v", err)
	}

	// - unmarshal the configs into the Config struct
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Error unmarshaling configs: %v", err)
	}

	return &config
}