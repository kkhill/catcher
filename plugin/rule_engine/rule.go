package plugin

import (
	"log"
)

type Rule struct {
	Name string
	Id   string
	If   *If
	Then *Then
}

type If struct {
	Event    []interface{}
	State    []interface{}
	Property []interface{}
	Context  []interface{}
}

type Then struct {
	Service []interface{}
}

func ParseRules(data interface{}) []*Rule {

	rules := make([]*Rule, 0)
	rulesData, ok := data.([]map[string]interface{})
	if !ok {
		log.Println("invalid rules data")
		return rules
	}

	for _, ruleData := range rulesData {
		// TODO generate a unique id
		id := "some id"
		ruleData["id"] = id

		name, ok := ruleData["name"].(string)
		if !ok {
			ruleData["name"] = ruleData["id"]
		}
		ifData, ok := ruleData["if"]
		if !ok {
			log.Printf("fail to load rule: %v, missing if expression \n", name)
			continue
		}
		thenData, ok := ruleData["then"]
		if !ok {
			log.Printf("fail to load rule: %v, missing then expression \n", name)
			continue
		}

		// parse if
		if_, err := ParseIf(ifData)
		if err != nil {
			log.Printf("%v: %v \n", name, err)
			continue
		}
		// parse then
		then_, err := ParseThen(thenData)
		if err != nil {
			log.Printf("%v: %v \n", name, err)
			continue
		}

		rule := &Rule{
			Name: name,
			Id:   id,
			If:   if_,
			Then: then_,
		}
		rules = append(rules, rule)
	}

	return rules
}

func ParseIf(data interface{}) (*If, error) {
	return nil, nil
}

func ParseThen(data interface{}) (*Then, error) {
	return nil, nil
}
