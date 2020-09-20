package plugin

import (
	_ "catcher/plugin/addon"
	_ "catcher/plugin/driver"
	"log"
)

func init() {
	log.Println("init plugins")
}
