package core

import "container/list"

type Core struct {
	eventBus *EventBus
	monitor  *Monitor
	storage  *Storage
	things   list.List
}

var core *Core

func init() {
	core = &Core{
		eventBus: &EventBus{bus: make(map[Event][]Action)},
		monitor:  &Monitor{},
		storage:  &Storage{},
		things:   *list.New(),
	}
}
