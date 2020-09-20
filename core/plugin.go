package core

import (
	"log"
	"reflect"
)

// make plugin different from driver,
// because they are confused and plugins need an order, while driver dont
type Plugin interface {
	Setup(interface{})
}

type PluginRegistry map[string]reflect.Value

func (dr PluginRegistry) Register(plugin Plugin) {

	t := reflect.TypeOf(plugin)
	dr[t.Elem().Name()] = reflect.ValueOf(plugin)
	log.Printf("register plugin: %v \n", t.Elem().Name())
}

func (dr PluginRegistry) Load(plugins map[string]interface{}) {

	for name, config := range plugins {
		if plugin, ok := dr[name]; ok {
			if config == nil {
				config = new(string)
			} // function Call receive at least one argument
			args := []reflect.Value{reflect.ValueOf(config)}
			plugin.MethodByName("Setup").Call(args)
			log.Printf("set up plugin: %v \n", name)
		} else {
			log.Printf("can not find this plugin: %v \n", name)
		}
	}
}
