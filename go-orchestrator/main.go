package main

import "orchestrator/utils"


func main() {
	utils.InitLogger()
	utils.Log.Info().Msg("Processing request")


    utils.Log.Info().Msg("User created successfully")
}