package main

import (
	"fmt"
	"orchestrator/utils"

	"github.com/spf13/viper"
)


func main() {
	utils.InitLogger()
	utils.LoadCofigs()

	utils.Log.Info().Msg("Processing request")
    utils.Log.Info().Msg("User created successfully")

	config := viper.GetStringMapString("example")
	fmt.Println(config)
	fmt.Println(config["key1"])

}