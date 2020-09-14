package core

type Core struct {
	ServiceRegistry *ServiceRegistry
	EventBus        *EventBus
	Monitor         *Monitor
}

var Root *Core

func init() {
	bus := &EventBus{}
	Root = &Core{
		EventBus: bus,
		Monitor:  NewMonitor(bus),
	}
}
