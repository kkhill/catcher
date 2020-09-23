package core

import (
	"container/list"
	"reflect"
)

var Root *Core

type Core struct {
	//DriverRegistry  DriverRegistry // driver can not be dependent any other
	PluginRegistry  PluginRegistry
	ServiceRegistry *ServiceRegistry
	EventBus        *EventBus
	Monitor         *Monitor
}

func init() {

	bus := &EventBus{
		bus: make(map[EventType]*list.List),
	}
	Root = &Core{
		//DriverRegistry: make(map[string]reflect.Value),
		PluginRegistry: make(map[string]reflect.Value),
		EventBus:       bus,
		Monitor:        NewMonitor(bus),
	}
}
