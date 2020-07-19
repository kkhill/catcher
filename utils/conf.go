package utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path"
)

func LoadComponents(dir string) map[string]interface{} {
	// TODO 加载configuration
	p := path.Join(dir, CONFIGURATION)
	data, err := ioutil.ReadFile(p)
	if err != nil {
		log.Fatalln(err)
	}
	components := make(map[string]interface{})
	yaml.Unmarshal(data, components)
	fmt.Println(components)
	return components
}

func loadAutomation() {
	// TODO 加载automation
}
