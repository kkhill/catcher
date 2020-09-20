package addon

import (
	"catcher/core"
	engine "catcher/plugin/addon/rulengine"
	"catcher/utils"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path"
)

type RuleEngine struct{}

func init() {

	// register plugin
	var plugin core.Plugin
	plugin = &RuleEngine{}
	core.Root.PluginRegistry.Register(plugin)
	log.Println("I have be initialized")
}

func (r *RuleEngine) Setup(config interface{}) {

	dataPath := os.Getenv(utils.DATA_PATH)
	rulePath := path.Join(dataPath, utils.CONFIG_DIR, utils.RULE_FILE)
	data, err := ioutil.ReadFile(rulePath)
	if err != nil {
		log.Println("failed to load rule.yaml")
		log.Println(err)
		return
	}

	rulesData := make([]map[string]interface{}, 0)
	err = yaml.Unmarshal(data, &rulesData)
	if err != nil {
		log.Println("failed to unmarshal rule.yaml")
		log.Println(err)
		return
	}
	rules := engine.ParseRules(rulesData)
	engine.ActivateRules(rules)
}
