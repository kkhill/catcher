package main

import (
	"catcher/core"
	_ "catcher/plugin"
	"catcher/utils"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func Start() {

	// set environment
	// TODO make user data in /home/user
	dir, _ := os.Getwd()
	dataPath := path.Join(dir, utils.DATA_DIR)
	os.Setenv(utils.DATA_PATH, dataPath)

	// initialize driver
	driverPath := path.Join(dataPath, utils.CONFIG_DIR, utils.THING_FILE)
	data, err := ioutil.ReadFile(driverPath)
	if err != nil {
		log.Fatalln("Can not open things.yaml")
	}
	drivers := make(map[string]interface{})
	err = yaml.Unmarshal(data, drivers)
	if err != nil {
		log.Fatalln("Syntax err in things.yaml")
	}

	log.Println("Start loading driver...")
	core.Root.PluginRegistry.Load(drivers)

	// initialize plugins
	pluginPath := path.Join(dataPath, utils.CONFIG_DIR, utils.PLUGIN_FILE)
	data, err = ioutil.ReadFile(pluginPath)
	if err != nil {
		log.Fatalln("Can not open plugins.yaml")
	}
	plugins := make(map[string]interface{})
	err = yaml.Unmarshal(data, plugins)
	if err != nil {
		log.Fatalln("Syntax err in plugins.yaml")
	}
	log.Println("Start loading plugins...")
	core.Root.PluginRegistry.Load(plugins)

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
