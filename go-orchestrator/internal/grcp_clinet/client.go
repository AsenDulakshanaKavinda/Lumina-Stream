package grpc_client

// This file contains the gRPC client code to connect to the embedder server and send requests.

import (
	"context"
	"fmt"
	"time"

	pb "grpc-go-server/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"orchestrator/utils"
)

func Client() {
	// setup the connection to the embedder server
	conn, err := grpc.NewClient(
		utils.GetConfig().GRPCConfig.ServerAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		utils.Log.Error().Err(err).Msg("Error while connecting to embedder server")
	}
	defer conn.Close()

	// - sending a request and getting response
	client := pb.NewEmbedderClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.Embed(ctx, &pb.TextRequest{ Text: "text to embed" })
	if err != nil {
		utils.Log.Error().Err(err).Msg("Error while receiving response")
	}
	utils.Log.Info().Msg("Received response from server")
	fmt.Println("Response: ", response.Vector)


}

