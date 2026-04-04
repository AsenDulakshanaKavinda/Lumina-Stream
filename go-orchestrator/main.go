package main

import (
	// "fmt"

	"fmt"
	"orchestrator/utils"
)


func main() {
	utils.InitLogger()
	fmt.Println(utils.GetConfig().Logging)


	utils.Log.Info().Msg("Processing request")
    utils.Log.Info().Msg("User created successfully")



}