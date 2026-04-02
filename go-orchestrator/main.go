package main

import "orchestrator/utils"


func main() {
	utils.InitLogger()
	utils.Log.Info().
        Str("action", "create_user").
        Str("status", "attempting").
        Msg("Processing request")


    utils.Log.Info().Msg("User created successfully")
}