package main

import (
	"catcher/core"
	_ "catcher/driver"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"reflect"
)

func Start() {

	data, err := ioutil.ReadFile("config/driver.yaml")
	if err != nil {
		log.Fatalln("Can not open driver.yaml")
	}
	drivers := make(map[string]interface{})
	err = yaml.Unmarshal(data, drivers)
	if err != nil {
		log.Fatalln("Syntax err in driver.yaml")
	}

	log.Println("Start loading drivers...")
	for name, config := range drivers {
		if driver, ok := core.DriverRegistry[name]; ok {
			driver.MethodByName("Setup").Call([]reflect.Value{reflect.ValueOf(config)})
			log.Printf("set up driver: %v \n", name)
		}
	}

	log.Println("Start loading automations...")

	// test event bus
	//core.Root.EventBus.Listen(core.ServiceCalled, )

	// test monitor
	things := core.Root.Monitor.GetThingsId()
	fmt.Println(things)
	services := core.Root.Monitor.GetServicesOfThing("lovely")
	fmt.Println(services)
	core.Root.Monitor.CallService("lovely", "Open", make(map[string]interface{}))

}

func main() {
	Start()
}
