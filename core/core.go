package core

import (
	"reflect"
)

type Core struct {
	ServiceRegistry *ServiceRegistry
	EventBus        *EventBus
	Monitor         *Monitor
	ThingRegistry   map[string]reflect.Type
}

type Service struct {
	id     string
	name   string
	action func(interface{}) interface{}
}

type ServiceRegistry struct {
	services map[string][]string // serviceId -> thingId
}

type Event struct {
	EventType string
	EventData interface{}
}

var Root *Core

func init() {
	Root = &Core{
		EventBus: &EventBus{},
		Monitor:  NewMonitor(),
	}
}
