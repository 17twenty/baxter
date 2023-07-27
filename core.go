package baxter

import (
	"encoding/json"
	"fmt"
	"log"
)

type Subscriber struct {
	eventName string
	callback  EventProcessorSignature
}

type Event struct {
	eventName string
	meta      json.RawMessage
}

func (evt Event) IsEmpty() bool {
	return evt.eventName == "" && evt.meta == nil
}

type BaxterProvider interface {
	Init() error
	Start() error
	Stop()
	Subscribe(event string, subHandler EventProcessorSignature)
	Publish(event string, meta json.RawMessage)
}

// Init
// The one magic thing I do is catch SIGINT and call stop on our instance
// TODO: add magic
func Init(f func() BaxterProvider) error {
	s := f()
	err := s.Init()
	if err != nil {
		return fmt.Errorf("baxter: %s", err)
	}
	instance = s
	return nil
}

var (
	// instance is a singleton. You don't want multiple baxters.
	instance          BaxterProvider
	standardComplaint = "baxter: You haven't called Init() -- and if you did, it errored"
)

// All these guys do is delegate to the singleton
func Start() error {
	// I can fail to start if something is horribly wrong
	if instance == nil {
		log.Fatalln(standardComplaint)
	}
	return instance.Start()
}

func Stop() {
	if instance == nil {
		log.Fatalln(standardComplaint)
	}
	instance.Stop()
}

type EventProcessorSignature func(event string, meta json.RawMessage)

func Subscribe(event string, subHandler EventProcessorSignature) {
	if instance == nil {
		log.Fatalln(standardComplaint)
	}
	instance.Subscribe(event, subHandler)
}

func Publish(event string, meta json.RawMessage) {
	// I have to publish. My life depends on it!
	if instance == nil {
		log.Fatalln(standardComplaint)
	}
	instance.Publish(event, meta)
}
