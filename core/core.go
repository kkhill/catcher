package core

type Core struct {
	serviceRegistry *ServiceRegistry
	eventBus        *EventBus
	monitor         *Monitor
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

type Thing struct {
	state      string
	properties map[string]interface{}
	service    []*Service
}

type Monitor struct {
	things []*Thing
}
