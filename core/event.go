package core

import (
	"container/list"
	"time"
)

type EventType struct {
	name      string
	source    string
	timestamp time.Time
	data      map[string]interface{}
}

type EventBus struct {
	bus map[*EventType]*list.List // EventType -> [ EventConsumer ]
}

func (eventBus *EventBus) Listen(event *EventType, consumer *EventConsumer) {

	if consumers, ok := eventBus.bus[event]; !ok {
		consumers = list.New()
		consumers.PushBack(consumer)
		eventBus.bus[event] = consumers
	} else {
		consumers.PushBack(consumer)
	}
}

func (eventBus *EventBus) Fire(event *EventType) {

	if consumers, ok := eventBus.bus[event]; ok {
		for e := consumers.Front(); e != nil; e = e.Next() {
			consumer := e.Value.(EventConsumer)
			err := consumer.execute(event)
			if err != nil {

			}
		}
	}
}

type EventConsumer interface {
	execute(e *EventType) (err error)
}

type EventProducer struct {
}

// to process event and dispatch
type EventProcessingAgent struct {
}

type Context struct {
}

type GlobalState struct {
}
