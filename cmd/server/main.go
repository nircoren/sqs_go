package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/nircoren/lightblocks/internal/server"
	"github.com/nircoren/lightblocks/internal/server/util"
	"github.com/nircoren/lightblocks/pkg/queue/sqs"
)

func main() {
	orderedMap := server.NewOrderedMap()
	logger, err := util.SetupLogger("logs/sqs_messages.log")

	if err != nil {
		fmt.Printf("Failed to set up logger: %v\n", err)
		return
	}

	config, err := godotenv.Read()
	if err != nil {
		fmt.Println("Error reading .env file: ", err)
		return
	}
	// Dependency Injection of SQS
	SQSService, err := sqs.New(config)
	if err != nil {
		fmt.Println("Error creating SQS service: ", err)
		return
	}

	queueProvider := server.NewReceiveMessagesService(SQSService)

	server.ReceiveMessages(queueProvider, orderedMap, logger)

}
