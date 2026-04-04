package main

import (
	"fmt"
	
	"orchestrator/utils"
)


func main() {
	utils.InitLogger()
	utils.LoadConfigs()
	fmt.Println("Server Address:", utils.AppConfig.ServerAddr)
    fmt.Println("DB Username:", utils.AppConfig.ExampleParams.Username)

	utils.Log.Info().Msg("Processing request")
    utils.Log.Info().Msg("User created successfully")



}