package main

import (
	"context"
	"github.com/AntonOcean/temporal-learn/app"
	"go.temporal.io/sdk/client"
	"log"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "my-first-workflow-2",
		TaskQueue: "greeting-tasks",
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, app.GreetSomeone, "kek 2")
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}
	log.Println("Workflow result:", result)
}
