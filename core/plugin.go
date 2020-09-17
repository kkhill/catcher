package core

import (
	"log"
	"reflect"
)

// make plugin different from driver,
// because they are confused and plugins need an order, while drivers dont
type Plugin interface {
	Setup(interface{})
}

type PluginRegistry map[string]reflect.Value

func (dr PluginRegistry) RegisterPlugin(driver Driver) {

	t := reflect.TypeOf(driver)
	dr[t.Elem().Name()] = reflect.ValueOf(driver)
	log.Printf("register driver: %v \n", t.Elem().Name())
}

func (dr PluginRegistry) LoadPlugins(plugins map[string]interface{}) {

	for name, config := range plugins {
		if plugin, ok := dr[name]; ok {
			plugin.MethodByName("Setup").Call([]reflect.Value{reflect.ValueOf(config)})
			log.Printf("set up driver: %v \n", name)
		} else {
			log.Printf("can not find this driver: %v \n", name)
		}
	}
}
