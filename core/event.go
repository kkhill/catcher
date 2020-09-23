package core

import (
	"container/list"
	"log"
	"time"
)

type EventType string

// define basic event types that core system will use
const (
	StateChanged    EventType = "StateChanged"
	PropertyChanged EventType = "PropertyChanged"
	ServiceCalled   EventType = "ServiceCalled"
	PluginLoaded    EventType = "PluginLoaded"
	SystemStarted   EventType = "SystemStarted"
	SystemStopped   EventType = "SystemStopped"
)

type Event struct {
	EventType EventType
	Source    string
	Timestamp time.Time
	Data      interface{}
}

type EventBus struct {
	// listen the EventType, but fire with a Event
	bus map[EventType]*list.List // EventType -> [ EventConsumer ]
}

func (eventBus *EventBus) Listen(eventType EventType, consumer EventConsumer) {

	if consumers, ok := eventBus.bus[eventType]; !ok {
		consumers = list.New()
		consumers.PushBack(consumer)
		eventBus.bus[eventType] = consumers
	} else {
		consumers.PushBack(consumer)
	}
}

func (eventBus *EventBus) Fire(event *Event) {

	log.Printf("fire event: %v \n", event.EventType)
	if consumers, ok := eventBus.bus[event.EventType]; ok {
		for e := consumers.Front(); e != nil; e = e.Next() {
			consumer := e.Value.(EventConsumer)
			err := consumer.Execute(event)
			if err != nil {
				log.Printf("failed to execute consumer function \n")
			}
			log.Printf("executed consumer function %v \n", event.EventType)
		}
	}
}

type EventConsumer interface {
	Execute(e *Event) error
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
