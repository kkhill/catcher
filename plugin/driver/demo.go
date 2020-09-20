// a driver file, developer can create another package to define things
package driver

import (
	"catcher/core"
	"catcher/plugin/driver/demo"
)

type DriverDemo struct{} // a driver object

// Must registry this driver object to plugin registry, otherwise, core system will not know it.
func init() {

	// register driver of Things
	var driver core.Plugin
	driver = &DriverDemo{}
	core.Root.PluginRegistry.Register(driver)
}

// create things if needed
func (d *DriverDemo) Setup(config interface{}) {

	// construct a thing
	var thing core.Thing
	thing = &demo.ThingDemo{
		Brightness:  50,
		Temperature: 24.7,
	}
	core.Root.Monitor.RegistryThing(thing)
}
