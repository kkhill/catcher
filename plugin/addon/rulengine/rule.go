package rulengine

import (
	"errors"
	"log"
)

type Event struct {
	Type   string
	Things []string
	Data   map[string]interface{}
}

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

type Rule struct {
	Name      string
	Id        string
	Event     []*Event
	Condition []*Condition
	Action    []*Action
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
		name, ok := ruleData["name"].(string)
		if !ok {
			ruleData["name"] = id
		}
		eventData, ok := ruleData["event"].([]interface{})
		if !ok {
			log.Printf("fail to load rule: %v, wrong 'if' expression \n", name)
			continue
		}
		conditionData := ruleData["condition"].([]interface{})
		actionData, ok := ruleData["action"].([]interface{})
		if !ok {
			log.Printf("fail to load rule: %v, wrong 'action' expression \n", name)
			continue
		}

		events, err := parseEvents(eventData)
		if err != nil || events == nil {
			log.Printf("%v: %v \n", name, err)
			continue
		}
		// allowed to miss condition, but not event and action
		conditions, err := parseConditions(conditionData)
		if err != nil {
			log.Printf("%v: %v \n", name, err)
			continue
		}
		actions, err := parseActions(actionData)
		if err != nil || actions == nil {
			log.Printf("%v: %v \n", name, err)
			continue
		}
		rule := &Rule{
			Name:      name,
			Id:        id,
			Event:     events,
			Condition: conditions,
			Action:    actions,
		}
		rules = append(rules, rule)
	}

	return rules
}

func parseEvents(data []interface{}) ([]*Event, error) {
	if data == nil || len(data) == 0 {
		return nil, errors.New("without any event")
	}
	events := make([]*Event, len(data))
	for i, e := range data {
		ee := e.(map[interface{}]interface{})
		if _, ok := ee["type"]; !ok {
			return nil, errors.New("miss event type")
		}
		eThings := ee["thing"].([]interface{})
		things := make([]string, len(eThings))
		for i, item := range eThings {
			things[i] = item.(string)
		}
		eData := ee["data"].(map[interface{}]interface{})
		data := make(map[string]interface{}, len(eData))
		for k, v := range eData {
			data[k.(string)] = v
		}
		event := &Event{
			Type:   ee["type"].(string),
			Things: things,
			Data:   data,
		}
		events[i] = event
	}

	return events, nil
}

func parseConditions(data []interface{}) ([]*Condition, error) {
	if data == nil || len(data) == 0 {
		return nil, nil // conditions can be missed
	}
	conditions := make([]*Condition, len(data))
	for i, c := range data {
		cc := c.(map[interface{}]interface{})
		if _, ok := cc["type"]; !ok {
			return nil, errors.New("miss condition type")
		}
		cThings := cc["thing"].([]interface{})
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

	if data == nil || len(data) == 0 {
		return nil, errors.New("without any action")
	}
	actions := make([]*Action, len(data))
	for i, a := range data {
		aa := a.(map[interface{}]interface{})
		if _, ok := aa["type"]; !ok {
			return nil, errors.New("miss action type")
		}
		aThings := aa["thing"].([]interface{})
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
