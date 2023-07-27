package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/17twenty/baxter"
)

func main() {

	baxter.Init(baxter.InMemory(10)) // We can pass in different backing stores for our queue management

	baxter.Subscribe("event.test", ProcessEvent)
	baxter.Subscribe("event.test", func(event string, meta json.RawMessage) {
		log.Println("I AM THE GREATEST", string(meta))
		time.Sleep(4 * time.Second)
		baxter.Publish("event.test", []byte("OMG NO"))
	})

	baxter.Start()

	baxter.Publish("event.test", []byte("hello"))

	log.Println("Sleeping main for 10...")
	time.Sleep(time.Second * 10)

	log.Println("Stopping baxter")
	baxter.Stop()
}

func ProcessEvent(event string, meta json.RawMessage) {
	log.Println("Received event", event, "with", string(meta))
}
