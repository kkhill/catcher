package demo

import (
	"fmt"
)

type ThingDemo struct {
	// define properties
	Brightness  int32
	Temperature float32
}

func (demo *ThingDemo) GetName() string {
	// return a lovely name. name will be the part of id which is unique
	return "lovely"
}

func (demo *ThingDemo) GetStates() []string {

	// register all states of thing, and the first one will be init state
	return []string{"off", "on"}
}

func (demo *ThingDemo) GetProperties() map[string]interface{} {

	// register all properties which the thing has
	m := make(map[string]interface{})
	m["brightness"] = demo.Brightness
	m["temperature"] = demo.Temperature
	return m
}

func (demo *ThingDemo) GetServices() map[string]func(map[string]interface{}) (interface{}, error) {

	// register all services which the thing supports
	m := make(map[string]func(map[string]interface{}) (interface{}, error))
	m["Open"] = demo.Open
	m["Close"] = demo.Close
	return m
}

// the function should be func(interface{}) (interface{}, error))
func (demo *ThingDemo) Open(map[string]interface{}) (interface{}, error) {
	fmt.Println("open demo thing")
	return nil, nil
}

func (demo *ThingDemo) Close(map[string]interface{}) (interface{}, error) {
	fmt.Println("close demo thing")
	return nil, nil
}
