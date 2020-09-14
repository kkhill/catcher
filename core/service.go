package core

type Service struct {
	id     string
	name   string
	action func(interface{}) interface{}
}

type ServiceRegistry struct {
	services map[string][]string // serviceId -> thingId
}
