package main

import (
	"github.com/AntonOcean/temporal-learn/app"
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	// TODO: modify the statement below to specify the task queue name
	w := worker.New(c, "greeting-tasks", worker.Options{})

	w.RegisterWorkflow(app.GreetSomeone)
	w.RegisterActivity(app.GreetInSpanish)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
