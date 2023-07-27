package baxter

import (
	"fmt"
	"log"
)

type Store interface {
	Init() error
	Start() error
	Stop()
	Subscribe(event string, subHandler EventProcessorSignature)
	Publish(event string, meta string)
}

// Init
// The one magic thing I do is catch SIGINT and call stop on our instance
// TODO: add magic
func Init(f func() Store) error {
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
	instance          Store
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

type EventProcessorSignature func(event string, meta string)

func Subscribe(event string, subHandler EventProcessorSignature) {
	if instance == nil {
		log.Fatalln(standardComplaint)
	}
	instance.Subscribe(event, subHandler)
}

func Publish(event string, meta string) {
	// I have to publish. My life depends on it!
	if instance == nil {
		log.Fatalln(standardComplaint)
	}
	instance.Publish(event, meta)
}
