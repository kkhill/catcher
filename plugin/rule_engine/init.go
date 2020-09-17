package plugin

import (
	"catcher/core"
	"io/ioutil"
	"log"
)

type RuleEngine struct{}

func init() {

	// register plugin
	var plugin core.Plugin
	plugin = &RuleEngine{}
	core.Root.PluginRegistry.RegisterPlugin(plugin)
}

func (demo *RuleEngine) Setup(config interface{}) {

	data, err := ioutil.ReadFile("config/rule.yaml")
	if err != nil {
		log.Fatalln("Syntax error in rule.yaml")
	}
	rules := ParseRules(data)
	RegisterRules(rules)
}

func RegisterRules(rules []*Rule) {
	// TODO: put rules to core system
}
