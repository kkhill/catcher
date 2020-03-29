package core

type Event struct {
	//EventSource string
	EventType string
	EventData interface{}
}

type Action struct {
	call func()
}

type EventBus struct {
	bus map[Event][]Action
}

func (eventBus *EventBus) Listen(event Event, action Action) {
	if Actions, ok := eventBus.bus[event]; !ok {
		eventBus.bus[event] = append([]Action{}, action)
	} else {
		Actions = append(Actions, action)
	}
}

func (eventBus *EventBus) Trigger(event Event) {

	for _, action := range eventBus.bus[event] {
		go action.call()
	}

}
