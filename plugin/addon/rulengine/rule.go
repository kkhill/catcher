package rulengine

import (
	"catcher/core"
	"errors"
	"log"
)

type Condition struct {
	Type      string
	Name      string
	Things    []string
	Threshold interface{}
	Operation string
}

type Action struct {
	Type   string
	Things []string
	Name   string
}

// a rule is a bunch of conditions and actions
type Rule struct {
	Name       string
	Id         string
	Conditions []*Condition
	Actions    []*Action
}

type Rules map[core.EventType][]*Rule

func ParseRules(data []map[string]interface{}) map[core.EventType][]*Rule {

	rules := make([]*Rule, 0)
	for _, ruleData := range data {
		// TODO generate a unique id
		id := "some id"
		name, ok := ruleData["name"].(string)
		if !ok {
			ruleData["name"] = id
		}
		eventData, ok := ruleData["event"].(map[interface{}]interface{})
		if !ok {
			log.Printf("fail to load rule: %v, wrong 'if' expression \n", name)
			continue
		}
		conditionsData := ruleData["conditions"].([]interface{})
		actionsData, ok := ruleData["actions"].([]interface{})
		if !ok {
			log.Printf("fail to load rule: %v, wrong 'action' expression \n", name)
			continue
		}

		event, err := parseEvent(eventData)
		if err != nil || event == nil {
			log.Printf("%v: %v \n", name, err)
			continue
		}
		// allowed to miss condition, but not event and action
		conditions, err := parseConditions(conditionsData)
		if err != nil {
			log.Printf("%v: %v \n", name, err)
			continue
		}
		actions, err := parseActions(actionsData)
		if err != nil || actions == nil {
			log.Printf("%v: %v \n", name, err)
			continue
		}
		rule := &Rule{
			Name:       name,
			Id:         id,
			Conditions: conditions,
			Actions:    actions,
		}
		rules = append(rules, rule)
	}

	return rules
}

func parseEvent(eventData map[interface{}]interface{}) (*Event, error) {

	if _, ok := eventData["type"]; !ok {
		return nil, errors.New("miss event type")
	}
	event := &Event{
		Type:  eventData["type"].(string),
		Thing: eventData["thing"].(string),
		Data:  eventData["data"].(map[interface{}]interface{}),
	}
	return event, nil
}

func parseConditions(data []interface{}) ([]*Condition, error) {

	conditions := make([]*Condition, len(data))
	for i, c := range data {
		cc := c.(map[interface{}]interface{})
		if _, ok := cc["type"]; !ok {
			return nil, errors.New("miss condition type")
		}
		cThings := cc["things"].([]interface{})
		things := make([]string, len(cThings))
		for i, item := range cThings {
			things[i] = item.(string)
		}

		condition := &Condition{
			Type:      cc["type"].(string),
			Name:      cc["name"].(string),
			Things:    things,
			Threshold: cc["threshold"],
			Operation: cc["operation"].(string),
		}
		conditions[i] = condition
	}

	return conditions, nil
}

func parseActions(data []interface{}) ([]*Action, error) {

	actions := make([]*Action, len(data))
	for i, a := range data {
		aa := a.(map[interface{}]interface{})
		if _, ok := aa["type"]; !ok {
			return nil, errors.New("miss action type")
		}
		aThings := aa["things"].([]interface{})
		things := make([]string, len(aThings))
		for i, item := range aThings {
			things[i] = item.(string)
		}
		action := &Action{
			Type:   aa["type"].(string),
			Things: things,
			Name:   aa["name"].(string),
		}
		actions[i] = action
	}

	return actions, nil
}
