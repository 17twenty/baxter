package baxter

import (
	"context"
	"log"
)

// This guy needs a mutex AND a waitgroup
//
//	lock           = &sync.Mutex{}
type inMemory struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func InMemory(depth int) func() Store {
	return func() Store {
		return &inMemory{}
	}
}

func (i *inMemory) Init() error {
	i.ctx, i.cancelFunc = context.WithCancel(context.Background())
	return nil
}

func (i *inMemory) Subscribe(event string, subHandler EventProcessorSignature) {

	log.Println("Subscribing...")
}

func (i *inMemory) Start() error {
	// I can fail to start if something is horribly wrong...
	// unlikely on inMemory
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done(): // if cancel() execute
				return
			default:
				i.getAndDispatch()
			}
		}
	}(i.ctx)

	return nil
}

func (i *inMemory) getAndDispatch() {
	// I grab the latest message from the queue and send it to all the subscribers
}

func (i *inMemory) Stop() {
	// TODO start the consumer that roundrobins the
	i.cancelFunc()
}

func (i *inMemory) Publish(event string, meta string) {
	// I have to publish. My life depends on it!
}
