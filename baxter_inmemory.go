package baxter

import (
	"context"
	"sync"
)

type Queue struct {
	sync.Mutex
	events []Event
}

func (q *Queue) Push(evt Event) {
	q.Lock()
	defer q.Unlock()
	q.events = append(q.events, evt)
}

func (q *Queue) Pop() Event {
	q.Lock()
	defer q.Unlock()
	if len(q.events) == 0 {
		return Event{}
	}
	item := q.events[0]
	q.events[0] = Event{}
	q.events = q.events[1:]
	return item
}

type inMemory struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
	//
	// Subscribers
	subMutex sync.Mutex
	subs     []Subscriber
	// Message Queue
	events Queue
}

func InMemory(capacity int) func() BaxterProvider {
	return func() BaxterProvider {
		return &inMemory{
			events: Queue{
				events: make([]Event, 0, capacity),
			},
		}
	}
}

func (inst *inMemory) Init() error {
	// This guy has nothing to do
	// for in memory but would defintely fail otherwise
	return nil
}

func (inst *inMemory) Subscribe(event string, subHandler EventProcessorSignature) {
	inst.subMutex.Lock()
	defer inst.subMutex.Unlock()

	inst.subs = append(inst.subs, Subscriber{
		eventName: event,
		callback:  subHandler,
	})
}

func (inst *inMemory) Start() error {
	inst.ctx, inst.cancelFunc = context.WithCancel(context.Background())

	// I can fail to start if something is horribly wrong...
	// unlikely on inMemory
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done(): // if cancel() execute
				return
			default:
				inst.getAndDispatch()
			}
		}
	}(inst.ctx)

	return nil
}

func (inst *inMemory) getAndDispatch() {
	// I grab the latest message(s) from the queue and send it
	// to all my subscribers

	latestEvent := inst.events.Pop()
	if !latestEvent.IsEmpty() {
		// Dispatch to all
		for j := range inst.subs {
			if inst.subs[j].eventName == latestEvent.eventName {
				inst.subs[j].callback(latestEvent.eventName, string(latestEvent.meta))
			}
		}
	}
}

func (inst *inMemory) Stop() {
	inst.cancelFunc()
}

func (inst *inMemory) Publish(event string, meta string) {
	// I have to publish. My life depends on it!
	inst.events.Push(Event{
		event, []byte(meta),
	})
}
