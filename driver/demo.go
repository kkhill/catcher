package driver

import (
	"catcher/common"
	"catcher/core"
	"fmt"
)

type DriverDemo struct{}

// add thing objects to runtime
func (demo *DriverDemo) Setup(config interface{}) {

	//data := config.(string)
	fmt.Println("Yeah! i have been setup")
	// construct a thing
	var thing common.Thing
	thing = &ThingDemo{}
	// register services for thing

	core.Root.Monitor.RegistryThing(thing)
}

type ThingDemo struct {
	brightness int
	services   map[string]func(interface{}) (interface{}, error)
}

func (demo *ThingDemo) GetName() string {
	return "lovely"
}

func (demo *ThingDemo) GetStates() []string {
	return []string{"off", "on"}
}

func (demo *ThingDemo) GetProperties() []string {
	return []string{"brightness", "power"}
}

func (demo *ThingDemo) GetServices() map[string]func(map[string]interface{}) (interface{}, error) {

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

func init() {

	// register driver of Things
	var driver common.Driver
	driver = &DriverDemo{}
	common.RegisterDriver(driver)
}
