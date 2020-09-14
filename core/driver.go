package core

import (
	"log"
	"reflect"
)

type Driver interface {
	Setup(interface{})
}

var DriverRegistry map[string]reflect.Value

func init() {
	DriverRegistry = make(map[string]reflect.Value)
}

func RegisterDriver(driver Driver) {
	t := reflect.TypeOf(driver)
	DriverRegistry[t.Elem().Name()] = reflect.ValueOf(driver)
	log.Printf("register driver: %v \n", t.Elem().Name())
}
