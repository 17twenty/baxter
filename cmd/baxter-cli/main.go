package main

import (
	"log"
	"time"

	"github.com/17twenty/baxter"
)

func main() {

	baxter.Init(baxter.InMemory()) // We can pass in different backing stores for our queue management

	baxter.Subscribe("event.test", ProcessEvent)

	baxter.Start()

	baxter.Publish("event.test", "hello")

	log.Println("Sleeping main for 10...")
	time.Sleep(time.Second * 10)

	log.Println("Stopping baxter")
	baxter.Stop()
}

func ProcessEvent(event, meta string) {
	// I can't fail, I _have_ to react
	log.Println("Received event", event, "with", meta)
}