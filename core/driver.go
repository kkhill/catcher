package core

import (
	"log"
	"reflect"
)

type Driver interface {
	Setup(interface{})
}

type DriverRegistry map[string]reflect.Value

func (dr DriverRegistry) RegisterDriver(driver Driver) {

	t := reflect.TypeOf(driver)
	dr[t.Elem().Name()] = reflect.ValueOf(driver)
	log.Printf("register driver: %v \n", t.Elem().Name())
}

func (dr DriverRegistry) LoadDrivers(drivers map[string]interface{}) {

	for name, config := range drivers {
		if driver, ok := dr[name]; ok {
			driver.MethodByName("Setup").Call([]reflect.Value{reflect.ValueOf(config)})
			log.Printf("set up driver: %v \n", name)
		} else {
			log.Printf("can not find this driver: %v \n", name)
		}
	}
}
