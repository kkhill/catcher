package core

import (
	"catcher/common"
	"log"
	"reflect"
)

type Monitor struct {
	things   map[string]common.Thing
	services map[string]map[string]bool // id -> services set
}

func NewMonitor() *Monitor {
	return &Monitor{
		things:   make(map[string]common.Thing),
		services: make(map[string]map[string]bool),
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

func (monitor *Monitor) RegistryThing(thing common.Thing) {

	// generate id of thing
	// TODO id should be unique, and semantic
	id := thing.GetName()
	monitor.things[thing.GetName()] = thing
	// save services
	services := thing.GetServices()
	if _, ok := monitor.services[id]; !ok {
		monitor.services[id] = make(map[string]bool)
	}
	ss := monitor.services[id]
	for k, _ := range services {
		ss[k] = true
	}
}

func (monitor *Monitor) CallService(id string, service string, data map[string]interface{}) (interface{}, error) {
	// call service
	if services, ok := monitor.services[id]; ok {
		if services[service] {
			thing := monitor.things[id]
			thingObject := reflect.ValueOf(thing)
			param := reflect.ValueOf(data) // can not be zero value
			thingObject.MethodByName("Open").Call([]reflect.Value{param})
		} else {
			log.Printf("can not find this service: %v\n", service)
		}
	} else {
		log.Printf("can not find this thing: %v\n", id)
	}
	// fire service called event

	return nil, nil
}
