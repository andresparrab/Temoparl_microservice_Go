package main

import (
	"andresparrab/atest/app"
	"context"
	"log"

	"go.temporal.io/sdk/client"
)

func main() {
	// Create the client object
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "UpdtaingNumber",
		TaskQueue: app.NumberTaskQueue,
	}

	inputToTemporal := "Getting and updating the int from a mysql DB for testing purposes"

	we, err := c.ExecuteWorkflow(context.Background(), options, app.UpdateNumer, inputToTemporal)
	if err != nil {
		log.Fatalln("error starting UpdateNmber workflow", err)
	}
	printResults(we.GetID(), we.GetRunID())
	//fmt.Println("This is the we  ", we)
}

func printResults(workflowID, runID string) {
	log.Printf(
		"\nGetting  and Updating of INT from DB  \n",
	)
	log.Printf(
		"\nWorkflowID: %s RunID: %s\n",
		workflowID,
		runID,
	)
}
