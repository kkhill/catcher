package core

import (
	"log"
	"reflect"
	"time"
)

// every is a thing
// four critical elements of a thing: name, state, property, service
type Thing interface {
	GetName() string     // name will be part of identification
	GetStates() []string // the states that thing can be
	GetProperties() map[string]interface{}
	GetServices() map[string]func(map[string]interface{}) (interface{}, error) // service function
}

// monitor things and all operations of a thing should be execute through monitor
type Monitor struct {
	things     map[string]Thing
	states     map[string]map[string]bool        // current state will be true
	properties map[string]map[string]interface{} // property could be any type
	services   map[string]map[string]bool        // services set
	eventBus   *EventBus
}

func NewMonitor(bus *EventBus) *Monitor {
	return &Monitor{
		things:     make(map[string]Thing),
		states:     make(map[string]map[string]bool),
		properties: make(map[string]map[string]interface{}),
		services:   make(map[string]map[string]bool),
		eventBus:   bus,
	}
}

// use methods defined by interface to get information
func (monitor *Monitor) RegistryThing(thing Thing) {

	// save name
	// TODO id should be unique, and semantic
	id := thing.GetName()
	monitor.things[id] = thing

	// save states
	states := thing.GetStates()
	monitor.states[id] = make(map[string]bool)
	for i, state := range states {
		if i == 0 {
			monitor.states[id][state] = true
		} else {
			monitor.states[id][state] = false
		}
	}
	// save properties
	monitor.properties[id] = thing.GetProperties()
	// save services
	services := thing.GetServices()
	monitor.services[id] = make(map[string]bool)
	for service, _ := range services {
		monitor.services[id][service] = true
	}
}

func (monitor *Monitor) GetThingsId() []string {

	things := make([]string, len(monitor.things))
	i := 0
	for thing, _ := range monitor.things {
		things[i] = thing
		i++
	}
	return things
}

func (monitor *Monitor) GetServicesOfThing(thing string) []string {

	if services, ok := monitor.services[thing]; ok {
		ss := make([]string, len(services))
		i := 0
		for k, _ := range services {
			ss[i] = k
			i++
		}
		return ss
	} else {
		return []string{}
	}
}

func (monitor *Monitor) GetState(id string) string {
	if states, ok := monitor.states[id]; ok {
		for state, v := range states {
			if v {
				return state
			}
		}
	} else {
		log.Printf("can not find this thing: %v\n", id)
	}
	log.Printf("unknown bug, there is active state for this thing! %v \n", id)
	return ""
}

func (monitor *Monitor) SetState(id string, state string) {

	if states, ok := monitor.states[id]; ok {
		if _, ok := states[state]; ok {
			// make other state false
			data := make(map[string]interface{})
			changed := false
			for k, v := range states {
				if v == true && k != state {
					states[k] = false
					changed = true
					data["old_state"] = k
					break
				}
			}
			if changed {
				states[state] = true
				data["new_state"] = state
				monitor.eventBus.Fire(&Event{
					EventType: StateChanged,
					Source:    "unknown",
					Timestamp: time.Time{},
					Data:      data,
				})
				log.Printf("state changed: %v %v\n", id, state)
			}
		} else {
			log.Printf("can not find this state: %v\n", state)
		}
	} else {
		log.Printf("can not find this thing: %v\n", id)
	}
}

func (monitor *Monitor) GetProperty(id string, name string) interface{} {

	if properties, ok := monitor.properties[id]; ok {
		if v, ok := properties[name]; ok {
			return v
		} else {
			log.Printf("can not find this property: %v\n", name)
			return nil
		}
	} else {
		log.Printf("can not find this thing: %v\n", id)
		return nil
	}
}

func (monitor *Monitor) GetProperties(id string) map[string]interface{} {

	if properties, ok := monitor.properties[id]; ok {
		return properties
	} else {
		log.Printf("can not find this thing: %v\n", id)
		return nil
	}
}

func (monitor *Monitor) SetProperty(id string, name string, value interface{}) {

	if properties, ok := monitor.properties[id]; ok {
		if v, ok := properties[name]; ok {
			if v != value {
				properties[name] = value
				data := make(map[string]interface{})
				data["property"] = name
				data["old_value"] = v
				data["new_value"] = value
				monitor.eventBus.Fire(&Event{
					EventType: PropertyChanged,
					Source:    "unknown",
					Timestamp: time.Time{},
					Data:      data,
				})
				log.Printf("property changed: %v %v\n", id, name)
			}
		} else {
			log.Printf("can not find this property: %v\n", name)
		}
	} else {
		log.Printf("can not find this thing: %v\n", id)
	}
}

func (monitor *Monitor) SetProperties(id string, properties map[string]interface{}) {

	if ps, ok := monitor.properties[id]; ok {
		data := make([]map[string]interface{}, 0)
		for name, value := range ps {
			if value != properties[name] {
				d := make(map[string]interface{})
				d["property"] = name
				d["old_value"] = value
				d["new_value"] = properties[name]
				data = append(data, d)
			}
		}
		if len(data) != 0 {
			monitor.eventBus.Fire(&Event{
				EventType: PropertyChanged,
				Source:    "unknown",
				Timestamp: time.Time{},
				Data:      data,
			})
		}
	} else {
		log.Printf("can not find this thing: %v\n", id)
	}
}

func (monitor *Monitor) CallService(id string, service string, data interface{}) (interface{}, error) {

	if services, ok := monitor.services[id]; ok {
		if services[service] {
			thing := monitor.things[id]
			thingObject := reflect.ValueOf(thing)
			param := reflect.ValueOf(data) // can not be zero value
			thingObject.MethodByName("Open").Call([]reflect.Value{param})
			log.Printf("service called: %v.%v\n", id, service)
		} else {
			log.Printf("can not find this service: %v\n", service)
		}
	} else {
		log.Printf("can not find this thing: %v\n", id)
	}
	// fire service called event
	monitor.eventBus.Fire(&Event{
		EventType: ServiceCalled,
		Source:    "unknown",
		Timestamp: time.Now(),
		Data:      data,
	})
	// TODO return value
	return nil, nil
}
