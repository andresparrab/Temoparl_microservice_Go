package main

import (
	"fmt"
	"log"

	"andresparrab/atest/app"

	_ "github.com/go-sql-driver/mysql"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	app.ConnectToMySQL()
	// Create the temporal client object just once per process
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()
	// This worker hosts both Worker and Activity functions
	fmt.Println("Started the worker")
	numberWoker := worker.New(c, app.NumberTaskQueue, worker.Options{})

	fmt.Println("Register the workflow app.UpdateNumer ")
	numberWoker.RegisterWorkflow(app.UpdateNumer)

	fmt.Println("Register the Activity  app.GetInt ")
	numberWoker.RegisterActivity(app.GetInt)

	fmt.Println("Register the app.UpdateInt ")
	numberWoker.RegisterActivity(app.UpdateInt)

	// Start listening to the Task Queue
	err = numberWoker.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
