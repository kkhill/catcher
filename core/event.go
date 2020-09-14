package core

import (
	"container/list"
	"log"
	"time"
)

type EventType int32

const (
	StateChanged    EventType = 0
	PropertyChanged EventType = 1
	ServiceCalled   EventType = 2
)

type Event struct {
	EventType EventType
	Source    string
	Timestamp time.Time
	Data      interface{}
}

type EventBus struct {
	bus map[EventType]*list.List // EventType -> [ EventConsumer ]
}

func (eventBus *EventBus) Listen(event EventType, consumer *EventConsumer) {

	if consumers, ok := eventBus.bus[event]; !ok {
		consumers = list.New()
		consumers.PushBack(consumer)
		eventBus.bus[event] = consumers
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
			log.Printf("execute consumer function %v \n", event.EventType)
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
