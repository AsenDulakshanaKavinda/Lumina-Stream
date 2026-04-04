package utils

// this file is responsible for loading configurations from .env and YAML files using Viper and godotenv packages. 
// It defines a Config struct to hold the configuration values and a LoadConfigs function to read the configurations into the AppConfig variable. 
// Calls LoadConfigs to initialize the configuration before using it in the application.
// example usage - fmt.Println("DB Username:", utils.AppConfig.ExampleParams.Username)

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// todo -- // - change this according to the YAML structure - // todo -- //
type Config struct {
	Logging struct {
		LogFile string `mapstructure:"log_file"`
		LogDir string `mapstructure:"log_dir"`
		MaxSize string `mapstructure:"max-size"`
		MaxAge string `mapstructure:"max_age"`
		MaxBackups string `mapstructure:"max_backups"`
		Compress bool `mapstructure: "compress"`
	}

    ExampleParams struct {
        Username string `mapstructure:"username"`
        Password string `mapstructure:"password"`
        Address  string `mapstructure:"address"`
        Database string `mapstructure:"database"`
    } `mapstructure:"exampleparams"` 
    
    ServerAddr string `mapstructure:"server_address"`
}

var AppConfig Config


func LoadConfigs() {
	// - read the env file using godotenv, determine the config file name based on the ENV variable, 
	// and load the YAML configuration using Viper.
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error while loading .env file: %v", err)
	}
	config_env := os.Getenv("ENV")
	config_name := "config." + config_env 

	// - set up Viper to read the YAML configuration file, and unmarshal the configuration values into the AppConfig variable.
	viper.AddConfigPath("./configs")
	viper.SetConfigName(config_name)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error while loading configs: %v", err)
	}

	// - handle any errors that occur during the loading process and log them appropriately.
	viper.Unmarshal(&AppConfig)
}