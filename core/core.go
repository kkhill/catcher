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

type EventBus struct {
	bus map[*Event][]*Service
}

func (eventBus *EventBus) Listen(event *Event, service *Service) {
	if services, ok := eventBus.bus[event]; !ok {
		eventBus.bus[event] = []*Service{service}
	} else {
		flag := true
		for _, item := range services {
			if item.id == service.id {
				flag = false
				break
			}
		}
		if flag {
			services = append(services, service)
		}
	}
}

func (eventBus *EventBus) Fire(event *Event) {
	for _, service := range eventBus.bus[event] {
		go service.action(event.EventData)
	}
}

type Thing struct {
	state      string
	properties map[string]interface{}
	service    []*Service
}

type Monitor struct {
	things []*Thing
}
